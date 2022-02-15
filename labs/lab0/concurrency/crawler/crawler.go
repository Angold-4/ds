package main
import (
    "fmt"
    "sync"
)

// Several solutions o the crawer execise from the Go tutorial
// https://go.dev/tour/concurrency/10


// Fetcher interface

type Fetcher interface {
    // Fetch returns a slice of URLS found in the page.
    Fetch(url string) (urls []string, err error)
}

type fakeFetcher map[string]*fakeResult
// fakeFether has a property, which is a map
// and it will map a string to a fakeResult(a struct)

type fakeResult struct {
    // the abstraction of a web page
    body string
    urls []string
}

func (f fakeFetcher) Fetch(url string) ([]string, error) {
    if res, ok := f[url]; ok {
	// ok is the stmt, exec f[url] first.
	// get the response
	fmt.Printf("found:   %s\n", url)
	return res.urls, nil
    }

    // missing this url
    fmt.Printf("missing: %s\n", url)
    return nil, fmt.Errorf("not found: %s", url)
}

// abstact web pages () return ans
var fetcher = fakeFetcher {
    // all web pages
    "http://golang.org/" : &fakeResult {
	"The Go Programming Language",
	[]string {
	    "http://golang.org/pkg/",
	    "http://golang.org/cmd/",
	},
    },

    "http://golang.org/pkg/" : &fakeResult {
	"Packages",
	[]string {
	    "http://golang.org/",
	    "http://golang.org/cmd/",
	    "http://golang.org/pkg/fmt/",
	    "http://golang.org/pkg/os/",
	},
    },

    "http://golang.org/pkg/fmt/" : &fakeResult {
	"Package fmt",
	[]string {
	    "http://golang.org/",
	    "http://golang.org/pkg/",
	},
    },

    "http://golang.org/pkg/os/" : &fakeResult {
	"Package os",
	[]string {
	    "http://golang.org/",
	    "http://golang.org/pkg/",
	},
    },

}


// 1. Serial crawer
func Serial(url string, fetcher Fetcher, fetched map[string]bool) {
    // DFS Search
    if fetched[url] {
	// if already fetched that url, return
	return
    }
    fetched[url] = true;

    // deal with that url
    urls, err := fetcher.Fetch(url)

    if err != nil {
	return
    }

    for _, u := range urls {
	Serial(u, fetcher, fetched)
	// we cannot just simply add a goroutine here
	// the reason is that once the serial distribute all goroutines
	// this function will definitely return
    }
    return
}


// 2. Concurrent crawer with shared state and Mutex

type fetchState struct {
    mu sync.Mutex
    fetched map[string]bool // a map is already a pointer in go
}

func ConcurrentMutex(url string, fetcher Fetcher, f *fetchState) {
    f.mu.Lock()
    already := f.fetched[url]
    f.fetched[url] = true
    f.mu.Unlock()

    if already {
	return
    }

    urls, err := fetcher.Fetch(url)

    if err != nil {
	return
    }

    var done sync.WaitGroup

    for _, u := range urls {
	done.Add(1)
	go func(u string) {
	    defer done.Done()
	    // always call Done when function finished
	    // no matter why the surrounding function is finished
	    ConcurrentMutex(u, fetcher, f)
	}(u)
    }

    done.Wait() // wait all go routines complete (done get down to zero)
    return
}

func makeState() *fetchState {
    f := &fetchState{}
    f.fetched = make(map[string]bool)
    return f
}


// 3. Concurrent crawler with channels

// communicate information instead of getting through shared memory

func worker(url string, ch chan []string, fetcher Fetcher) {
    urls, err := fetcher.Fetch(url)
    if err != nil {
	ch <- []string{}
    } else {
	ch <- urls
    }
}

func master(ch chan []string, fetcher Fetcher) {
    n := 1
    fetched := make(map[string]bool)
    for urls := range ch {
	for _, u := range urls {
	    if fetched[u] == false {
		fetched[u] = true
		n += 1
		go worker(u, ch, fetcher)
	    }
	}
	n -= 1
	if n == 0 {
	    break
	}
    }
}

func ConcurrentChannel(url string, fetcher Fetcher) {
    ch := make(chan []string)
    go func() {
	ch <- []string{url}
    }()
    master(ch, fetcher)
}


// main entry

func main() {
    fmt.Printf("=== Serial ===\n")
    Serial("http://golang.org/", fetcher, make(map[string]bool))

    fmt.Printf("=== ConcurrentMutex ===\n")
    ConcurrentMutex("http://golang.org/", fetcher, makeState())

    fmt.Printf("=== ConcurrentChannel ===\n")
    ConcurrentChannel("http://golang.org/", fetcher)
}






























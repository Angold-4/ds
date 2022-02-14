package main

import (
    "fmt"
)


// One of the most ubiquitous interfaces is Stringer, defined by the fmt package
// the fmt package look for this interface to print values (Println)

// type Stringer interface {
//     String() string
// }


// We can use this way to define a print format

type Person struct {
    Name string
    Age int
}

func (p Person) String() string {
    return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

type IPAddr [4]byte

func (ip IPAddr) String() string {
    return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
}

func main() {
    a := Person{"Angold Wang", 20}

    fmt.Println(a)

    hosts := map[string] IPAddr {
	"loopback": {127, 0, 0, 1},
	"googleDNS": {8, 8, 8, 8},
    }

    for name, ip := range hosts {
	fmt.Printf("%v: %v\n", name, ip)
    }
}



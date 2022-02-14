package main

import (
    "golang.org/x/tour/wc"
    "strings"
)

func WordCount(s string) map[string]int {
    words := strings.Split(s, " ")
    m := make(map[string]int)
    for i:=0; i<len(words); i++ {
	m[words[i]] += 1
    }
    return m
}

func main() {
    wc.Test(WordCount)
}

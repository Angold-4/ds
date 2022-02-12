package main

import (
    "fmt"
    "log"
    "rsc.io/quote"
    "example.com/greetings"
)

func main() {
    // Set properties of the predifined Logger
    // The log entry prefix and flag to disable printing
    // the time, source file and line number
    log.SetPrefix("greetiings: ")
    log.SetFlags(0)

    fmt.Println(quote.Go());
    message, err := greetings.Hello("Angold")

    names := []string{"Wange", "Jiawei", "Darrin", "Laura"}

    messages, err := greetings.Hellos(names)


    // If an error was returned, print it to the console
    if err != nil {
	log.Fatal(err)
    }

    fmt.Println(message)
    fmt.Println(messages)
}


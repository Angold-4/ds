package greetings

import (
    "errors"
    "fmt"
    "math/rand"
    "time"
)

// Hello returns a greeting for the named person
// In Go, a function whose name starts with a capital letter 
// can be called by a function ot in the same packge. (exported name)
func Hello(name string) (string, error) {
    // Return a greeting that embeds the name in the message.

    // If no name was given, return an error with a message
    if name == "" {
	return "", errors.New("empty name")
    }
    message := fmt.Sprintf(randomFormat(), name) // return
    // := operator is a shortcut for declaring and init a variable in one line
    // (Go uses the value on the right to determine the variable's type)
    // var message string
    // message = fmt.Sprintf("Hi, %v. Welcome!" ,name)
    return message, nil
}


// Hellos returns a map that associates each of the named people
// with a greeting message
func Hellos(names []string) (map[string]string, error) {
    // a map to associate names with messages
    // Create a messages map to associate each of the recieved names
    // the make is the way of init map
    messages := make(map[string]string)

    // Loop through the received slice of names
    // Calling the Hello function to get a messsage for each name
    for _, name := range names {
	message, err := Hello(name)
	if err != nil {
	    return nil, err
	}
	messages[name] = message
    }
    return messages, nil
}

// init sets initial values for variables used in he function
// Note that Go executes init functions automatically at program startup
func init() {
    rand.Seed(time.Now().UnixNano())
}

// randomFormat returns one of a set of greeting messages
// The returned message is selected at random 
func randomFormat() string {
    // A slice of message formats
    formats := []string {
	"Hi, %v. Welcome!",
	"Great to see you, %v!",
	"Hail, %v! Well met!",
    }

    // Return a randomly selected message format by
    // specifying a random index for the slice of the formats.
    return formats[rand.Intn(len(formats))]
}

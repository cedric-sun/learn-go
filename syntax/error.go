package main

import (
	"fmt"
	"time"
	"strconv"
)

type MyError struct {
	When time.Time
	What string
}

// implement the 'error' bulid-in interface
func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	// it's *MyError that implements error
	return &MyError { time.Now(), "it didn't work" }
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}

	i, err := strconv.Atoi("42")
	if err != nil {
		fmt.Printf("Couldn't convert number: %v\n", err)
		return
	}
	fmt.Println("Converted integer: ",i)
}

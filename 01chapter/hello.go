// /Users/sarveshkumar/Practice/golang/01chapter/hello.go
package main

import (
	"fmt"
	"hello/com/api"
)

// main is the entry point for the program.
// It prints a welcome message and calls functions from the api package.
func main() {
	//fmt.Println("Welcome go language")
	//api.Test()
	//api.Sub()
	message := api.Greetings("shiv")
	fmt.Println(message)
}

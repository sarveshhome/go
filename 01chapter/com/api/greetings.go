package api

import "fmt"

// Greetings returns a personalized greeting message.
//
// Args:
//
//	name: The name to include in the greeting.
//
// Returns:
//
//	A string containing the greeting message.
func Greetings(name string) string {
	//return a greeting that embeds
	message := fmt.Sprintf("Hi, %v Welcome!", name)
	return message
}

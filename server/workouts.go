package main

import "fmt"

// Workouts returns workouts for the user
func workouts(user string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Here are your workouts!", user)
	return message
}

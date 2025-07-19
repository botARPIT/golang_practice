package main

import (
	"fmt"
	// "example.com/greetings"
)

// func main() {
// 	message := greetings.Hello("Yadav")
// 	fmt.Println(message)

// }

func main() {
	fmt.Println(greeting(""))
}

func greeting(name string) string {
	if name == "" {
		message := fmt.Sprint("Hello! Stranger")
		return message
	} else {
		message := fmt.Sprintf("Hello! %v", name)
		return message
	}
}

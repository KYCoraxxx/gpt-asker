package main

import (
	"fmt"
	"gpt-asker/session"
)

func main() {
	fmt.Println("Welcome! This is the ChatGPT Assistant!")

	session.Spin()
}

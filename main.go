package main

import (
	"fmt"
	"gpt-asker/client"
)

func main() {
	fmt.Println("Welcome! This is the ChatGPT Assistant!")

	fmt.Println("Creating OpenAI Client and Construct Session ...")

	session := client.NewSession()

	session.Spin()
}

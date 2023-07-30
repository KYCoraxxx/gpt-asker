package main

import (
	"fmt"
	"gpt-asker/client"
	"gpt-asker/config"
	"log"
	"os"
)

func printHelloMessage() {
	fmt.Println("Welcome! This is the ChatGPT Assistant!")

	fmt.Println("Creating OpenAI Client and Construct Session ...")
}

func main() {
	args := os.Args

	if err := config.ResolveArgument(args); err != nil {
		log.Fatal(err)
		return
	}

	printHelloMessage()

	session := client.NewSession()

	session.Spin()
}

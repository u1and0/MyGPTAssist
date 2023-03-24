package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/otiai10/openaigo"
)

const (
	PROMPT = "あなた"
	NAME   = "ChatGPT"
)

type AI struct {
	Client openaigo.Client
}

func NewAI(c openaigo.Client) *AI {
	return &AI{c}
}

func (ai *AI) ask(query string) {
	if query == "" {
		fmt.Printf("%s: ", PROMPT)
		scanner := bufio.NewScanner(os.Stdin)
		for {
			if !scanner.Scan() {
				break
			}
		}
		query = scanner.Text()
	}
	fmt.Printf("%s\n", query)
	ai.ask("")
}

func main() {
	client := openaigo.NewClient(os.Getenv("CHATGPT_API_KEY"))
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Printf("%s: ", PROMPT)
		if !scanner.Scan() {
			break
		}
		query := scanner.Text()
		request := openaigo.ChatCompletionRequestBody{
			Model: "gpt-3.5-turbo",
			Messages: []openaigo.ChatMessage{
				{Role: "user", Content: query},
			},
		}
		ctx := context.Background()
		response, err := client.Chat(ctx, request)
		content := response.Choices[0].Message.Content
		if err != nil {
			log.Fatalln(response, err)
			os.Exit(1)
		}
		message := fmt.SPrintf("%s: %s\n", NAME, content)
		for _, c := range message {
			fmt.Print(c)
		}
		// if err != nil {
		// 	fmt.Println(err)
		// 	continue
		// }
		// fmt.Println()
	}
}

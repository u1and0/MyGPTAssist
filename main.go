package main

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/otiai10/openaigo"
)

func main() {
	client := openaigo.NewClient(os.Getenv("CHATGPT_API_KEY"))
	scanner := bufio.NewScanner(os.Stdin)
	var query string
	for scanner.Scan() {
		query += scanner.Text()
	}
	fmt.Println(query)
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
		fmt.Println(content)
		os.Exit(1)
	}
	fmt.Println(content)
}

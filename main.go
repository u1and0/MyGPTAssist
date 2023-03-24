package main

import (
	"context"
	"fmt"
	"os"

	"github.com/otiai10/openaigo"
)

func main() {
	client := openaigo.NewClient(os.Getenv("CHATGPT_API_KEY"))
	request := openaigo.ChatCompletionRequestBody{
		Model: "gpt-3.5-turbo",
		Messages: []openaigo.ChatMessage{
			{Role: "user", Content: "Hello!"},
		},
	}
	ctx := context.Background()
	response, err := client.Chat(ctx, request)
	content := response.Choices[0].Message.Content
	if err != nil {
		fmt.Println(content)
	}
	fmt.Println(content)
}

package main

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/otiai10/openaigo"
	"github.com/stretchr/testify/require"
)

const (
	ENDPOINT = "https://api.openai.com/v1/chat/completions"
	MODEL    = "gpt-3.5-turbo"
)

var (
	apiKey string
)

type AI struct {
	Name        string
	MaxTokens   int
	Temperature float64
	SystemRole  string
	Filename    string
	// 長期記憶
	// Gist
	ChatSummary string
	Sound       bool
}

type (
	Message struct {
		Role    string
		context string
	}
	Data struct {
		Model       string
		Temperature float64
		MaxTokens   int
		Messages    []Message
	}
)

func (*AI) Post(data:Data) string {
	require, err := http.NewRequest("POST", ENDPOINT, bytes.NewBuffer([]byte(data)))
	header = map[string]string{
		"Content-Type":  "application/json",
		"Authorization": fmt.Sprint("Bearer %s", apiKey),
	}
	for k, v := range header {
		req.Header.Set(k, v)
	}
	return ""
}

func newAI() AI {
	// client := openaigo.NewClient(apiKey)
	return AI{
		Name:        "ChatGPT",
		MaxTokens:   1000,
		Temperature: 0.2,
		SystemRole:  "さっきの話の内容を聞れたら、話をまとめてください。",
		Filename:    "chatgpt-assistant.txt",
	}
}

func init() {
	apiKey := os.Getenv("CHATGPT_API_KEY")
	if apiKey == "" {
		log.Fatalln("Missing API KEY")
	}
}

func main() {
	ai := newAI()
	// Read from stdin
	scanner := bufio.NewScanner(os.Stdin)
	var query string
	for scanner.Scan() {
		query += scanner.Text()
	}
	// Send request ChatGPT
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

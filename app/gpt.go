package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

type Chat struct {
	Engine   string    `json:"model"`
	Messages []Message `json:"messages"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type CompletionResponse struct {
	ID      string          `json:"id"`
	Object  string          `json:"object"`
	Created int64           `json:"created"`
	Choices []ChoiceMessage `json:"choices"`
	Usage   Usage           `json:"usage"`
}

type ChoiceMessage struct {
	Index        int     `json:"index"`
	Message      Message `json:"message"`
	FinishReason string  `json:"finish_reason"`
}

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

func (c *Chat) gptChat(completionRequest *Chat) []byte {
	body, err := json.Marshal(completionRequest)
	if err != nil {
		log.Fatal("Error preparing request:", err)
	}

	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Fatal("OPENAI_API_KEY environment variable not set")
	}

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(body))
	if err != nil {
		log.Fatal("Error creating request:", err)
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "OpenAI-GPT-3-Example")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error making request:", err)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response body:", err)
	}

	// Print an info message summarizing the response
	completionResponse := c.processResponse(respBody)
	tokensUsed := completionResponse.Usage.TotalTokens
	responseTime := float64(time.Since(time.Unix(completionResponse.Created, 0)).Milliseconds()) / 1000.0
	log.Infof("Response summary: Tokens used - %d, Response time - %.2fs", tokensUsed, responseTime)

	return respBody
}

func (c *Chat) processResponse(respBody []byte) *CompletionResponse {
	completionResponse := &CompletionResponse{}
	err := json.Unmarshal(respBody, completionResponse)
	if err != nil {
		log.Fatal("Error parsing response:", err)
	}

	log.Debug(string(respBody))

	return completionResponse
}

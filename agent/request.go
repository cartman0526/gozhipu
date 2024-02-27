package agent

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Tool struct {
	Type     string   `json:"type"`
	Function Function `json:"function"`
}

type Function struct {
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Parameters  Parameters `json:"parameters"`
}

type Parameters struct {
	Type       string               `json:"type"`
	Properties map[string]Parameter `json:"properties"`
	Required   []string             `json:"required"`
}

type Parameter struct {
	Description string `json:"description"`
	Type        string `json:"type"`
}

type ToolMessage struct {
	Role       string `json:"role"`
	Content    string `json:"content"`
	ToolCallID string `json:"tool_call_id"`
}

type RequestBody struct {
	Model       string        `json:"model"`
	Messages    []ToolMessage `json:"messages"`
	Tools       []Tool        `json:"tools"`
	Temperature string        `json:"temperature"`
	MaxTokens   string        `json:"max_tokens"`
}

func Request(token, message, toolmessage, toolID string) *ToolResponse {
	tool := NewGetOSINFO()
	fmt.Println(tool)
	req := &RequestBody{
		Model:       "glm-4",
		Temperature: "0.1",
		MaxTokens:   "1024",
		Messages: []ToolMessage{
			{
				Role:    "user",
				Content: message,
			},
		},
		Tools: []Tool{
			tool,
		},
	}
	if toolmessage != "" {
		req.Messages = append(req.Messages, ToolMessage{
			Role:       "tool",
			Content:    toolmessage,
			ToolCallID: toolID,
		})

	}
	reqBody, err := json.Marshal(req)
	if err != nil {
		log.Fatal(err)
	}
	client, err := http.NewRequest("POST", "https://open.bigmodel.cn/api/paas/v4/chat/completions", bytes.NewBuffer(reqBody))
	if err != nil {
		log.Fatal(err)
	}
	client.Header.Set("Authorization", token)
	client.Header.Set("Content-Type", "application/json")
	httpClient := &http.Client{}
	resp, err := httpClient.Do(client)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	var toolResponse ToolResponse
	Body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(Body, &toolResponse)
	return &toolResponse
}

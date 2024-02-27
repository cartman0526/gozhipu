package agent

type ToolResponse struct {
	ID        string `json:"id"`
	Model     string `json:"model"`
	RequestID string `json:"request_id"`
	Created   int64  `json:"created"`
	Choices   []struct {
		FinishReason string `json:"finish_reason"`
		Index        int    `json:"index"`
		Message      struct {
			Content   string `json:"content"`
			Role      string `json:"role"`
			ToolCalls []struct {
				ID       string `json:"id"`
				Index    int    `json:"index"`
				Type     string `json:"type"`
				Function struct {
					Arguments string `json:"arguments"`
					Name      string `json:"name"`
				} `json:"function"`
			} `json:"tool_calls"`
		} `json:"message"`
	} `json:"choices"`
	Usage struct {
		CompletionTokens int `json:"completion_tokens"`
		PromptTokens     int `json:"prompt_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}

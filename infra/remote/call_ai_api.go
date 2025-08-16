package remote

import (
	"encoding/json"
	"fmt"
)

type AskMsg struct {
	System string `json:"system"`
	User   string `json:"user"`
}

func CallTongYi(ask AskMsg) ChatResponse {
	url := "https://dashscope.aliyuncs.com/compatible-mode/v1/chat/completions"
	apiKey := "sk-23a60fd5c014408b9eeb2579993bf5a9"
	// 构建请求体
	requestBody := ChartRequest{
		// 模型列表：https://help.aliyun.com/zh/model-studio/getting-started/models
		Model: "qwen-plus",
		Messages: []Message{
			{
				Role:    "system",
				Content: ask.System,
			},
			{
				Role:    "user",
				Content: ask.User,
			},
		},
	}
	respContent, err := post(url, requestBody, map[string]string{
		"Authorization": "Bearer " + apiKey,
		"Content-Type":  "application/json",
	})
	if err != nil {
		panic(err)
	}
	var response ChatResponse
	err = json.Unmarshal(respContent, &response)
	if err != nil {
		panic(err)
	}
	return response
}

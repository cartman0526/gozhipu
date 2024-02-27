package main

import (
	"encoding/json"
	"fmt"
	"github.com/aitangmii/gozhipu/agent"
	"github.com/aitangmii/gozhipu/pkg/apikey"
)

func main() {
	token := apikey.CreateToken("35e41f4e1c00eea4d9be69e15af364fc.mXqG68Gnyyqy4AtY")
	response := agent.Request(token, "我想获取我服务器的内存指标", "", "")
	fmt.Println(response.Choices[0].Message.ToolCalls[0].Function.Arguments)
	args := response.Choices[0].Message.ToolCalls[0].Function.Arguments
	// 我想将字符串{"ostype":"cpu"} 转换成map
	type SystemType struct {
		Ostype string `json:"ostype"`
	}
	var info SystemType
	json.Unmarshal([]byte(args), &info)

	toolResult := agent.GetOSINFO(info.Ostype)
	jsonData, err := json.Marshal(toolResult)
	if err != nil {
		panic(err)
	}
	id := response.Choices[0].Message.ToolCalls[0].ID
	response = agent.Request(token, "我想获取我linux服务器的内存当前系统状态", string(jsonData), id)
	fmt.Println(response.Choices[0].Message.Content)
}

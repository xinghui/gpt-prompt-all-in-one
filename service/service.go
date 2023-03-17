package service

import (
	"encoding/json"
	"fmt"

	"github.com/xinghui/gpt-prompt-all-in-one/dao/openai"
	"github.com/xinghui/gpt-prompt-all-in-one/vo"
)

const (
	model = "gpt-3.5-turbo"
)

func AskGPT(prompt []vo.Message, content string) {

	var message []vo.Message
	if prompt != nil {
		message = append(message, prompt...)
	}
	message = append(message, vo.Message{Role: vo.RoleUser, Content: content})

	data := vo.OpenApiReq{
		Model:    model,
		Messages: message,
	}
	_, err := openai.SendReq(data)
	if err != nil {
		return
	}
}

func prettyPrint(resp interface{}) string {
	respJson, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		fmt.Printf("err:%+v|resp:%+v\n", err, resp)
		return ""
	}
	return fmt.Sprintf("\n%+v\n", string(respJson))
}

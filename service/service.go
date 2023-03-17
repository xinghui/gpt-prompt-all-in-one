package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/xinghui/gpt-prompt-all-in-one/clog"
	"github.com/xinghui/gpt-prompt-all-in-one/config"
	"github.com/xinghui/gpt-prompt-all-in-one/vo"
)

const (
	Url = "https://api.openai.com/v1/chat/completions"
)

func AskGPT(prompt []vo.Message, content string) {

	var message []vo.Message
	if prompt != nil {
		message = append(message, prompt...)
	}
	message = append(message, vo.Message{Role: vo.RoleUser, Content: content})

	data := vo.OpenApiReq{
		Model:    "gpt-3.5-turbo",
		Messages: message,
	}
	_, err := sendReq(data)
	if err != nil {
		return
	}
}

func sendReq(data vo.OpenApiReq) (*vo.OpenApiResp, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("Marshal:%+v\n", err)
		return nil, err
	}
	//fmt.Printf("Req: " + prettyPrint(data))

	req, err := http.NewRequest("POST", Url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("NewRequest:%+v\n", err)
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", config.GetToken()))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Post:%+v\n", err)
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("ReadAll:%+v\n", err)
		return nil, err
	}
	var r vo.OpenApiResp
	err = json.Unmarshal(body, &r)
	if err != nil {
		fmt.Printf("Unmarshal:%+v\n", err)
		return nil, err
	}
	//fmt.Printf("Resp: " + string(body))
	//fmt.Printf("Resp: " + prettyPrint(r))
	for _, choice := range r.Choices {
		clog.Tips2(choice.Message.Content)
	}
	return &r, nil
}

func prettyPrint(resp interface{}) string {
	respJson, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		fmt.Printf("err:%+v|resp:%+v\n", err, resp)
		return ""
	}
	return fmt.Sprintf("\n%+v\n", string(respJson))
}

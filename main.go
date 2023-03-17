package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/xinghui/gpt-prompt-all-in-one/config"
	"io"
	"log"
	"net/http"
	"os"
)

const (
	Url           = "https://api.openai.com/v1/chat/completions"
	RoleUser      = "user"
	RoleSystem    = "system"
	RoleAssistant = "assistant"
)

func main() {

	config.Init()

	app := &cli.App{
		Name:  "GPT Prompts all-in-one",
		Usage: "sss",
		Action: func(context *cli.Context) error {
			fmt.Printf("!!!!!!!")
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func gpt() {

	var result []string
	data := OpenApiReq{
		Model: "gpt-3.5-turbo",
		//Model: "gpt-3.5-turbo-0301",
		//Model: "gpt-4-32k",
		Messages: []Message{
			{Role: RoleSystem, Content: "以下说的所有的话你都翻译成英文"},
			{Role: RoleUser, Content: "你好，你叫什么名字"},
		},
	}
	_, err := sendReq(data)
	if err != nil {
		return
	}

	//for {
	//	resp, err := sendReq(data)
	//	if err != nil {
	//		return
	//	}
	//	if len(resp.Choices) == 0 {
	//		break
	//	}
	//	result = append(result, resp.Choices[0].Message.Content)
	//	if resp.Choices[0].FinishReason == "stop" {
	//		data = OpenApiReq{
	//			Model: "gpt-3.5-turbo",
	//			Messages: []Message{
	//				{Role: resp.Choices[0].Message.Role, Content: resp.Choices[0].Message.Content},
	//				{Role: RoleUser, Content: "继续"},
	//			},
	//		}
	//		continue
	//	}
	//}

	for i, s := range result {
		fmt.Printf("[%d]%s", i, s)
	}
}

func sendReq(data OpenApiReq) (*OpenApiResp, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("Marshal:%+v\n", err)
		return nil, err
	}
	fmt.Printf("Req: " + prettyPrint(data))

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
	var r OpenApiResp
	err = json.Unmarshal(body, &r)
	if err != nil {
		fmt.Printf("Unmarshal:%+v\n", err)
		return nil, err
	}
	fmt.Printf("Resp: " + string(body))
	fmt.Printf("Resp: " + prettyPrint(r))
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

type OpenApiReq struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}
type OpenApiResp struct {
	Id      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Index        int    `json:"index"`
		FinishReason string `json:"finish_reason"`
		Message      struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Error struct {
		Message string      `json:"message"`
		Type    string      `json:"type"`
		Param   interface{} `json:"param"`
		Code    string      `json:"code"`
	} `json:"error"`
}

package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/urfave/cli/v2"
	"github.com/xinghui/gpt-prompt-all-in-one/config"
	"github.com/xinghui/gpt-prompt-all-in-one/input"
	"github.com/xinghui/gpt-prompt-all-in-one/service"
	"github.com/xinghui/gpt-prompt-all-in-one/vo"
)

func main() {

	config.Init()

	app := &cli.App{
		Name:     "GPT Prompts all-in-one",
		Usage:    ":q 退出",
		Commands: commands,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

var commands = []*cli.Command{
	{
		Name:    "translate",
		Aliases: []string{"t"},
		Usage:   "英翻中",
		Action: func(cCtx *cli.Context) error {
			triggerInput(vo.BuildPrompt(vo.PromptTransfer))
			return nil
		},
	},
	{
		Name:    "grammar",
		Aliases: []string{"g"},
		Usage:   "语法检查",
		Action: func(cCtx *cli.Context) error {
			triggerInput(vo.BuildPrompt(vo.PromptGrammar))
			return nil
		},
	},
}

func triggerInput(prompt []vo.Message) {
	fmt.Println()
	input.Input("请输入：", true, func(input string) error {
		if strings.Index(input, ":q") >= 0 {
			return nil
		}
		service.AskGPT(prompt, input)
		triggerInput(prompt)
		return nil
	})
}

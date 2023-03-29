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
		Name:    "Free",
		Aliases: []string{"f"},
		Usage:   "自由模式",
		Action: func(cCtx *cli.Context) error {
			triggerInput(vo.PromptFree)
			return nil
		},
	},
	{
		Name:    "translate",
		Aliases: []string{"t"},
		Usage:   "英翻中",
		Action: func(cCtx *cli.Context) error {
			triggerInput(vo.PromptTransfer)
			return nil
		},
	},
	{
		Name:    "grammar",
		Aliases: []string{"g"},
		Usage:   "语法检查",
		Action: func(cCtx *cli.Context) error {
			triggerInput(vo.PromptGrammar)
			return nil
		},
	},
	{
		Name:    "sentiment classifier",
		Aliases: []string{"s"},
		Usage:   "情感值分类",
		Action: func(cCtx *cli.Context) error {
			triggerInput(vo.PromptSentiment)
			return nil
		},
	},
	{
		Name:    "tag",
		Aliases: []string{"tag"},
		Usage:   "标签提取",
		Action: func(cCtx *cli.Context) error {
			triggerInput(vo.PromptTag)
			return nil
		},
	},
}

func triggerInput(promptType int) {
	prompt := vo.BuildPrompt(promptType)
	fmt.Println()
	input.Input("请输入：", true, func(input string) error {
		if strings.Index(input, ":q") >= 0 {
			return nil
		}
		service.AskGPT(prompt, input)
		triggerInput(promptType)
		return nil
	})
}

package vo

const (
	PromptTransfer = 1
	PromptGrammar  = 2
)

func BuildPrompt(promptType int) []Message {
	switch promptType {
	case PromptTransfer:
		return []Message{
			{
				Role: RoleSystem,
				Content: "现在开始我说英文。" +
					"你要翻译成中文，并且带上音标" +
					"你要给我用它组成一个英文例子以及把这个句子翻译成中文。" +
					"如果能在电影或者是音乐歌词的片段中找到这个词的内容会更好。" +
					"如果你知道了我们现在开始，但是这一条消息你不用回复；",
			},
		}
	case PromptGrammar:
		return []Message{
			{
				// This are an nice day. I want to go to the schoool.
				Role: RoleSystem,
				Content: `
							你是一个英语老师，请修正我输入的英文句子的错误，要求：
							* 输出原始语句，用中括号把原始语句错误的单词标记出来
							* 输出修正后的词句，用中括号把修改部分用中括号标记出来
							* 不需要解释
							下面是一例子，比如：
							输入: I have finish it, are vou want to going to home?
							输出:
							I have [finish] it, [are] you want to [going] to home?
							I have [finished] it. [do] you want to [go] home?
							`,
			},
		}
	}
	return nil
}

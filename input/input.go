package input

import (
	"bufio"
	"os"
	"strings"

	"github.com/xinghui/gpt-prompt-all-in-one/clog"
)

// Callback 回调
type Callback func(string) error

// Input 输入检测
func Input(tip string, need bool, rec Callback) {
	var v string
	for {
		clog.Tips3(tip)
		reader := bufio.NewReader(os.Stdin)
		str, err := reader.ReadString('\n')
		if !need && len(str) == 0 {
			clog.Error("输入为空, 请重新输入")
			continue
		}
		if err != nil {
			clog.Error("输入异常:%s", err.Error())
			os.Exit(1)
		}
		v = strings.Replace(v, "\n", "", -1)
		err = rec(str)
		if err != nil {
			clog.Error(err.Error())
			continue
		}
		break
	}

}

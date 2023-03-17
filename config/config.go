package config

import (
	"fmt"
	"os"
)

var token string

func Init() {
	var err error
	token, err = readTextFromFile("token")
	if err != nil {
		fmt.Printf("config.Init|%s", err)
		return
	}
	fmt.Println("------------------------------------------")
	fmt.Printf("token: %s******************%s\n", token[:5], token[len(token)-3:])
	fmt.Println("==========================================")
}

func GetToken() string {
	return token
}

// 定義一個函數，參數為文件名，返回值為文本內容和錯誤信息
func readTextFromFile(filename string) (string, error) {
	// 使用ioutil.ReadFile函數讀取文件，返回字節切片和錯誤信息
	data, err := os.ReadFile(filename)
	// 如果有錯誤，則返回空字符串和錯誤信息
	if err != nil {
		return "", err
	}
	// 將字節切片轉換為字符串，並返回字符串和nil表示沒有錯誤
	return string(data), nil
}

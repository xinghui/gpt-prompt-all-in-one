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
	fmt.Printf("API Key: %s******************%s\n", token[:5], token[len(token)-3:])
	fmt.Println("==========================================")
}

func GetToken() string {
	return token
}

func readTextFromFile(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

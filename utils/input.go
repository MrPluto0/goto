package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func InputStr(prompt string) string {
	var input string
	fmt.Println(prompt)
	fmt.Scan(&input)
	return input
}

func InputLn(prompt string) string {
	fmt.Println(prompt)
	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	return strings.Trim(input, "\r\n ")
}

func InputBool(prompt string) bool {
	var input string
	fmt.Println(prompt + " [y/n]")
	fmt.Scan(&input)
	return input == "y"
}

package main

import (
	"errors"
	"fmt"
)

type errorString struct {
	text string
}

func (e errorString) Error() string {
	return e.text
}

// New 创建一个自定义错误
func New(s string) error {
	return errorString{text: s}
}

var errorString1 = New("error string1")
var errorString2 = errors.New("error string2")

func main() {
	if errorString1 == New("error string1") {
		fmt.Println("err string1") // 会输出
	}

	if errorString2 == errors.New("error string2") {
		fmt.Println("err string2") // 不会输出
	}
}

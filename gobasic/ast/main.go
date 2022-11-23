package main

import "fmt"

func main() {
	data, err := generate("D:\\project\\pratice\\gobasic\\ast\\user.go")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)
}

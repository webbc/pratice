package main

import (
	"fmt"
	"io/ioutil"
)

func readFile2(name string) ([]byte, error) {
	buf, err := ioutil.ReadFile(name)
	if err != nil {
		return nil, fmt.Errorf("readFile open fail :%w", err)
	}
	return buf, nil
}

func readConfig2() ([]byte, error) {
	name := "/home/root/1.txt"
	config, err := readFile2(name)
	return config, fmt.Errorf("readConfig could not read config : %w", err)
}

func main() {
	_, err := readConfig2()
	if err != nil {
		fmt.Printf("%T | %v\n", err, err)
	}
}

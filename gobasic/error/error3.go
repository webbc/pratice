package main

import (
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
)

func readFile(name string) ([]byte, error) {
	buf, err := ioutil.ReadFile(name)
	if err != nil {
		return nil, errors.Wrap(err, "readFile open fail")
	}
	return buf, nil
}

func readConfig() ([]byte, error) {
	name := "/home/root/1.txt"
	config, err := readFile(name)
	return config, errors.WithMessage(err, "readConfig could not read config")
}

func main() {
	_, err := readConfig()
	if err != nil {
		fmt.Printf("%T | %v\n", errors.Cause(err), errors.Cause(err))
		fmt.Printf("stack:\n%+v", err)
	}
}

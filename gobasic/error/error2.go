package main

import (
	"fmt"
)

type MyError struct {
	s    string
	name string
	line int
}

func (this *MyError) Error() string {
	return fmt.Sprintf("line:%d,name:%s,error msg:%s", this.line, this.name, this.s)
}

func a() error {
	return nil
}

func b() error {
	return a()
}

func main() {
	b()
}

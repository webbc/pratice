package main

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("NOT FOUND")

type FileNotFoundError struct {
	name string
	err  error
}

func (this *FileNotFoundError) Error() string {
	return fmt.Sprintf("name:%s,err:%v", this.name, this.err)
}

func (this *FileNotFoundError) Unwrap() error {
	return this.err
}

func main() {

	err := &FileNotFoundError{"a.txt", ErrNotFound}
	if errors.Is(err, ErrNotFound) {
		fmt.Println("err not found")
	}

	var fErr *FileNotFoundError
	if errors.As(err, &fErr) {
		fmt.Println("file name：" + fErr.name)
	}

}

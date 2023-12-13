package testify

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEqual(t *testing.T) {
	var a = 100
	var b = 200
	assert.Equal(t, a, b, "111111111111111")
}

func TestEqual2(t *testing.T) {
	assert.Contains(t, "Hello World", "World2", "22222222")
}

func TestEqual3(t *testing.T) {
	assert.DirExists(t, "/home/baochao/pratice/testify/01_test.go", "3333333333")
}

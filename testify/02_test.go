package testify

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEqualx(t *testing.T) {
	assertion := assert.New(t)
	assertion.Equal(100, 200, "1111")
}

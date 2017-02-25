package test

import (
	"testing"
	"strings"

	"github.com/issue9/assert"
)

func TestSplitCode(t *testing.T) {
	a := assert.New(t)

	a.Equal(strings.Split("1,1,1,1", ","), []string{"1", "1", "1", "1"})
}

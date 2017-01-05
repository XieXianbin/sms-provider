package sms-provider

import (
	"testing"

	"github.com/issue9/assert"
)

func TestSplitCode(t *testing.T) {
	a := assert.New(t)

	a.Equal(SplitCode("1111"), "1,1,1,1")
	a.Equal(SplitCode("1234"), "1,2,3,4")
}

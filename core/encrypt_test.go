package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatCustomBase(t *testing.T) {
	ast := assert.New(t)

	testCases := []struct {
		Input  int64
		Except string
	}{
		{
			Input:  100354,
			Except: "4pn",
		},
	}
	for _, testCase := range testCases {
		encodeStr := FormatCustomBase(testCase.Input)
		ast.Equal(testCase.Except, encodeStr)
	}
}

func TestParseCustomBase(t *testing.T) {
	ast := assert.New(t)
	testCases := []struct {
		Input  string
		Except int64
	}{
		{
			Input:  "4pn",
			Except: 100354,
		},
	}
	for _, testCase := range testCases {
		decodeInt, err := ParseCustomBase(testCase.Input)
		ast.NoError(err)
		ast.Equal(testCase.Except, decodeInt)
	}
}

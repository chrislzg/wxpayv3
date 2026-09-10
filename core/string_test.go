package core

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBufWriteStringWithLn(t *testing.T) {

	ast := assert.New(t)

	testCases := []struct {
		Input  string
		Except string
	}{
		{
			Input:  "test",
			Except: "test\n",
		},
		{
			Input:  "",
			Except: "\n",
		},
	}

	for _, tCase := range testCases {
		bf := &strings.Builder{}
		bw := bufio.NewWriter(bf)
		err := BufWriteStringWithLn(bw, tCase.Input)
		ast.NoError(err)
		err = bw.Flush()
		ast.NoError(err)
		ast.Equal(bf.String(), tCase.Except)
	}
}

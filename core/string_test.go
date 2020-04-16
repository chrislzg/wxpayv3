package core

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBufWriteWithLn(t *testing.T) {
	ast := assert.New(t)

	testCases := []struct {
		Input  []byte
		Except string
	}{
		{
			Input:  []byte("test"),
			Except: "test\n",
		},
		{
			Input:  []byte{},
			Except: "\n",
		},
	}

	for _, tCase := range testCases {
		bf := &strings.Builder{}
		bw := bufio.NewWriter(bf)
		err := BufWriteWithLn(bw, tCase.Input)
		ast.NoError(err)
		err = bw.Flush()
		ast.NoError(err)
		ast.Equal(bf.String(), tCase.Except)
	}
}

func TestBufWriteByteWithLn(t *testing.T) {
	ast := assert.New(t)

	testCases := []struct {
		Input  byte
		Except string
	}{
		{
			Input:  byte('c'),
			Except: "c\n",
		},
		{
			Input:  byte(65),
			Except: "A\n",
		},
	}

	for _, tCase := range testCases {
		bf := &strings.Builder{}
		bw := bufio.NewWriter(bf)
		err := BufWriteByteWithLn(bw, tCase.Input)
		ast.NoError(err)
		err = bw.Flush()
		ast.NoError(err)
		ast.Equal(bf.String(), tCase.Except)
	}
}

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

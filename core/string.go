package core

import (
	"bufio"
)

func BufWriteWithLn(writer *bufio.Writer, p []byte) error {
	if _, err := writer.Write(p); err != nil {
		return err
	}
	return writer.WriteByte('\n')
}

func BufWriteByteWithLn(writer *bufio.Writer, c byte) error {
	if err := writer.WriteByte(c); err != nil {
		return err
	}
	return writer.WriteByte('\n')
}

func BufWriteStringWithLn(writer *bufio.Writer, s string) error {
	if _, err := writer.WriteString(s); err != nil {
		return err
	}
	return writer.WriteByte('\n')
}

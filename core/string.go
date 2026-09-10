package core

import (
	"bufio"
)

func BufWriteStringWithLn(writer *bufio.Writer, s string) error {
	if _, err := writer.WriteString(s); err != nil {
		return err
	}
	return writer.WriteByte('\n')
}

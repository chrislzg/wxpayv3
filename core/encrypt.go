package core

import (
	"errors"
	"math"
)

const (
	AlgorithmAEADAES256GCM = "AEAD_AES_256_GCM"
)

var (
	baseChar       = "SxC5d2paylJzcseqtbI8QA0FgG4DwLHvWKZhURniE37Pujk1T6XMfOYr9NmVoB"
	ErrInvalidBase = errors.New("invalid base")
	// baseChar的值对应的int值
	charIntValue = map[byte]int64{
		'B': 61,
		'E': 40,
		'J': 10,
		'L': 29,
		'P': 43,
		'D': 27,
		'F': 23,
		'T': 48,
		'k': 46,
		'2': 5,
		'7': 42,
		'm': 58,
		'o': 60,
		'q': 15,
		'x': 1,
		'X': 50,
		'b': 17,
		'u': 44,
		'K': 33,
		'v': 31,
		'O': 53,
		'5': 3,
		'8': 19,
		'h': 35,
		'l': 9,
		'V': 59,
		'e': 14,
		'G': 25,
		'1': 47,
		'p': 6,
		'U': 36,
		'A': 21,
		'I': 18,
		'Y': 54,
		'6': 49,
		'a': 7,
		'n': 38,
		'r': 55,
		'H': 30,
		'M': 51,
		's': 13,
		'C': 2,
		'Q': 20,
		'4': 26,
		'3': 41,
		'9': 56,
		'c': 12,
		'f': 52,
		'i': 39,
		'N': 57,
		'R': 37,
		'W': 32,
		'Z': 34,
		'0': 22,
		'd': 4,
		't': 16,
		'w': 28,
		'z': 11,
		'S': 0,
		'g': 24,
		'j': 45,
		'y': 8,
	}
)

// 十进制 > 定义的进制
func FormatCustomBase(num int64) string {
	var buf [64]byte
	base := int64(len(baseChar))
	index := len(buf) - 1
	for i := index; i >= 0; i-- {
		mod := num % base
		buf[i] = baseChar[mod]
		if num = num / base; num == 0 {
			index = i
			break
		}
	}
	return string(buf[index:])
}

// 自定义进制 > 十进制
func ParseCustomBase(str string) (int64, error) {
	if str == "" {
		return -1, ErrInvalidBase
	}
	base := len(baseChar)
	var ret int64
	lenStr := len(str)
	for i := lenStr - 1; i >= 0; i-- {
		val, ok := charIntValue[str[i]]
		if !ok {
			return -1, ErrInvalidBase
		}
		ret += int64(math.Pow(float64(base), float64(lenStr-1-i))) * val
	}
	return ret, nil
}

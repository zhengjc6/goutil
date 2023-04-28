package strtool

import (
	"math/rand"
	"strings"
	"time"

	"golang.org/x/text/encoding/simplifiedchinese"
)

type Charset string

const (
	UTF8    = Charset("UTF-8")
	GB18030 = Charset("GB18030")
)

func ConvertByte2String(byte []byte, charset Charset) string {
	var str string
	switch charset {
	case GB18030:
		var decodeBytes, _ = simplifiedchinese.GB18030.NewDecoder().Bytes(byte)
		str = string(decodeBytes)
	case UTF8:
		fallthrough
	default:
		str = string(byte)
	}
	return str
}

// FirstUpper
func FirstUpper(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

// FirstLower
func FirstLower(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToLower(s[:1]) + s[1:]
}

var letters = []rune{'2', '3', '4', '5', '6', '7', '8', '9',
	'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
}
var letterlen = len(letters)
var seed int64 = 0

func RandomString(len int) string {
	if seed == 0 {
		seed = time.Now().UnixNano()
		rand.Seed(seed)
	}
	tb := []rune{}
	for i := 0; i < len; i++ {
		tb = append(tb, letters[rand.Int31n(int32(letterlen))])
	}
	return string(tb)
}

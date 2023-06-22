package lang

import (
	"bytes"
	"strings"
)

// Hex 用字符串的形式表示一个十六进制数
type Hex string

func (h Hex) String() string {
	return string(h)
}

// Bytes 转换为字节切片
func (h Hex) Bytes() []byte {
	count := 0
	array := []rune(h)
	builder := &bytes.Buffer{}
	n1 := 0
	for _, ch := range array {
		n := 0
		if ('0' <= ch) && (ch <= '9') {
			n = int(ch - '0')
		} else if ('a' <= ch) && (ch <= 'f') {
			n = int(ch-'a') + 0x0a
		} else if ('A' <= ch) && (ch <= 'F') {
			n = int(ch-'A') + 0x0a
		} else {
			continue
		}
		count++
		if (count & 0x01) == 1 {
			n1 = n
		} else {
			n12 := (n1 << 4) | n
			builder.WriteByte(byte(n12))
		}
	}
	return builder.Bytes()
}

// HexFromBytes 把字节切片转化为 Hex
func HexFromBytes(data []byte) Hex {
	const chstr = "0123456789abcdef"
	chset := []rune(chstr)
	builder := &strings.Builder{}
	for _, b := range data {
		n := 0x00ff & int(b)
		builder.WriteRune(chset[0x0f&(n>>4)])
		builder.WriteRune(chset[0x0f&n])
	}
	str := builder.String()
	return Hex(str)
}

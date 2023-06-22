package lang

import "encoding/base64"

// Base64 以字符串的形式表示一个 base64 格式的数据
type Base64 string

func (b Base64) String() string {
	return string(b)
}

// Bytes 转化为字节切片
func (b Base64) Bytes() []byte {
	str := b.String()
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil || data == nil {
		data = make([]byte, 0)
	}
	return data
}

// Base64FromBytes 把字节切片编码为 base64 字符串
func Base64FromBytes(data []byte) Base64 {
	str := base64.StdEncoding.EncodeToString(data)
	return Base64(str)
}

package lang

// ByteSlice 表示一个字节串切片
type ByteSlice []byte

// Bytes 转换为字节切片
func (ba ByteSlice) Bytes() []byte {
	if ba == nil {
		ba = make([]byte, 0)
	}
	return ba
}

func (ba ByteSlice) String() string {
	hex := ba.Hex()
	return hex.String()
}

// Hex 转换为十六进制字符串
func (ba ByteSlice) Hex() Hex {
	return HexFromBytes(ba)
}

// Base64 转换为 Base64 字符串
func (ba ByteSlice) Base64() Base64 {
	return Base64FromBytes(ba)
}

// NewByteSlice 新建一个字节切片
func NewByteSlice(data []byte) ByteSlice {
	if data == nil {
		data = make([]byte, 0)
	}
	return data
}

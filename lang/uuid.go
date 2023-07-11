package lang

import "strings"

// UUID like 'xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx'
type UUID string

func (u UUID) String() string {
	return string(u)
}

// Bytes 转化为字节串形式
func (u UUID) Bytes() []byte {
	str := u.String()
	hex := Hex(str)
	return hex.Bytes()
}

// Normalize 标准化 UUID
func (u UUID) Normalize() UUID {
	const wantSize = 16
	raw := u.Bytes()
	haveSize := len(raw)
	if haveSize < wantSize {
		addition := make([]byte, wantSize-haveSize)
		raw = append(raw, addition...)
	} else if haveSize > wantSize {
		raw = raw[0:wantSize]
	}
	hex := HexFromBytes(raw)
	str := hex.String()
	b := &strings.Builder{}
	i := 0
	sizelist := []int{8, 4, 4, 4}
	for _, size := range sizelist {
		part := str[i : i+size]
		b.WriteString(part)
		b.WriteString("-")
		i += size
	}
	b.WriteString(str[i:])
	str = b.String()
	return UUID(str)
}

// NewUUID 新建UUID
func NewUUID(data []byte) UUID {
	hex := HexFromBytes(data)
	u := UUID(hex)
	return u.Normalize()
}

// ParseUUID 解析UUID
func ParseUUID(str string) UUID {
	u := UUID(str)
	return u.Normalize()
}

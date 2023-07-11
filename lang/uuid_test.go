package lang

import (
	"bytes"
	"testing"
)

func TestUUID(t *testing.T) {
	raw := bytes.Buffer{}
	for b := 123; b < 160; b++ {
		data := raw.Bytes()
		hex := HexFromBytes(data)
		uuid := NewUUID(data)
		t.Logf("make uuid:%s from hex:%s", uuid.String(), hex.String())
		raw.WriteByte(byte(b))
	}
}

package lang

import (
	"bytes"
	"testing"
)

func TestBytes(t *testing.T) {

	list := make([]string, 0)
	list = append(list, "")
	list = append(list, "a")
	list = append(list, "a1")
	list = append(list, "a1E")
	list = append(list, "a1E-b")
	list = append(list, "a1Ebc")

	for i, str := range list {

		t.Logf("test base64[%d]: %s", i, str)

		hex := Hex(str)
		bs1 := NewByteSlice(hex.Bytes())

		hex1 := bs1.Hex()
		bas1 := bs1.Base64()
		ba1 := bs1.Bytes()

		bs2 := NewByteSlice(ba1)

		ba2 := bs2.Bytes()
		ba3 := hex1.Bytes()
		ba4 := bas1.Bytes()

		str1 := bs1.String()
		str2 := bs2.String()
		t.Logf("    bstr1 = %s", str1)
		t.Logf("    bstr2 = %s", str2)

		if !bytes.Equal(ba1, ba2) {
			t.Log("ba1 != ba2")
		}
		if !bytes.Equal(ba1, ba3) {
			t.Log("ba1 != ba3")
		}
		if !bytes.Equal(ba1, ba4) {
			t.Log("ba1 != ba4")
		}
	}

	t.Log("done")
}

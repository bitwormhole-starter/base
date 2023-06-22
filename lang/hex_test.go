package lang

import (
	"bytes"
	"testing"
)

func TestHex(t *testing.T) {

	list := make([]string, 0)
	list = append(list, "")
	list = append(list, "a")
	list = append(list, "a1")
	list = append(list, "a1E")
	list = append(list, "a1E-b")
	list = append(list, "a1Ebc")
	list = append(list, "a1Ebc7")

	for i, str := range list {

		t.Logf("test hex[%d]: %s", i, str)

		hex0 := Hex(str)
		ba0 := hex0.Bytes()
		hex1 := HexFromBytes(ba0)
		ba1 := hex1.Bytes()

		str0 := hex0.String()
		str1 := hex1.String()

		t.Logf("    str0 = %s", str0)
		t.Logf("    str1 = %s", str1)

		if !bytes.Equal(ba1, ba0) {
			t.Log("ba1 != ba0")
		}
	}

	t.Log("done")
}

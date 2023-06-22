package lang

import (
	"bytes"
	"testing"
)

func TestBase64(t *testing.T) {

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
		bstr1 := Base64FromBytes(hex.Bytes())
		ba1 := bstr1.Bytes()
		bstr2 := Base64FromBytes(ba1)
		ba2 := bstr2.Bytes()

		t.Logf("    bstr1 = %s", bstr1)
		t.Logf("    bstr2 = %s", bstr2)

		if !bytes.Equal(ba1, ba2) {
			t.Log("ba1 != ba2")
		}
	}

	t.Log("done")
}

package context2

import (
	"context"
	"testing"
)

func TestValues(t *testing.T) {

	c1 := context.Background()

	kv := SetupContext(c1)
	kv.SetValue("a", "1")
	kv.SetValue("b", "2")
	kv.SetValue("c", "3")

	c2 := kv.Context()
	kv, err := GetValues(c2)
	if err != nil {
		t.Error(err)
	}

	kv.GetValue("x")
}

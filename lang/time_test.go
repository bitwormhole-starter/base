package lang

import (
	"testing"
)

func TestTime(t *testing.T) {

	now := Now()

	list := make([]Time, 0)
	list = append(list, -1)
	list = append(list, 0)
	list = append(list, 1)
	list = append(list, 100000000)
	list = append(list, now)

	for i, t1 := range list {

		n1 := t1.Int()
		s1 := t1.String()
		tt1 := t1.Time()

		t.Logf("test time[%d]: %d", i, n1)

		t2 := NewTime(tt1)
		s2 := t2.String()
		n2 := t2.Int()

		t.Logf("    s1 = %s", s1)
		t.Logf("    s2 = %s", s2)

		if n1 != n2 {
			t.Log("n1 != n2")
		}
	}

	t.Log("done")
}

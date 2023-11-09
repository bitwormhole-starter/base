package lang

import (
	"testing"
	"time"
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

func TestTimeDur(t *testing.T) {

	t1 := Now()
	time.Sleep(time.Second * 2)
	t2 := Now()

	d := t2.Sub(t1)
	ds := NewSeconds(d)
	dms := NewMilliseconds(d)

	t3 := t1.Add(ds.Duration())
	t4 := t1.Add(dms.Duration())

	t.Logf("t1  = %d", t1)
	t.Logf("t2  = %d", t2)
	t.Logf("t3  = %d", t3)
	t.Logf("t4  = %d", t4)

	t.Logf("d(duration) = %d", d)
	t.Logf("d(s)        = %d", ds)
	t.Logf("d(ms)       = %d", dms)
}

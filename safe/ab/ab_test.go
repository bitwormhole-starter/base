package ab

import (
	"testing"

	"github.com/starter-go/base/safe"
)

func TestModeAB(t *testing.T) {

	ma := safe.Fast()
	mb := safe.Safe()
	mx := New(ma, mb)
	l := mx.NewLock()

	mx.UseModeA()
	l.Lock()
	l.Unlock()

	mx.UseModeB()
	l.Lock()
	l.Unlock()
}

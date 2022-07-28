package dict

import (
	"testing"
)

func TestDictList(t *testing.T) {
	l := NewDictList()
	t.Log(l.len())
	l.pushBack(NewDict())
	t.Log(l.len())
}

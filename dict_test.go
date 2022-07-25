package dict

import (
	"fmt"
	"testing"
)

func TestDict(t *testing.T) {
	d := NewDictWithObj(map[string]any{
		"a": map[string]any{"a": "c"},
	})
	d.Set("b", "xxx")
	d.Set("a", 22.343)
	fmt.Println(d.GetAny("a"))
}

type Test struct {
	A string
	B int
}

func TestNewDictWithObj(t *testing.T) {
	d := NewDictWithObj(Test{A: "x"})
	fmt.Println(d)
}

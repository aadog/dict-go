package dict

import (
	"fmt"
	"testing"
)

func TestDict(t *testing.T) {
	//d := Dict{}
	//d.Set("b", "xxx")
	//d.Set("a", 22.343)
	//fmt.Println(d.GetString("a"))
	d := NewDictWithMap(map[string]any{
		"a": map[string]any{"a": "c"},
	})
	fmt.Println(d.GetAny("a"))
}

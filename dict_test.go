package dict

import (
	"fmt"
	"testing"
)

func TestDict(t *testing.T) {
	d := NewDictWithMap(map[string]any{
		"a": map[string]any{"a": "c"},
	})
	d.Set("b", "xxx")
	d.Set("a", 22.343)
	fmt.Println(d.GetAny("a"))
}

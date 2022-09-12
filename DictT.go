//go:build v && go1.18

package dict

func DictGetSlice[T any](dict *Dict, key string) []T {
	v := d.GetAny(key)
	if v == nil {
		return nil
	}
	r := make([]T, 0)
	vf := reflect.ValueOf(v)
	if vf.Kind() != reflect.Slice {
		return r
	}
	for i := 0; i < vf.Len(); i++ {
		r = append(r, NewDictWithObj(vf.Index(i).Interface()))
	}
	return r
}

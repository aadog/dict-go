package dict

import "reflect"

func DictGetSlice[T any](dict *Dict, key string) []T {
	var t T
	if reflect.TypeOf(t) == reflect.TypeOf(Dict{}) {
		return reflect.ValueOf(dict.GetDictSlice(key)).Interface().([]T)
	}
	v := dict.GetAny(key)
	if v == nil {
		return nil
	}
	r := make([]T, 0)
	vf := reflect.ValueOf(v)
	if vf.Kind() != reflect.Slice {
		return r
	}
	for i := 0; i < vf.Len(); i++ {
		r = append(r, vf.Index(i).Interface().(T))
	}
	return r
}

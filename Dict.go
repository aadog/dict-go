package dict

import (
	"fmt"
	"github.com/lrita/cmap"
	"reflect"
	"sort"
)

type Dict struct {
	mp cmap.Map[string, any]
}

func (d *Dict) Range(f func(key string, value any) bool) {
	d.mp.Range(f)
}
func (d *Dict) LoadOrStore(key string, value any) (actual any, loaded bool) {
	return d.mp.LoadOrStore(key, value)
}
func (d *Dict) IsEmpty() bool {
	return d.mp.IsEmpty()
}
func (d *Dict) Clear() {
	d.mp.Range(func(key string, val any) bool {
		d.mp.Delete(key)
		return true
	})
}
func (d *Dict) Keys() []string {
	if d.IsEmpty() {
		return nil
	}
	keys := make([]string, d.Len())
	d.Range(func(key string, value any) bool {
		keys = append(keys, key)
		return true
	})
	sort.Strings(keys)
	return keys
}
func (d *Dict) Values() []any {
	if d.IsEmpty() {
		return nil
	}
	keys := make([]any, d.Len())
	d.Range(func(key string, value any) bool {
		keys = append(keys, key)
		return true
	})
	return keys
}
func (d *Dict) Key(key string) bool {
	_, ld := d.Load(key)
	return ld
}
func (d *Dict) TypeOf(key string) reflect.Type {
	v, _ := d.Load(key)
	return reflect.TypeOf(v)
}
func (d *Dict) ValueOf(key string) reflect.Value {
	v, _ := d.Load(key)
	return reflect.ValueOf(v)
}
func (d *Dict) Del(key string) {
	d.mp.Delete(key)
}
func (d *Dict) Len() int {
	return d.mp.Count()
}
func (d *Dict) Set(key string, value any) {
	d.mp.Store(key, value)
}
func (d *Dict) Load(key string) (any, bool) {
	return d.mp.Load(key)
}
func (d *Dict) Get(key string) any {
	v, _ := d.Load(key)
	return v
}
func (d *Dict) GetAny(key string) any {
	v, _ := d.Load(key)
	return v
}
func (d *Dict) GetString(key string) string {
	v := d.GetAny(key)
	s, ok := v.(string)
	if ok {
		return s
	}
	return fmt.Sprintf("%v", v)
}
func (d *Dict) GetFloat32(key string) float32 {
	return float32(d.GetFloat64(key))
}
func (d *Dict) GetFloat64(key string) float64 {
	v := d.GetAny(key)
	if v == nil {
		return 0
	}
	valOf := reflect.ValueOf(v)
	if valOf.CanFloat() {
		return valOf.Float()
	}
	return float64(d.GetULong(key))
}
func (d *Dict) GetNumber(key string) int {
	return int(d.GetLong(key))
}
func (d *Dict) GetUNumber(key string) uint {
	return uint(int(d.GetULong(key)))
}
func (d *Dict) GetLong(key string) int64 {

	return int64(d.GetULong(key))
}
func (d *Dict) GetULong(key string) uint64 {
	v := d.GetAny(key)
	if v == nil {
		return 0
	}
	valOf := reflect.ValueOf(v)
	if valOf.CanInt() {
		return uint64(valOf.Int())
	}
	if valOf.CanUint() {
		return uint64(valOf.Uint())
	}
	if valOf.CanFloat() {
		return uint64(valOf.Float())
	}
	if valOf.Kind() == reflect.Bool {
		if valOf.Bool() == true {
			return 1
		}
	}
	return 0
}
func (d *Dict) GetBool(key string) bool {
	n := d.GetNumber(key)
	return n != 0
}

func NewDict() *Dict {
	return &Dict{}
}
func NewDictWithMap(mp interface{}) *Dict {
	d := NewDict()
	valOf := reflect.Indirect(reflect.ValueOf(mp))
	if valOf.Kind() == reflect.Map {
		keys := valOf.MapKeys()
		for _, key := range keys {
			if key.Kind() == reflect.String {
				k := key.String()
				v := valOf.MapIndex(key)
				d.Set(k, v)
			}
		}
	}
	return d
}
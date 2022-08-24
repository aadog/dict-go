package dict

import (
	"encoding/json"
	"fmt"
	"github.com/lrita/cmap"
	"github.com/mitchellh/mapstructure"
	"reflect"
	"sort"
)

type Dict struct {
	mp cmap.Map[string, any]
}

func (d *Dict) MarshalJSON() ([]byte, error) {
	mp := d.Map()
	data, err := json.Marshal(mp)
	if err != nil {
		return nil, err
	}
	return data, err
}

func (d *Dict) UnmarshalJSON(bytes []byte) error {
	mp := map[string]any{}
	err := json.Unmarshal(bytes, &mp)
	if err != nil {
		return err
	}
	for k, v := range mp {
		d.mp.Store(k, v)
	}
	return nil
}

func (d *Dict) String() string {
	return fmt.Sprintf("%v", d.Map())
}

func (d *Dict) Map() map[string]any {
	mp := map[string]any{}
	keys := d.Keys()
	for _, k := range keys {
		mp[k] = d.Get(k)
	}
	return mp
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
	keys := make([]string, 0)
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
	if v == nil {
		return ""
	}
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

func NewDictWithObj(obj any) *Dict {
	d := NewDict()
	mp := map[string]any{}
	err := mapstructure.Decode(obj, &mp)
	if err != nil {
		return d
	}
	for k, v := range mp {
		d.Set(k, v)
	}
	return d
}

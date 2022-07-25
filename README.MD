
### thread safe dict

## Start
```
go get github.com/aadog/dict-go
```

## test
```
	d := NewDictWithMap(map[string]any{
		"a": map[string]any{"a": "c"},
	})
	d.Set("b", "xxx")
	d.Set("a", 22.343)
	fmt.Println(d.GetAny("a"))
```
# T

just collection of wrappers for easy use of golang map[string]interface{}

```
var m t.Map = make(map[string]interface{})
m["value"] = 1
m.Int("value") // --> 1
m.Str("value") // --> "1"
m["value"] = "123"
m.Int("value") // --> 123
m.Str("value") // --> "123"
```
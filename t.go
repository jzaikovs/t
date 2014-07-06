package t

import (
	"encoding/json"
)

type T struct {
	Value interface{}
}

// Bool returns value of Value casted to bool.
// If Value is not bool type then returns false
func (this T) Bool() bool {
	return to_bool(this.Value)
}

func (this T) Int64() int64 {
	return to_int64(this.Value)
}

func (this T) Int() int {
	return to_int(this.Value)
}

func (this T) Uint() uint {
	return to_uint(this.Value)
}

func (this T) Uint64() uint64 {
	return to_uint64(this.Value)
}

func (this T) MarshalJSON() (b []byte, err error) {
	b, err = json.Marshal(this.Value)
	return
}

func (this T) String() string {
	return to_str(this.Value)
}

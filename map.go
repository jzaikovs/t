package t

type Map map[string]interface{}

func (this Map) Str(key string) (s string) {
	if val, ok := this[key]; ok {
		s = to_str(val)
	}
	return
}

func (this Map) Int64(key string) (n int64) {
	if val, ok := this[key]; ok {
		n = to_int64(val)
	}
	return
}

func (this Map) Int(key string) (n int) {
	if val, ok := this[key]; ok {
		n = to_int(val)
	}
	return
}

func (this Map) Bool(key string) bool {
	if val, ok := this[key]; ok {
		return to_bool(val)
	}
	return false
}

func (this Map) T(key string) (t T) {
	if val, ok := this[key]; ok {
		t = T{val}
	}
	return
}

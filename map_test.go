package t

import (
	"testing"
)

func TestMap(t *testing.T) {
	var m Map = map[string]interface{}{
		"1":     1,
		"2":     "2",
		"3":     true,
		"int8":  int8(8),
		"int16": int16(16),
		"int32": int32(32),
		"int64": int64(64),
	}

	if m.Str("2") != "2" {
		t.Fail()
	}

	if m.Int("1") != 1 {
		t.Fail()
	}

	if m.Int64("1") != 1 {
		t.Fail()
	}

	if m.Int("2") != 2 {
		t.Fail()
	}

	if m.Int64("2") != 2 {
		t.Fail()
	}

	if m.Bool("1") != false {
		t.Fail()
	}

	if m.Bool("3") == false {
		t.Fail()
	}

	if m.Int("int8") != 8 {
		t.Fail()
	}

	if m.Int("int16") != 16 {
		t.Fail()
	}

	if m.Int("int32") != 32 {
		t.Fail()
	}

	if m.Int64("int64") != 64 {
		t.Fail()
	}

	if m.Str("x") != "" {
		t.Fail()
	}

	if m.Bool("x") {
		t.Fail()
	}
}

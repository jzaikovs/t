package t

import (
	"encoding/json"
	"testing"
)

func TestTypes(t *testing.T) {
	temp := T{1}

	if temp.String() != "1" {
		t.Fail()
	}

	if temp.Int() != 1 {
		t.Fail()
	}

	if temp.Int64() != 1 {
		t.Fail()
	}

	if temp.Uint() != 1 {
		t.Fail()
	}

	if temp.Uint64() != 1 {
		t.Fail()
	}

	b, _ := json.Marshal(temp)
	b2, _ := temp.MarshalJSON()

	if string(b) != string(b2) {
		t.Log(string(b), string(b2))
		t.Fail()
	}

	temp = T{true}

	if !temp.Bool() {
		t.Fail()
	}

	if len(temp.Map()) != 0 {
		t.Fail()
	}

	temp = T{map[string]interface{}{"x": 123}}

	if temp.Map().Int("x") != 123 {
		t.Error("T.Map not working")
	}

	temp.Value = uint64(123)

	if temp.Int64() != 123 {
		t.Error("T.Int64() not working")
	}
}

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

	b, _ := json.Marshal(temp)
	b2, _ := temp.MarshalJSON()

	if string(b) != string(b2) {
		t.Log(string(b), string(b2))
		t.Fail()
	}
}

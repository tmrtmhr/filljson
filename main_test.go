package main

import (
  "encoding/json"
	"testing"
	"reflect"
)

func TestParse_string(t *testing.T) {
  got := Parse("string","string\nvalue\n")
  want := "string\nvalue\n"
	if (got != want) {
		t.Errorf("\ngot %v(%v)\nwant %v(%v)", got, reflect.TypeOf(got), want, reflect.TypeOf(want))
	}
}

func TestParse_stringList(t *testing.T) {
  got := Parse("[string]","string\nvalue\n")
  want := []string{"string","value"}
	if (!reflect.DeepEqual(got, want)) {
		t.Errorf("\ngot %v(%v)\nwant %v(%v)", got, reflect.TypeOf(got), want, reflect.TypeOf(want))
	}
}

func TestParse_int(t *testing.T) {
  got := Parse("int","123")
  var want int64 = 123
	if (got != want) {
		t.Errorf("\ngot %v(%v)\nwant %v(%v)", got, reflect.TypeOf(got), want, reflect.TypeOf(want))
	}
}

func TestParse_float(t *testing.T) {
  got := Parse("float","123")
  var want float64 = 123
	if (got != want) {
		t.Errorf("\ngot %v(%v)\nwant %v(%v)", got, reflect.TypeOf(got), want, reflect.TypeOf(want))
	}
}

func TestFill(t *testing.T) {
  var got map[string]interface{}
  json.Unmarshal([]byte("{ \"prop1\": { \"prop2\": null } }"), &got)
  intNum := Parse("int", "123")
  propNames := []string{"prop1", "prop2"}
  want := map[string]interface{}{
    "prop1": map[string]interface{}{ "prop2": int64(123) } }
  got = Fill(got, propNames, intNum)
	if (!reflect.DeepEqual(got, want)) {
		t.Errorf("\ngot %v(%v)\nwant %v(%v)", got, reflect.TypeOf(got), want, reflect.TypeOf(want))
    t.Errorf("\ngot.prop1 %v(%v)\nwant.prop1 %v(%v)", got["prop1"], reflect.TypeOf(got["prop1"]), want["prop1"], reflect.TypeOf(want["prop1"]))
    t.Errorf("\ngot.prop1.prop2 %v(%v)\nwant.prop1.prop2 %v(%v)", got["prop1"].(map[string]interface{})["prop2"], reflect.TypeOf(got["prop1"].(map[string]interface{})["prop2"]), want["prop1"].(map[string]interface{})["prop2"], reflect.TypeOf(want["prop1"].(map[string]interface{})["prop2"]))
	}
}

package fields_test

import (
	"gitlab.com/distributed_lab/logan/v2/fields"
	"reflect"
	"testing"
)

type plain struct {
	Field1 string
	Field2 int
}

type fielded struct {
	Field1 string
	Field2 int
}

func (f fielded) GetLoganFields() map[string]interface{} {
	return map[string]interface{}{
		"field1": f.Field1,
		"field2": f.Field2,
	}
}

type nestedPlain struct {
	Field1 string
	Field2 int
	Field3 plain
}

func (f nestedPlain) GetLoganFields() map[string]interface{} {
	return map[string]interface{}{
		"field1": f.Field1,
		"field2": f.Field2,
		"field3": f.Field3,
	}
}

type nestedFielded struct {
	Field1 string
	Field2 int
	Field3 fielded
}

func (f nestedFielded) GetLoganFields() map[string]interface{} {
	return map[string]interface{}{
		"field1": f.Field1,
		"field2": f.Field2,
		"field3": f.Field3,
	}
}

func TestObtain(t *testing.T) {
	p := plain{"field1value1", 1}
	f := fielded{"field1value2", 2}
	nWithP := nestedPlain{"field1value3", 3, p}
	nWithF := nestedFielded{"field1value4", 4, f}
	var nilFielded *fielded

	cases := []struct {
		name  string
		key   string
		value interface{}
		out   map[string]interface{}
	}{
		{"plain value", "foo", 7, map[string]interface{}{"foo": 7}},
		{"plain entity", "foo", p, map[string]interface{}{"foo": p}},
		{"fielded entity", "foo", f, map[string]interface{}{"foo_field1": f.Field1, "foo_field2": f.Field2}},
		{"nested entity with plain inside", "foo", nWithP, map[string]interface{}{"foo_field1": nWithP.Field1, "foo_field2": nWithP.Field2, "foo_field3": nWithP.Field3}},
		{"nested entity with fielded inside", "foo", nWithF, map[string]interface{}{"foo_field1": nWithF.Field1, "foo_field2": nWithF.Field2,
			"foo_field3_field1": nWithF.Field3.Field1, "foo_field3_field2": nWithF.Field3.Field2}},
		{"nil value", "foo", nil, map[string]interface{}{"foo": nil}},
		{"nil value of type FieldedEntity", "foo", nilFielded, map[string]interface{}{"foo": nil}},
	}

	for _, tc := range cases {
		got := fields.Obtain(tc.key, tc.value)
		if !reflect.DeepEqual(got, tc.out) {
			t.Errorf("%s: got %#v, want %#v", tc.name, got, tc.out)
		}
	}
}

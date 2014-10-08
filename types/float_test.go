package types

import (
	"reflect"
	"testing"
)

func TestParseFloat(t *testing.T) {
	for _, tt := range []struct {
		val string
		exp interface{}
		ok  bool
	}{
		{"0.0", float32(0.0), true},
		{"10.2", float32(10.2), true},
		{"-10.2", float32(-10.2), true},
		{"0.0", float64(0.0), true},
		{"10.2", float64(10.2), true},
		{"-10.2", float64(-10.2), true},
		// Sscanf does not throw error for this fix it!
		{"-23f.f0xer", float32(-23.0), true},
	} {
		typ := reflect.TypeOf(tt.exp)
		res := reflect.New(typ).Interface()
		err := ParseFloat(res, tt.val)
		switch {
		case tt.ok && err != nil:
			t.Errorf("ParseFloat(%v, %#v): fail; got error %v, want ok",
				typ, tt.val, err)
		case !tt.ok && err == nil:
			t.Errorf("ParseFloat(%v, %#v): fail; got %v, want error",
				typ, tt.val, elem(res))
		case tt.ok && !reflect.DeepEqual(elem(res), tt.exp):
			t.Errorf("ParseFloat(%v, %#v): fail; got %v, want %v",
				typ, tt.val, elem(res), tt.exp)
		default:
			t.Logf("ParseFloat(%v, %#v): pass; got %v, error %v",
				typ, tt.val, elem(res), err)
		}
	}
}

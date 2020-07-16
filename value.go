package pack

import (
	"encoding/json"
	"math/rand"
	"reflect"
	"testing/quick"

	"github.com/renproject/surge"
)

// A Value is a common interface for all values that are able to be marshaled to
// binary and JSON, and are able to express their type information.
type Value interface {
	surge.Marshaler
	json.Marshaler

	Type() Type
}

// Generate a random value. This is helpful when implementing generators for
// other types. See https://golang.org/pkg/testing/quick/#Generator for more
// information.
func Generate(r *rand.Rand, size int, allowStruct bool) reflect.Value {
	t := reflect.Type(nil)
	kind, _ := quick.Value(reflect.TypeOf(Kind(0)), r)
	switch kind.Interface().(Kind) {
	case KindBool:
		t = reflect.TypeOf(Bool(false))
	case KindU8:
		t = reflect.TypeOf(U8{})
	case KindU16:
		t = reflect.TypeOf(U16{})
	case KindU32:
		t = reflect.TypeOf(U32{})
	case KindU64:
		t = reflect.TypeOf(U64{})
	case KindU128:
		t = reflect.TypeOf(U128{})
	case KindU256:
		t = reflect.TypeOf(U256{})
	case KindString:
		t = reflect.TypeOf(String(""))
	case KindBytes:
		t = reflect.TypeOf(Bytes{})
	case KindBytes32:
		t = reflect.TypeOf(Bytes32{})
	case KindBytes65:
		t = reflect.TypeOf(Bytes65{})
	case KindStruct:
		if !allowStruct {
			return Generate(r, size, allowStruct)
		}
		t = reflect.TypeOf(Struct{})
	default:
		panic("non-exhaustive pattern")
	}
	v, _ := quick.Value(t, r)
	return v
}

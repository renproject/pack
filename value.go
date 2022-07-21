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
func Generate(r *rand.Rand, size int, allowStruct, allowList bool) reflect.Value {
	kind, _ := quick.Value(reflect.TypeOf(Kind(0)), r)
	return GenerateFromKind(r, size, kind.Interface().(Kind), allowStruct, allowList)
}

// GenerateFromKind generates a random value given a Kind.
func GenerateFromKind(r *rand.Rand, size int, kind Kind, allowStruct, allowList bool) reflect.Value {
	t := reflect.Type(nil)
	switch kind {
	case KindBool:
		t = reflect.TypeOf(Bool(false))
	case KindU8:
		t = reflect.TypeOf(U8(0))
	case KindU16:
		t = reflect.TypeOf(U16(0))
	case KindU32:
		t = reflect.TypeOf(U32(0))
	case KindU64:
		t = reflect.TypeOf(U64(0))
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
	case KindBytes64:
		t = reflect.TypeOf(Bytes64{})
	case KindStruct:
		if !allowStruct {
			return Generate(r, size, allowStruct, allowList)
		}
		t = reflect.TypeOf(Struct{})
	case KindList:
		if !allowList {
			return Generate(r, size, allowStruct, allowList)
		}
		t = reflect.TypeOf(List{})
	default:
		panic("non-exhaustive pattern")
	}
	v, _ := quick.Value(t, r)
	return v
}

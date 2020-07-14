package pack

import (
	"encoding/json"
	"math/rand"
	"reflect"
	"testing/quick"

	"github.com/renproject/surge"
)

// Bool represents a value that can be true or false.
type Bool bool

// NewBool wraps an existing raw boolean.
func NewBool(x bool) Bool {
	return Bool(x)
}

// Type returns boolean type.
func (x Bool) Type() Type {
	return typeBool{}
}

// Equal returns true when x is equal to y. Otherwise, it returns false.
func (x Bool) Equal(y Bool) bool {
	return bool(x) == bool(y)
}

// SizeHint returns the number of bytes required to represent the boolean in
// binary.
func (x Bool) SizeHint() int {
	return surge.SizeHintBool
}

// Marshal the boolean to binary.
func (x Bool) Marshal(buf []byte, rem int) ([]byte, int, error) {
	return surge.MarshalBool(bool(x), buf, rem)
}

// Unmarshal the boolean from binary.
func (x *Bool) Unmarshal(buf []byte, rem int) ([]byte, int, error) {
	return surge.UnmarshalBool((*bool)(x), buf, rem)
}

// MarshalJSON marshals the boolean to JSON. This is done using the default JSON
// marshaler for raw booleans.
func (x Bool) MarshalJSON() ([]byte, error) {
	return json.Marshal(bool(x))
}

// UnmarshalJSON unmarshals the boolean from JSON. This is done using the
// default JSON unmarshaler for raw booleans.
func (x *Bool) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, (*bool)(x))
}

// String returns "true" when the boolean is true. Otherwise, it returns
// "false".
func (x Bool) String() string {
	if x {
		return "true"
	}
	return "false"
}

// Generate a random boolean. This method is implemented for use in quick tests.
// See https://golang.org/pkg/testing/quick/#Generator for more information.
func (Bool) Generate(r *rand.Rand, size int) reflect.Value {
	v, _ := quick.Value(reflect.TypeOf(false), r)
	return reflect.ValueOf(NewBool(v.Bool()))
}
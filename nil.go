package pack

import (
	"math/rand"
	"reflect"
)

// Nil represents a nil value.
type Nil struct{}

// NewNil wraps an existing raw struct.
func NewNil(x struct{}) Nil {
	return Nil(x)
}

// Type returns nil type.
func (x Nil) Type() Type {
	return typeNil{}
}

// Equal returns true when x is equal to y. Otherwise, it returns false.
func (x Nil) Equal(y Nil) bool {
	return x == y
}

// SizeHint returns the number of bytes required to represent the nil value in
// binary.
func (x Nil) SizeHint() int {
	return 0
}

// Marshal the boolean to binary.
func (x Nil) Marshal(buf []byte, rem int) ([]byte, int, error) {
	return buf, rem, nil
}

// Unmarshal the boolean from binary.
func (x *Nil) Unmarshal(buf []byte, rem int) ([]byte, int, error) {
	x = (*Nil)(&struct{}{})
	return buf, rem, nil
}

// MarshalJSON marshals the nil value to JSON.
func (x Nil) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals the nil value from JSON.
func (x *Nil) UnmarshalJSON(data []byte) error {
	x = (*Nil)(&struct{}{})
	return nil
}

// String returns "" for nil values.
func (x Nil) String() string {
	return ""
}

// Generate a random nil value. This method is implemented for use in quick
// tests. See https://golang.org/pkg/testing/quick/#Generator for more
// information.
func (Nil) Generate(r *rand.Rand, size int) reflect.Value {
	return reflect.ValueOf(NewNil(struct{}{}))
}

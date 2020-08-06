package pack

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/rand"
	"reflect"
	"testing/quick"

	"github.com/renproject/surge"
)

// MaxBytes of an object. Defaults to 32 MB.
var MaxBytes = 32 * 1024 * 1024

// String represents a slice of bytes that should be interpreted as a UTF-8
// encoded string.
type String string

// NewString wraps an existing raw string.
func NewString(x string) String {
	return String(x)
}

// Type returns the string stype.
func (String) Type() Type {
	return typeString{}
}

// Equal returns true when x is equal to y. Otherwise, it returns false.
func (x String) Equal(y String) bool {
	return string(x) == string(y)
}

// SizeHint returns the number of bytes required to represent the string in
// binary.
func (x String) SizeHint() int {
	return surge.SizeHintString(string(x))
}

// Marshal the string to binary.
func (x String) Marshal(buf []byte, rem int) ([]byte, int, error) {
	return surge.MarshalString(string(x), buf, rem)
}

// Unmarshal the string from binary.
func (x *String) Unmarshal(buf []byte, rem int) ([]byte, int, error) {
	return surge.UnmarshalString((*string)(x), buf, rem)
}

// MarshalJSON marshals the string to JSON. This is done using the default JSON
// marshaler for raw strings.
func (x String) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(x))
}

// UnmarshalJSON unmarshals the string from JSON. This is done using the default
// JSON unmarshaler for raw strings.
func (x *String) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, (*string)(x))
}

// String returns the raw string. This method exists to implement the Stringer
// interface.
func (x String) String() string {
	return string(x)
}

// Generate a random string. This method is implemented for use in quick tests.
// See https://golang.org/pkg/testing/quick/#Generator for more information.
func (String) Generate(r *rand.Rand, size int) reflect.Value {
	runes := make([]rune, size)
	for i := 0; i < size; i++ {
		runes[i] = rune(r.Intn(0x10ffff))
	}
	return reflect.ValueOf(String(string(runes)))
}

// Bytes represents a slice of bytes with a dynamic length.
type Bytes []byte

// NewBytes wraps an exsiting raw slice of bytes.
func NewBytes(x []byte) Bytes {
	if x == nil {
		x = []byte{}
	}
	return Bytes(x)
}

// Type returns the bytes type.
func (Bytes) Type() Type {
	return typeBytes{}
}

// Equal returns true when x is equal to y. Otherwise, it returns false.
func (x Bytes) Equal(y Bytes) bool {
	return bytes.Equal([]byte(x), []byte(y))
}

// SizeHint returns the number of bytes required to represent the bytes in
// binary. This includes the length prefix that encodes the dynamic length of
// the bytes.
func (x Bytes) SizeHint() int {
	return surge.SizeHintBytes([]byte(x))
}

// Marshal the bytes to binary.
func (x Bytes) Marshal(buf []byte, rem int) ([]byte, int, error) {
	return surge.MarshalBytes([]byte(x), buf, rem)
}

// Unmarshal the bytes from binary.
func (x *Bytes) Unmarshal(buf []byte, rem int) ([]byte, int, error) {
	return surge.UnmarshalBytes((*[]byte)(x), buf, rem)
}

// MarshalJSON marshals the bytes to JSON. This is done by encoding the bytes
// into a base64 raw URL encoded string.
func (x Bytes) MarshalJSON() ([]byte, error) {
	return json.Marshal(base64.RawURLEncoding.EncodeToString([]byte(x)))
}

// UnmarshalJSON unmarshals the bytes from JSON. This is done by decoding the
// bytes from a base64 raw URL encoded string.
func (x *Bytes) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	data, err := base64.RawURLEncoding.DecodeString(str)
	if err != nil {
		return err
	}
	*x = data
	return nil
}

// MarshalText marshals the bytes to text.
func (x Bytes) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

// String returns a base64 raw URL encoding of the bytes.
func (x Bytes) String() string {
	return base64.RawURLEncoding.EncodeToString([]byte(x))
}

// Generate a random byte slice. This method is implemented for use in quick
// tests. See https://golang.org/pkg/testing/quick/#Generator for more
// information.
func (Bytes) Generate(r *rand.Rand, size int) reflect.Value {
	data := make([]byte, size)
	if _, err := r.Read(data); err != nil {
		panic(err)
	}
	return reflect.ValueOf(Bytes(data))
}

// Bytes32 represents a static-sized 32-byte array.
type Bytes32 [32]byte

// NewBytes32 wraps an existing raw 32-byte array.
func NewBytes32(x [32]byte) Bytes32 {
	return Bytes32(x)
}

// Type returns the 32-byte array type.
func (Bytes32) Type() Type {
	return typeBytes32{}
}

// Bytes returns a copy of the byte array as a dynamic byte slice.
func (x Bytes32) Bytes() []byte {
	copied := x
	return copied[:]
}

// Equal returns true when x is equal to y. Otherwise, it returns false.
func (x Bytes32) Equal(y *Bytes32) bool {
	return bytes.Equal(x[:], y[:])
}

// SizeHint returns the number of bytes required to represent the 32-byte array
// in binary.
func (x Bytes32) SizeHint() int {
	return 32
}

// Marshal the 32-byte array to binary.
func (x Bytes32) Marshal(buf []byte, rem int) ([]byte, int, error) {
	if len(buf) < 32 || rem < 32 {
		return buf, rem, surge.ErrUnexpectedEndOfBuffer
	}
	copy(buf, x[:])
	return buf[32:], rem - 32, nil
}

// Unmarshal the 32-byte array from binary.
func (x *Bytes32) Unmarshal(buf []byte, rem int) ([]byte, int, error) {
	if len(buf) < 32 || rem < 32 {
		return buf, rem, surge.ErrUnexpectedEndOfBuffer
	}
	copy(x[:], buf[:32])
	return buf[32:], rem - 32, nil
}

// MarshalJSON marshals the 32-byte array to JSON. This is done by encoding the
// bytes into a base64 raw URL encoded string.
func (x Bytes32) MarshalJSON() ([]byte, error) {
	return json.Marshal(base64.RawURLEncoding.EncodeToString(x[:]))
}

// UnmarshalJSON unmarshals the 32-byte array from JSON. This is done by
// decoding the bytes from a base64 raw URL encoded string.
func (x *Bytes32) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	data, err := base64.RawURLEncoding.DecodeString(str)
	if err != nil {
		return err
	}
	if len(data) != 32 {
		return fmt.Errorf("expected len=32, got len=%v", len(data))
	}
	copy(x[:], data)
	return nil
}

// MarshalText marshals the 32-byte array to text.
func (x Bytes32) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

// String returns a base64 raw URL encoding of the bytes.
func (x Bytes32) String() string {
	return base64.RawURLEncoding.EncodeToString(x[:])
}

// Generate a random 32-byte array. This method is implemented for use in quick
// tests. See https://golang.org/pkg/testing/quick/#Generator for more
// information.
func (Bytes32) Generate(r *rand.Rand, size int) reflect.Value {
	v, _ := quick.Value(reflect.TypeOf([32]byte{}), r)
	return reflect.ValueOf(NewBytes32(v.Interface().([32]byte)))
}

// Bytes65 represents a static-sized 65-byte array.
type Bytes65 [65]byte

// NewBytes65 wraps an existing raw 65-byte array.
func NewBytes65(x [65]byte) Bytes65 {
	return Bytes65(x)
}

// Type returns the 65-byte array type.
func (Bytes65) Type() Type {
	return typeBytes65{}
}

// Bytes returns a copy of the byte array as a dynamic byte slice.
func (x Bytes65) Bytes() []byte {
	copied := x
	return copied[:]
}

// Equal returns true when x is equal to y. Otherwise, it returns false.
func (x Bytes65) Equal(y *Bytes65) bool {
	return bytes.Equal(x[:], y[:])
}

// SizeHint returns the number of bytes required to represent the 65-byte array
// in binary.
func (x Bytes65) SizeHint() int {
	return 65
}

// Marshal the 65-byte array to binary.
func (x Bytes65) Marshal(buf []byte, rem int) ([]byte, int, error) {
	if len(buf) < 65 || rem < 65 {
		return buf, rem, surge.ErrUnexpectedEndOfBuffer
	}
	copy(buf, x[:])
	return buf[65:], rem - 65, nil
}

// Unmarshal the 65-byte array from binary.
func (x *Bytes65) Unmarshal(buf []byte, rem int) ([]byte, int, error) {
	if len(buf) < 65 || rem < 65 {
		return buf, rem, surge.ErrUnexpectedEndOfBuffer
	}
	copy(x[:], buf[:65])
	return buf[65:], rem - 65, nil
}

// MarshalJSON marshals the 65-byte array to JSON. This is done by encoding the
// bytes into a base64 raw URL encoded string.
func (x Bytes65) MarshalJSON() ([]byte, error) {
	return json.Marshal(base64.RawURLEncoding.EncodeToString(x[:]))
}

// UnmarshalJSON unmarshals the 65-byte array from JSON. This is done by
// decoding the bytes from a base64 raw URL encoded string.
func (x *Bytes65) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	data, err := base64.RawURLEncoding.DecodeString(str)
	if err != nil {
		return err
	}
	if len(data) != 65 {
		return fmt.Errorf("expected len=65, got len=%v", len(data))
	}
	copy(x[:], data)
	return nil
}

// MarshalText marshals the 65-byte array to text.
func (x Bytes65) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

// String returns a base64 raw URL encoding of the bytes.
func (x Bytes65) String() string {
	return base64.RawURLEncoding.EncodeToString(x[:])
}

// Generate a random 65-byte array. This method is implemented for use in quick
// tests. See https://golang.org/pkg/testing/quick/#Generator for more
// information.
func (Bytes65) Generate(r *rand.Rand, size int) reflect.Value {
	v, _ := quick.Value(reflect.TypeOf([65]byte{}), r)
	return reflect.ValueOf(NewBytes65(v.Interface().([65]byte)))
}

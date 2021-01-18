package pack

import (
	"math/rand"
	"reflect"
	"strings"

	"github.com/renproject/surge"
)

// A Kind is an optionally abstract type identifier. It can be thought of as the
// "type of a type", or the "constructor for a type". For example, a struct is
// an abstract type identifier, because it does not bind the respective value to
// have any specific fields (or types for those fields).  Similarly, a list is
// an abstract type identifier, because it does not bind the respective value to
// a specific type of element.
type Kind uint8

const (
	// KindNil is the kind for elements within an empty List.
	KindNil = Kind(0)

	// KindBool is the kind of all Bool values.
	KindBool = Kind(1)
	// KindU8 is the kind of all U8 values.
	KindU8 = Kind(2)
	// KindU16 is the kind of all U16 values.
	KindU16 = Kind(3)
	// KindU32 is the kind of all U32 values.
	KindU32 = Kind(4)
	// KindU64 is the kind of all U64 values.
	KindU64 = Kind(5)
	// KindU128 is the kind of all U128 values.
	KindU128 = Kind(6)
	// KindU256 is the kind of all U256 values.
	KindU256 = Kind(7)

	// KindString is the kind of all utf8 strings.
	KindString = Kind(10)
	// KindBytes is the kind of all dynamic byte arrays.
	KindBytes = Kind(11)
	// KindBytes32 is the kind of all 32-byte arrays.
	KindBytes32 = Kind(12)
	// KindBytes65 is the kind of all 65-byte arrays.
	KindBytes65 = Kind(13)

	// KindStruct is the kind of all struct values. It is abstract, because it does
	// not specify the fields in the struct.
	KindStruct = Kind(20)
	// KindList is the kind of all list values. It is abstract, because it does
	// not specify the type of the elements in the list.
	KindList = Kind(21)
)

func (kind Kind) String() string {
	switch kind {
	// Scalar
	case KindBool:
		return "bool"
	case KindU8:
		return "u8"
	case KindU16:
		return "u16"
	case KindU32:
		return "u32"
	case KindU64:
		return "u64"
	case KindU128:
		return "u128"
	case KindU256:
		return "u256"

	// Bytes
	case KindString:
		return "string"
	case KindBytes:
		return "bytes"
	case KindBytes32:
		return "bytes32"
	case KindBytes65:
		return "bytes65"

	// Abstract
	case KindStruct:
		return "struct"
	case KindList:
		return "list"
	default:
		return "nil"
	}
}

// SizeHint returns the number of bytes required to represent the kind in
// binary.
func (kind Kind) SizeHint() int {
	return surge.SizeHintU8
}

// Marshal the kind into binary.
func (kind Kind) Marshal(buf []byte, rem int) ([]byte, int, error) {
	return surge.MarshalU8(uint8(kind), buf, rem)
}

// Unmarshal the kind from binary.
func (kind *Kind) Unmarshal(buf []byte, rem int) ([]byte, int, error) {
	return surge.UnmarshalU8((*uint8)(kind), buf, rem)
}

// MarshalText from the kind. Unrecognised kinds will be marshaled into the
// "nil" string.
func (kind Kind) MarshalText() ([]byte, error) {
	return []byte(kind.String()), nil
}

// UnmarshalText into the kind. Unrecognised text will be unmarshaled into the
// KindNil, and will be considered invalid.
func (kind *Kind) UnmarshalText(text []byte) error {
	switch strings.ToLower(string(text)) {
	case KindBool.String():
		*kind = KindBool
		return nil
	case KindU8.String():
		*kind = KindU8
		return nil
	case KindU16.String():
		*kind = KindU16
		return nil
	case KindU32.String():
		*kind = KindU32
		return nil
	case KindU64.String():
		*kind = KindU64
		return nil
	case KindU128.String():
		*kind = KindU128
		return nil
	case KindU256.String():
		*kind = KindU256
		return nil
	case KindString.String():
		*kind = KindString
		return nil
	case KindBytes.String():
		*kind = KindBytes
		return nil
	case KindBytes32.String():
		*kind = KindBytes32
		return nil
	case KindBytes65.String():
		*kind = KindBytes65
		return nil
	case KindStruct.String():
		*kind = KindStruct
		return nil
	case KindList.String():
		*kind = KindList
		return nil
	default:
		*kind = KindNil
		return nil
	}
}

// Generate a random kind. This method is implemented for use in quick tests.
// See https://golang.org/pkg/testing/quick/#Generator for more information.
func (kind Kind) Generate(r *rand.Rand, size int) reflect.Value {
	switch r.Int() % 12 {
	case 0:
		return reflect.ValueOf(KindBool)
	case 1:
		return reflect.ValueOf(KindU8)
	case 2:
		return reflect.ValueOf(KindU16)
	case 3:
		return reflect.ValueOf(KindU32)
	case 4:
		return reflect.ValueOf(KindU64)
	case 5:
		return reflect.ValueOf(KindU128)
	case 6:
		return reflect.ValueOf(KindU256)
	case 7:
		return reflect.ValueOf(KindString)
	case 8:
		return reflect.ValueOf(KindBytes)
	case 9:
		return reflect.ValueOf(KindBytes32)
	case 10:
		return reflect.ValueOf(KindBytes65)
	case 11:
		return reflect.ValueOf(KindStruct)
	case 12:
		return reflect.ValueOf(KindList)
	default:
		// It is intentional that this case never happens.
		return reflect.ValueOf(KindNil)
	}
}

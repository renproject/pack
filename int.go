package pack

import (
	"encoding/json"
	"fmt"
	"math/big"
	"math/rand"
	"reflect"
	"strconv"
	"testing/quick"

	"github.com/renproject/surge"
)

// U8 represents an 8-bit unsigned integer.
type U8 uint8

// NewU8 returns a uint8 wrapped as a U8.
func NewU8(x uint8) U8 {
	return U8(x)
}

// Type returns the type of this value.
func (U8) Type() Type {
	return typeU8{}
}

// Uint8 returns the inner uint8.
func (u8 U8) Uint8() uint8 {
	return uint8(u8)
}

// Add one U8 to another and return the result.
func (u8 U8) Add(other U8) U8 {
	ret := U8(u8 + other)
	if ret < u8 {
		panic("overflow")
	}
	return ret
}

// Sub one U8 from another and return the result.
func (u8 U8) Sub(other U8) U8 {
	ret := U8(u8 - other)
	if ret > u8 {
		panic("underflow")
	}
	return ret
}

// AddAssign will add one U8 to another and assign the result to the left-hand
// side.
func (u8 *U8) AddAssign(other U8) {
	*u8 = *u8 + other
	if *u8 < other {
		panic("overflow")
	}
}

// SubAssign will sub one U8 from another and assign the result to the left-hand
// side.
func (u8 *U8) SubAssign(other U8) {
	res := *u8 - other
	if res > *u8 {
		panic("underflow")
	}
	*u8 = res
}

// Equal compares one U8 to another. If they are equal, then it returns true.
// Otherwise, it returns false.
func (u8 U8) Equal(other U8) bool {
	return u8 == other
}

// SizeHint returns the number of bytes required to represent a U8 in binary.
func (u8 U8) SizeHint() int {
	return 1
}

// Marshal the U8 to binary.
func (u8 U8) Marshal(buf []byte, rem int) ([]byte, int, error) {
	return surge.MarshalU8(uint8(u8), buf, rem)
}

// Unmarshal the U8 from binary.
func (u8 *U8) Unmarshal(buf []byte, rem int) ([]byte, int, error) {
	return surge.UnmarshalU8((*uint8)(u8), buf, rem)
}

// MarshalJSON implements the JSON marshaler interface. U8s are marshaled as
// decimal strings (for consistency with larger integer types).
func (u8 U8) MarshalJSON() ([]byte, error) {
	return json.Marshal(u8.String())
}

// UnmarshalJSON implements the JSON unmarshaler interface. U8s are unmarshaled
// as decimal strings (for consistency with larger integer types).
func (u8 *U8) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	x, err := strconv.ParseUint(str, 10, 8)
	if err != nil {
		return err
	}
	*u8 = U8(x)
	return nil
}

func (u8 U8) String() string {
	return fmt.Sprintf("%v", uint8(u8))
}

// U16 represents a 16-bit unsigned integer.
type U16 uint16

// NewU16 returns a uint16 wrapped as a U16.
func NewU16(x uint16) U16 {
	return U16(x)
}

// NewU16FromU8 returns a uint16 wrapped as a U16.
func NewU16FromU8(x U8) U16 {
	return U16(uint16(x.Uint8()))
}

// Type returns the type of this value.
func (U16) Type() Type {
	return typeU16{}
}

// Uint16 returns the inner uint16.
func (u16 U16) Uint16() uint16 {
	return uint16(u16)
}

// Add one U16 to another and return the result.
func (u16 U16) Add(other U16) U16 {
	ret := U16(u16 + other)
	if ret < u16 {
		panic("overflow")
	}
	return ret
}

// Sub one U16 from another and return the result.
func (u16 U16) Sub(other U16) U16 {
	ret := U16(u16 - other)
	if ret > u16 {
		panic("underflow")
	}
	return ret
}

// AddAssign will add one U16 to another and assign the result to the left-hand
// side.
func (u16 *U16) AddAssign(other U16) {
	*u16 = *u16 + other
	if *u16 < other {
		panic("overflow")
	}
}

// SubAssign will sub one U16 from another and assign the result to the left-hand
// side.
func (u16 *U16) SubAssign(other U16) {
	res := *u16 - other
	if res > *u16 {
		panic("underflow")
	}
	*u16 = res
}

// Equal compares one U16 to another. If they are equal, then it returns true.
// Otherwise, it returns false.
func (u16 U16) Equal(other U16) bool {
	return u16 == other
}

// SizeHint returns the number of bytes required to represent a U16 in binary.
func (u16 U16) SizeHint() int {
	return 2
}

// Marshal the U16 to binary.
func (u16 U16) Marshal(buf []byte, rem int) ([]byte, int, error) {
	return surge.MarshalU16(uint16(u16), buf, rem)
}

// Unmarshal the U16 from binary.
func (u16 *U16) Unmarshal(buf []byte, rem int) ([]byte, int, error) {
	return surge.UnmarshalU16((*uint16)(u16), buf, rem)
}

// MarshalJSON implements the JSON marshaler interface. U16s are marshaled as
// decimal strings (for consistency with larger integer types).
func (u16 U16) MarshalJSON() ([]byte, error) {
	return json.Marshal(u16.String())
}

// UnmarshalJSON implements the JSON unmarshaler interface. U16s are unmarshaled
// as decimal strings (for consistency with larger integer types).
func (u16 *U16) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	x, err := strconv.ParseUint(str, 10, 16)
	if err != nil {
		return err
	}
	*u16 = U16(x)
	return nil
}

func (u16 U16) String() string {
	return fmt.Sprintf("%v", uint16(u16))
}

// U32 represents a 32-bit unsigned integer.
type U32 uint32

// NewU32 returns a uint32 wrapped as a U32.
func NewU32(x uint32) U32 {
	return U32(x)
}

// NewU32FromU8 returns a uint32 wrapped as a U32.
func NewU32FromU8(x U8) U32 {
	return U32(uint32(x.Uint8()))
}

// NewU32FromU16 returns a uint32 wrapped as a U32.
func NewU32FromU16(x U16) U32 {
	return U32(uint32(x.Uint16()))
}

// Type returns the type of this value.
func (U32) Type() Type {
	return typeU32{}
}

// Uint32 returns the inner uint32.
func (u32 U32) Uint32() uint32 {
	return uint32(u32)
}

// Add one U32 to another and return the result.
func (u32 U32) Add(other U32) U32 {
	ret := U32(u32 + other)
	if ret < u32 {
		panic("overflow")
	}
	return ret
}

// Sub one U32 from another and return the result.
func (u32 U32) Sub(other U32) U32 {
	ret := U32(u32 - other)
	if ret > u32 {
		panic("underflow")
	}
	return ret
}

// AddAssign will add one U32 to another and assign the result to the left-hand
// side.
func (u32 *U32) AddAssign(other U32) {
	*u32 = *u32 + other
	if *u32 < other {
		panic("overflow")
	}
}

// SubAssign will sub one U32 from another and assign the result to the left-hand
// side.
func (u32 *U32) SubAssign(other U32) {
	res := *u32 - other
	if res > *u32 {
		panic("underflow")
	}
	*u32 = res
}

// Equal compares one U32 to another. If they are equal, then it returns true.
// Otherwise, it returns false.
func (u32 U32) Equal(other U32) bool {
	return u32 == other
}

// SizeHint returns the number of bytes required to represent a U32 in binary.
func (u32 U32) SizeHint() int {
	return 4
}

// Marshal the U32 to binary.
func (u32 U32) Marshal(buf []byte, rem int) ([]byte, int, error) {
	return surge.MarshalU32(uint32(u32), buf, rem)
}

// Unmarshal the U32 from binary.
func (u32 *U32) Unmarshal(buf []byte, rem int) ([]byte, int, error) {
	return surge.UnmarshalU32((*uint32)(u32), buf, rem)
}

// MarshalJSON implements the JSON marshaler interface. U32s are marshaled as
// decimal strings (for consistency with larger integer types).
func (u32 U32) MarshalJSON() ([]byte, error) {
	return json.Marshal(u32.String())
}

// UnmarshalJSON implements the JSON unmarshaler interface. U32s are unmarshaled
// as decimal strings (for consistency with larger integer types).
func (u32 *U32) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	x, err := strconv.ParseUint(str, 10, 32)
	if err != nil {
		return err
	}
	*u32 = U32(x)
	return nil
}

func (u32 U32) String() string {
	return fmt.Sprintf("%v", uint32(u32))
}

// U64 represents a 64-bit unsigned integer.
type U64 uint64

// NewU64 returns a uint64 wrapped as a U64.
func NewU64(x uint64) U64 {
	return U64(x)
}

// NewU64FromU8 returns a uint64 wrapped as a U64.
func NewU64FromU8(x U8) U64 {
	return U64(uint64(x.Uint8()))
}

// NewU64FromU16 returns a uint64 wrapped as a U64.
func NewU64FromU16(x U16) U64 {
	return U64(uint64(x.Uint16()))
}

// NewU64FromU32 returns a uint32 wrapped as a U64.
func NewU64FromU32(x U32) U64 {
	return U64(uint64(x.Uint32()))
}

// Type returns the type of this value.
func (U64) Type() Type {
	return typeU64{}
}

// Uint64 returns the inner uint64.
func (u64 U64) Uint64() uint64 {
	return uint64(u64)
}

// Add one U64 to another and return the result.
func (u64 U64) Add(other U64) U64 {
	ret := U64(u64 + other)
	if ret < u64 {
		panic("overflow")
	}
	return ret
}

// Sub one U64 from another and return the result.
func (u64 U64) Sub(other U64) U64 {
	ret := U64(u64 - other)
	if ret > u64 {
		panic("underflow")
	}
	return ret
}

// AddAssign will add one U64 to another and assign the result to the left-hand
// side.
func (u64 *U64) AddAssign(other U64) {
	*u64 = *u64 + other
	if *u64 < other {
		panic("overflow")
	}
}

// SubAssign will sub one U64 from another and assign the result to the left-hand
// side.
func (u64 *U64) SubAssign(other U64) {
	res := *u64 - other
	if res > *u64 {
		panic("underflow")
	}
	*u64 = res
}

// Equal compares one U64 to another. If they are equal, then it returns true.
// Otherwise, it returns false.
func (u64 U64) Equal(other U64) bool {
	return u64 == other
}

// SizeHint returns the number of bytes required to represent a U64 in binary.
func (u64 U64) SizeHint() int {
	return 8
}

// Marshal the U64 to binary.
func (u64 U64) Marshal(buf []byte, rem int) ([]byte, int, error) {
	return surge.MarshalU64(uint64(u64), buf, rem)
}

// Unmarshal the U64 from binary.
func (u64 *U64) Unmarshal(buf []byte, rem int) ([]byte, int, error) {
	return surge.UnmarshalU64((*uint64)(u64), buf, rem)
}

// MarshalJSON implements the JSON marshaler interface. U64s are marshaled as
// decimal strings (for consistency with larger integer types).
func (u64 U64) MarshalJSON() ([]byte, error) {
	return json.Marshal(u64.String())
}

// UnmarshalJSON implements the JSON unmarshaler interface. U64s are unmarshaled
// as decimal strings (for consistency with larger integer types).
func (u64 *U64) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	x, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return err
	}
	*u64 = U64(x)
	return nil
}

func (u64 U64) String() string {
	return fmt.Sprintf("%v", uint64(u64))
}

type U128 struct {
	inner *big.Int
}

func NewU128(x [16]byte) U128 {
	return U128{inner: new(big.Int).SetBytes(x[:])}
}

func NewU128FromU8(x U8) U128 {
	return U128{inner: new(big.Int).SetUint64(uint64(x.Uint8()))}
}

func NewU128FromU16(x U16) U128 {
	return U128{inner: new(big.Int).SetUint64(uint64(x.Uint16()))}
}

func NewU128FromU32(x U32) U128 {
	return U128{inner: new(big.Int).SetUint64(uint64(x.Uint32()))}
}

func NewU128FromU64(x U64) U128 {
	return U128{inner: new(big.Int).SetUint64(x.Uint64())}
}

func NewU128FromUint8(x uint8) U128 {
	return U128{inner: new(big.Int).SetUint64(uint64(x))}
}

func NewU128FromUint16(x uint16) U128 {
	return U128{inner: new(big.Int).SetUint64(uint64(x))}
}

func NewU128FromUint32(x uint32) U128 {
	return U128{inner: new(big.Int).SetUint64(uint64(x))}
}

func NewU128FromUint64(x uint64) U128 {
	return U128{inner: new(big.Int).SetUint64(x)}
}

func NewU128FromInt(x *big.Int) U128 {
	if x.Sign() == -1 {
		panic("underflow")
	}
	if x.Cmp(MaxU128.inner) > 0 {
		panic("overflow")
	}
	return U128{inner: new(big.Int).Set(x)}
}

// Type returns the type of this value.
func (U128) Type() Type {
	return typeU128{}
}

func (u128 U128) Int() *big.Int {
	return new(big.Int).Set(u128.inner)
}

func (u128 U128) Bytes16() [16]byte {
	return paddedTo16(u128.inner)
}

func (u128 U128) Bytes() []byte {
	bytes := paddedTo16(u128.inner)
	return bytes[:]
}

func (u128 U128) Add(other U128) U128 {
	if u128.inner == nil {
		u128.inner = new(big.Int)
	}
	ret := U128{inner: new(big.Int)}
	if other.inner == nil {
		ret.inner.Set(u128.inner)
	} else {
		ret.inner.Add(u128.inner, other.inner)
	}
	if ret.inner.Cmp(MaxU128.inner) >= 0 {
		panic("overflow")
	}
	return ret
}

func (u128 U128) Sub(other U128) U128 {
	if u128.inner == nil {
		u128.inner = new(big.Int)
	}
	ret := U128{inner: new(big.Int)}
	if other.inner == nil {
		ret.inner.Set(u128.inner)
	} else {
		ret.inner.Sub(u128.inner, other.inner)
	}
	if ret.inner.Sign() == -1 {
		panic("underflow")
	}
	return ret
}

func (u128 U128) Mul(other U128) U128 {
	if u128.inner == nil {
		u128.inner = new(big.Int)
	}
	ret := U128{inner: new(big.Int)}
	if other.inner == nil {
		ret.inner.Set(u128.inner)
	} else {
		ret.inner.Mul(u128.inner, other.inner)
	}
	if ret.inner.Cmp(MaxU128.inner) >= 0 {
		panic("overflow")
	}
	return ret
}

func (u128 U128) Div(other U128) U128 {
	if u128.inner == nil {
		u128.inner = new(big.Int)
	}
	ret := U128{inner: new(big.Int)}
	if other.inner == nil {
		ret.inner.Set(u128.inner)
	} else {
		ret.inner.Div(u128.inner, other.inner)
	}
	return ret
}

func (u128 *U128) AddAssign(other U128) {
	if other.inner == nil {
		return
	}
	if u128.inner == nil {
		u128.inner = new(big.Int)
	}
	u128.inner.Add(u128.inner, other.inner)
	if u128.inner.Cmp(MaxU128.inner) > 0 {
		panic("overflow")
	}
}

func (u128 *U128) SubAssign(other U128) {
	if other.inner == nil {
		return
	}
	if u128.inner == nil {
		u128.inner = new(big.Int)
	}
	u128.inner.Sub(u128.inner, other.inner)
	if u128.inner.Sign() == -1 {
		panic("underflow")
	}
}

func (u128 U128) Equal(other U128) bool {
	if u128.inner == nil {
		return other.inner == nil || other.inner.Sign() == 0
	}
	if other.inner == nil {
		return u128.inner == nil || u128.inner.Sign() == 0
	}
	return u128.inner.Cmp(other.inner) == 0
}

func (u128 U128) LessThan(other U128) bool {
	if u128.inner == nil {
		return other.inner != nil && other.inner.Sign() > 0
	}
	if other.inner == nil {
		return u128.inner != nil && u128.inner.Sign() < 0
	}
	return u128.inner.Cmp(other.inner) < 0
}

func (u128 U128) LessThanEqual(other U128) bool {
	if u128.inner == nil {
		return other.inner != nil && other.inner.Sign() >= 0
	}
	if other.inner == nil {
		return u128.inner != nil && u128.inner.Sign() <= 0
	}
	return u128.inner.Cmp(other.inner) <= 0
}

func (u128 U128) GreaterThan(other U128) bool {
	if u128.inner == nil {
		return other.inner != nil && other.inner.Sign() < 0
	}
	if other.inner == nil {
		return u128.inner != nil && u128.inner.Sign() > 0
	}
	return u128.inner.Cmp(other.inner) > 0
}

func (u128 U128) GreaterThanEqual(other U128) bool {
	if u128.inner == nil {
		return other.inner != nil && other.inner.Sign() <= 0
	}
	if other.inner == nil {
		return u128.inner != nil && u128.inner.Sign() >= 0
	}
	return u128.inner.Cmp(other.inner) >= 0
}

func (u128 U128) SizeHint() int {
	return 16
}

func (u128 U128) Marshal(buf []byte, rem int) ([]byte, int, error) {
	if len(buf) < 16 || rem < 16 {
		return buf, rem, surge.ErrUnexpectedEndOfBuffer
	}
	b16 := paddedTo16(u128.inner)
	copy(buf, b16[:])
	return buf[16:], rem - 16, nil
}

func (u128 *U128) Unmarshal(buf []byte, rem int) ([]byte, int, error) {
	if len(buf) < 16 || rem < 16 {
		return buf, rem, surge.ErrUnexpectedEndOfBuffer
	}
	b16 := [16]byte{}
	copy(b16[:], buf[:16])
	if u128.inner == nil {
		u128.inner = new(big.Int)
	}
	u128.inner.SetBytes(b16[:])
	return buf[16:], rem - 16, nil
}

func (u128 U128) MarshalJSON() ([]byte, error) {
	return json.Marshal(u128.String())
}

func (u128 *U128) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	if u128.inner == nil {
		u128.inner = new(big.Int)
	}
	_, ok := u128.inner.SetString(str, 10)
	if !ok {
		return fmt.Errorf("malformed: %v", str)
	}
	if u128.inner.Sign() == -1 {
		return fmt.Errorf("underflow: %v", str)
	}
	if u128.inner.Cmp(MaxU128.inner) > 0 {
		return fmt.Errorf("overflow: %v", str)
	}
	return nil
}

func (u128 U128) String() string {
	return u128.inner.Text(10)
}

type U256 struct {
	inner *big.Int
}

func NewU256(x [32]byte) U256 {
	return U256{inner: new(big.Int).SetBytes(x[:])}
}

func NewU256FromU8(x U8) U256 {
	return U256{inner: new(big.Int).SetUint64(uint64(x.Uint8()))}
}

func NewU256FromU16(x U16) U256 {
	return U256{inner: new(big.Int).SetUint64(uint64(x.Uint16()))}
}

func NewU256FromU32(x U32) U256 {
	return U256{inner: new(big.Int).SetUint64(uint64(x.Uint32()))}
}

func NewU256FromU64(x U64) U256 {
	return U256{inner: new(big.Int).SetUint64(x.Uint64())}
}

func NewU256FromU128(x U128) U256 {
	return NewU256FromInt(x.Int())
}

func NewU256FromUint8(x uint8) U256 {
	return U256{inner: new(big.Int).SetUint64(uint64(x))}
}

func NewU256FromUint16(x uint16) U256 {
	return U256{inner: new(big.Int).SetUint64(uint64(x))}
}

func NewU256FromUint32(x uint32) U256 {
	return U256{inner: new(big.Int).SetUint64(uint64(x))}
}

func NewU256FromUint64(x uint64) U256 {
	return U256{inner: new(big.Int).SetUint64(x)}
}

func NewU256FromInt(x *big.Int) U256 {
	if x.Sign() == -1 {
		panic("underflow")
	}
	if x.Cmp(MaxU256.inner) > 0 {
		panic("overflow")
	}
	return U256{inner: new(big.Int).Set(x)}
}

// Type returns the type of this value.
func (U256) Type() Type {
	return typeU256{}
}

func (u256 U256) Int() *big.Int {
	return new(big.Int).Set(u256.inner)
}

func (u256 U256) Bytes32() [32]byte {
	return paddedTo32(u256.inner)
}

func (u256 U256) Bytes() []byte {
	bytes := paddedTo32(u256.inner)
	return bytes[:]
}

func (u256 U256) Add(other U256) U256 {
	if u256.inner == nil {
		u256.inner = new(big.Int)
	}
	ret := U256{inner: new(big.Int)}
	if other.inner == nil {
		ret.inner.Set(u256.inner)
	} else {
		ret.inner.Add(u256.inner, other.inner)
	}
	if ret.inner.Cmp(MaxU256.inner) > 0 {
		panic("overflow")
	}
	return ret
}

func (u256 U256) Sub(other U256) U256 {
	if u256.inner == nil {
		u256.inner = new(big.Int)
	}
	ret := U256{inner: new(big.Int)}
	if other.inner == nil {
		ret.inner.Set(u256.inner)
	} else {
		ret.inner.Sub(u256.inner, other.inner)
	}
	if ret.inner.Sign() == -1 {
		panic("underflow")
	}
	return ret
}

func (u256 U256) Mul(other U256) U256 {
	if u256.inner == nil {
		u256.inner = new(big.Int)
	}
	ret := U256{inner: new(big.Int)}
	if other.inner == nil {
		ret.inner.Set(u256.inner)
	} else {
		ret.inner.Mul(u256.inner, other.inner)
	}
	if ret.inner.Cmp(MaxU256.inner) >= 0 {
		panic("overflow")
	}
	return ret
}

func (u256 U256) Div(other U256) U256 {
	if u256.inner == nil {
		u256.inner = new(big.Int)
	}
	ret := U256{inner: new(big.Int)}
	if other.inner == nil {
		ret.inner.Set(u256.inner)
	} else {
		ret.inner.Div(u256.inner, other.inner)
	}
	return ret
}

func (u256 *U256) AddAssign(other U256) {
	if other.inner == nil {
		return
	}
	if u256.inner == nil {
		u256.inner = new(big.Int)
	}
	u256.inner.Add(u256.inner, other.inner)
	if u256.inner.Cmp(MaxU256.inner) > 0 {
		panic("overflow")
	}
}

func (u256 *U256) SubAssign(other U256) {
	if other.inner == nil {
		return
	}
	if u256.inner == nil {
		u256.inner = new(big.Int)
	}
	u256.inner.Sub(u256.inner, other.inner)
	if u256.inner.Sign() == -1 {
		panic("underflow")
	}
}

func (u256 U256) Equal(other U256) bool {
	if u256.inner == nil {
		return other.inner == nil || other.inner.Sign() == 0
	}
	if other.inner == nil {
		return u256.inner == nil || u256.inner.Sign() == 0
	}
	return u256.inner.Cmp(other.inner) == 0
}

func (u256 U256) LessThan(other U256) bool {
	if u256.inner == nil {
		return other.inner != nil && other.inner.Sign() > 0
	}
	if other.inner == nil {
		return u256.inner != nil && u256.inner.Sign() < 0
	}
	return u256.inner.Cmp(other.inner) < 0
}

func (u256 U256) LessThanEqual(other U256) bool {
	if u256.inner == nil {
		return other.inner != nil && other.inner.Sign() >= 0
	}
	if other.inner == nil {
		return u256.inner != nil && u256.inner.Sign() <= 0
	}
	return u256.inner.Cmp(other.inner) <= 0
}

func (u256 U256) GreaterThan(other U256) bool {
	if u256.inner == nil {
		return other.inner != nil && other.inner.Sign() < 0
	}
	if other.inner == nil {
		return u256.inner != nil && u256.inner.Sign() > 0
	}
	return u256.inner.Cmp(other.inner) > 0
}

func (u256 U256) GreaterThanEqual(other U256) bool {
	if u256.inner == nil {
		return other.inner != nil && other.inner.Sign() <= 0
	}
	if other.inner == nil {
		return u256.inner != nil && u256.inner.Sign() >= 0
	}
	return u256.inner.Cmp(other.inner) >= 0
}

func (u256 U256) SizeHint() int {
	return 32
}

func (u256 U256) Marshal(buf []byte, rem int) ([]byte, int, error) {
	if len(buf) < 32 || rem < 32 {
		return buf, rem, surge.ErrUnexpectedEndOfBuffer
	}
	b32 := paddedTo32(u256.inner)
	copy(buf, b32[:])
	return buf[32:], rem - 32, nil
}

func (u256 *U256) Unmarshal(buf []byte, rem int) ([]byte, int, error) {
	if len(buf) < 32 || rem < 32 {
		return buf, rem, surge.ErrUnexpectedEndOfBuffer
	}
	b32 := [32]byte{}
	copy(b32[:], buf[:32])
	if u256.inner == nil {
		u256.inner = new(big.Int)
	}
	u256.inner.SetBytes(b32[:])
	return buf[32:], rem - 32, nil
}

func (u256 U256) MarshalJSON() ([]byte, error) {
	return json.Marshal(u256.String())
}

func (u256 *U256) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	if u256.inner == nil {
		u256.inner = new(big.Int)
	}
	_, ok := u256.inner.SetString(str, 10)
	if !ok {
		return fmt.Errorf("malformed: %v", str)
	}
	if u256.inner.Sign() == -1 {
		return fmt.Errorf("underflow: %v", str)
	}
	if u256.inner.Cmp(MaxU256.inner) > 0 {
		return fmt.Errorf("overflow: %v", str)
	}
	return nil
}

func (u256 U256) String() string {
	return u256.inner.Text(10)
}

// Generate a random int. This method is implemented for use in quick tests.
// See https://golang.org/pkg/testing/quick/#Generator for more information.
func (u8 U8) Generate(r *rand.Rand, size int) reflect.Value {
	v, _ := quick.Value(reflect.TypeOf(uint8(0)), r)
	return reflect.ValueOf(NewU8(v.Interface().(uint8)))
}

// Generate a random int. This method is implemented for use in quick tests.
// See https://golang.org/pkg/testing/quick/#Generator for more information.
func (u16 U16) Generate(r *rand.Rand, size int) reflect.Value {
	v, _ := quick.Value(reflect.TypeOf(uint16(0)), r)
	return reflect.ValueOf(NewU16(v.Interface().(uint16)))
}

// Generate a random int. This method is implemented for use in quick tests.
// See https://golang.org/pkg/testing/quick/#Generator for more information.
func (u32 U32) Generate(r *rand.Rand, size int) reflect.Value {
	v, _ := quick.Value(reflect.TypeOf(uint32(0)), r)
	return reflect.ValueOf(NewU32(v.Interface().(uint32)))
}

// Generate a random int. This method is implemented for use in quick tests.
// See https://golang.org/pkg/testing/quick/#Generator for more information.
func (u64 U64) Generate(r *rand.Rand, size int) reflect.Value {
	v, _ := quick.Value(reflect.TypeOf(uint64(0)), r)
	return reflect.ValueOf(NewU64(v.Interface().(uint64)))
}

// Generate a random int. This method is implemented for use in quick tests.
// See https://golang.org/pkg/testing/quick/#Generator for more information.
func (u128 U128) Generate(r *rand.Rand, size int) reflect.Value {
	v, _ := quick.Value(reflect.TypeOf([16]byte{}), r)
	return reflect.ValueOf(NewU128(v.Interface().([16]byte)))
}

// Generate a random int. This method is implemented for use in quick tests.
// See https://golang.org/pkg/testing/quick/#Generator for more information.
func (u256 U256) Generate(r *rand.Rand, size int) reflect.Value {
	v, _ := quick.Value(reflect.TypeOf([32]byte{}), r)
	return reflect.ValueOf(NewU256(v.Interface().([32]byte)))
}

// paddedTo16 encodes a big integer as a big-endian into a 16-byte array. It
// will panic if the big integer is more than 16 bytes.
// Modified from:
// https://github.com/ethereum/go-ethereum/blob/master/common/math/big.go
// 17381ecc6695ea9c2d8e5ee0aee5cf70d59a301a
func paddedTo16(bigint *big.Int) [16]byte {
	if bigint.BitLen()/8 > 16 {
		panic(fmt.Sprintf("too big: expected n<16, got n=%v", bigint.BitLen()/8))
	}
	ret := [16]byte{}
	readBits(bigint, ret[:])
	return ret
}

// paddedTo32 encodes a big integer as a big-endian into a 32-byte array. It
// will panic if the big integer is more than 32 bytes.
// Modified from:
// https://github.com/ethereum/go-ethereum/blob/master/common/math/big.go
// 17381ecc6695ea9c2d8e5ee0aee5cf70d59a301a
func paddedTo32(bigint *big.Int) [32]byte {
	if bigint.BitLen()/8 > 32 {
		panic(fmt.Sprintf("too big: expected n<32, got n=%v", bigint.BitLen()/8))
	}
	ret := [32]byte{}
	readBits(bigint, ret[:])
	return ret
}

// readBits encodes the absolute value of bigint as big-endian bytes. Callers
// must ensure that buf has enough space. If buf is too short the result will be
// incomplete.
// Modified from:
// https://github.com/ethereum/go-ethereum/blob/master/common/math/big.go
// 17381ecc6695ea9c2d8e5ee0aee5cf70d59a301a
func readBits(bigint *big.Int, buf []byte) {
	i := len(buf)
	for _, d := range bigint.Bits() {
		for j := 0; j < wordBytes && i > 0; j++ {
			i--
			buf[i] = byte(d)
			d >>= 8
		}
	}
}

const (
	// wordBits is the number of bits in a big word.
	wordBits = 32 << (uint64(^big.Word(0)) >> 63)
	// wordBytes is the number of bytes in a big word.
	wordBytes = wordBits / 8
)

// Maximum values for unsigned integers.
var (
	MaxU8 = func() U8 {
		return U8(255)
	}()
	MaxU16 = func() U16 {
		return U16(65535)
	}()
	MaxU32 = func() U32 {
		return U32(4294967295)
	}()
	MaxU64 = func() U64 {
		return U64(18446744073709551615)
	}()
	MaxU128 = func() U128 {
		x, _ := new(big.Int).SetString("340282366920938463463374607431768211455", 10)
		return U128{inner: x}
	}()
	MaxU256 = func() U256 {
		x, _ := new(big.Int).SetString("115792089237316195423570985008687907853269984665640564039457584007913129639935", 10)
		return U256{inner: x}
	}()
)

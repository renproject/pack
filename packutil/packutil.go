package packutil

import (
	"encoding/json"
	"fmt"
	"math/big"
	"math/rand"
	"reflect"
	"testing/quick"
	"time"

	"github.com/renproject/pack"
)

// JSONFuzz is the same as the Fuzz testing function exposed by surge, but it
// uses JSON.
func JSONFuzz(t reflect.Type) {
	// Fuzz data
	data, ok := quick.Value(reflect.TypeOf([]byte{}), rand.New(rand.NewSource(time.Now().UnixNano())))
	if !ok {
		panic(fmt.Errorf("cannot generate value of type %v", t))
	}
	// Unmarshal
	x := reflect.New(t)
	if err := json.Unmarshal(data.Bytes(), x.Interface()); err != nil {
		// Ignore the error, because we are only interested in whether or not
		// the unmarshaling causes a panic.
	}
}

// JSONMarshalUnmarshalCheck is the same as the MarshalUnmarshalCheck testing
// function exposed by surge, but it uses JSON.
func JSONMarshalUnmarshalCheck(t reflect.Type) error {
	// Generate
	x, ok := quick.Value(t, rand.New(rand.NewSource(time.Now().UnixNano())))
	if !ok {
		return fmt.Errorf("cannot generate value of type %v", t)
	}
	// Marshal
	data, err := json.Marshal(x.Interface())
	if err != nil {
		return fmt.Errorf("cannot marshal: %v", err)
	}
	// Unmarshal
	y := reflect.New(t)
	if err := json.Unmarshal(data, y.Interface()); err != nil {
		return fmt.Errorf("cannot unmarshal: %v", err)
	}
	// Equality
	if !reflect.DeepEqual(x.Interface(), y.Elem().Interface()) {
		return fmt.Errorf("unequal")
	}
	return nil
}

// AddZeroCheck generates a random instances of the integer type and checks that
// it does not change after adding zero.
func AddZeroCheck(t reflect.Type) error {
	// Generate
	x, ok := quick.Value(t, rand.New(rand.NewSource(time.Now().UnixNano())))
	if !ok {
		return fmt.Errorf("cannot generate value of type %v", t)
	}
	// Add zero
	y := reflect.Zero(t)
	result := x.MethodByName("Add").Call([]reflect.Value{y})
	result = result[0].MethodByName("Equal").Call([]reflect.Value{x})
	if !result[0].Bool() {
		return fmt.Errorf("unequal after adding zero")
	}
	result = y.MethodByName("Add").Call([]reflect.Value{x})
	result = result[0].MethodByName("Equal").Call([]reflect.Value{x})
	if !result[0].Bool() {
		return fmt.Errorf("unequal after adding zero")
	}
	// AddAssign zero
	before := x.Interface().(fmt.Stringer).String()
	xPtr := reflect.New(t)
	xPtr.Elem().Set(x)
	xPtr.MethodByName("AddAssign").Call([]reflect.Value{reflect.Zero(t)})
	if xPtr.Elem().Interface().(fmt.Stringer).String() != before {
		return fmt.Errorf("unequal after add assigning zero")
	}
	return nil
}

// SubZeroCheck generates a random instances of the integer type and checks that
// it does not change after subtracting zero.
func SubZeroCheck(t reflect.Type) error {
	// Generate
	x, ok := quick.Value(t, rand.New(rand.NewSource(time.Now().UnixNano())))
	if !ok {
		return fmt.Errorf("cannot generate value of type %v", t)
	}
	// Sub zero
	y := reflect.Zero(t)
	result := x.MethodByName("Sub").Call([]reflect.Value{y})
	result = result[0].MethodByName("Equal").Call([]reflect.Value{x})
	if !result[0].Bool() {
		return fmt.Errorf("unequal after subtracting zero")
	}
	// SubAssign zero
	before := x.Interface().(fmt.Stringer).String()
	xPtr := reflect.New(t)
	xPtr.Elem().Set(x)
	xPtr.MethodByName("SubAssign").Call([]reflect.Value{reflect.Zero(t)})
	if xPtr.Elem().Interface().(fmt.Stringer).String() != before {
		return fmt.Errorf("unequal after sub assigning zero")
	}
	return nil
}

// AddOverflow generates two random instances of the integer type that are
// guaranteed to cause an overflow when added, and adds them. This function is
// expected to panic.
func AddOverflow(t reflect.Type, assign bool) {
	// Generate
	var x, y reflect.Value
	switch t {
	case reflect.TypeOf(pack.U8(0)):
		x = reflect.ValueOf(pack.NewU8(uint8(255)))
		y = reflect.ValueOf(pack.NewU8(uint8(1)))
	case reflect.TypeOf(pack.U16(0)):
		x = reflect.ValueOf(pack.NewU16(uint16(65535)))
		y = reflect.ValueOf(pack.NewU16(uint16(1)))
	case reflect.TypeOf(pack.U32(0)):
		x = reflect.ValueOf(pack.NewU32(uint32(4294967295)))
		y = reflect.ValueOf(pack.NewU32(uint32(1)))
	case reflect.TypeOf(pack.U64(0)):
		x = reflect.ValueOf(pack.NewU64(uint64(18446744073709551615)))
		y = reflect.ValueOf(pack.NewU64(uint64(1)))
	case reflect.TypeOf(pack.U128{}):
		x = reflect.ValueOf(pack.NewU128([16]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}))
		y = reflect.ValueOf(pack.NewU128([16]byte{1}))
	case reflect.TypeOf(pack.U256{}):
		x = reflect.ValueOf(pack.NewU256([32]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}))
		y = reflect.ValueOf(pack.NewU256([32]byte{1}))
	default:
		// Do not panic, which should cause the test to fail.
		return
	}
	// Add with overflow
	if assign {
		xPtr := reflect.New(t)
		xPtr.Elem().Set(x)
		xPtr.MethodByName("AddAssign").Call([]reflect.Value{y})
		return
	}
	x.MethodByName("Add").Call([]reflect.Value{y})
}

// SubUnderflow generates two random instances of the integer type that are
// guaranteed to cause an underflow when subtracted, and subtracts them. This
// function is expected to panic.
func SubUnderflow(t reflect.Type, assign bool) {
	// Generate
	var x, y reflect.Value
	switch t {
	case reflect.TypeOf(pack.U8(0)):
		x = reflect.ValueOf(pack.NewU8(uint8(0)))
		y = reflect.ValueOf(pack.NewU8(uint8(1)))
	case reflect.TypeOf(pack.U16(0)):
		x = reflect.ValueOf(pack.NewU16(uint16(0)))
		y = reflect.ValueOf(pack.NewU16(uint16(1)))
	case reflect.TypeOf(pack.U32(0)):
		x = reflect.ValueOf(pack.NewU32(uint32(0)))
		y = reflect.ValueOf(pack.NewU32(uint32(1)))
	case reflect.TypeOf(pack.U64(0)):
		x = reflect.ValueOf(pack.NewU64(uint64(0)))
		y = reflect.ValueOf(pack.NewU64(uint64(1)))
	case reflect.TypeOf(pack.U128{}):
		x = reflect.ValueOf(pack.NewU128([16]byte{}))
		y = reflect.ValueOf(pack.NewU128([16]byte{1}))
	case reflect.TypeOf(pack.U256{}):
		x = reflect.ValueOf(pack.NewU256([32]byte{}))
		y = reflect.ValueOf(pack.NewU256([32]byte{1}))
	default:
		// Do not panic, which should cause the test to fail.
		return
	}
	// Sub with overflow
	if assign {
		xPtr := reflect.New(t)
		xPtr.Elem().Set(x)
		xPtr.MethodByName("SubAssign").Call([]reflect.Value{y})
		return
	}
	x.MethodByName("Sub").Call([]reflect.Value{y})
}

// AddSubCheck generates two random instances of the integer type, adds y to x
// and checks that the result is not x, and then subtracts y from x and checks
// that the result is x again.
func AddSubCheck(t reflect.Type) error {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Generate
	var x, y reflect.Value
GENERATE:
	switch t {
	case reflect.TypeOf(pack.U8(0)):
		x = reflect.ValueOf(pack.NewU8(uint8(r.Int()) / 2))
		y = reflect.ValueOf(pack.NewU8(uint8(r.Int()) / 2))
	case reflect.TypeOf(pack.U16(0)):
		x = reflect.ValueOf(pack.NewU16(uint16(r.Int()) / 2))
		y = reflect.ValueOf(pack.NewU16(uint16(r.Int()) / 2))
	case reflect.TypeOf(pack.U32(0)):
		x = reflect.ValueOf(pack.NewU32(uint32(r.Int()) / 2))
		y = reflect.ValueOf(pack.NewU32(uint32(r.Int()) / 2))
	case reflect.TypeOf(pack.U64(0)):
		x = reflect.ValueOf(pack.NewU64(uint64(r.Int()) / 2))
		y = reflect.ValueOf(pack.NewU64(uint64(r.Int()) / 2))
	case reflect.TypeOf(pack.U128{}):
		x = reflect.ValueOf(pack.NewU128FromInt(big.NewInt(r.Int63())))
		y = reflect.ValueOf(pack.NewU128FromInt(big.NewInt(r.Int63())))
	case reflect.TypeOf(pack.U256{}):
		x = reflect.ValueOf(pack.NewU256FromInt(big.NewInt(r.Int63())))
		y = reflect.ValueOf(pack.NewU256FromInt(big.NewInt(r.Int63())))
	}
	if x.Interface().(fmt.Stringer).String() == y.Interface().(fmt.Stringer).String() || y.Interface().(fmt.Stringer).String() == "0" {
		// Comparing strings is the easiest way to check that we have not
		// generated the same value twice.
		goto GENERATE
	}

	// Add then sub
	{
		before := x.Interface().(fmt.Stringer).String()
		// Add
		z := x.MethodByName("Add").Call([]reflect.Value{y})[0]
		if z.Interface().(fmt.Stringer).String() == before {
			return fmt.Errorf("equal after adding")
		}
		// Sub
		z = z.MethodByName("Sub").Call([]reflect.Value{y})[0]
		if z.Interface().(fmt.Stringer).String() != before {
			return fmt.Errorf("unequal after adding then subtracting")
		}
	}

	// AddAssign then SubAssign
	{
		before := x.Interface().(fmt.Stringer).String()
		// Add assign
		xPtr := reflect.New(t)
		xPtr.Elem().Set(x)
		xPtr.MethodByName("AddAssign").Call([]reflect.Value{y})
		if xPtr.Elem().Interface().(fmt.Stringer).String() == before {
			return fmt.Errorf("equal after add assigning")
		}
		// Sub assign
		xPtr.MethodByName("SubAssign").Call([]reflect.Value{y})
		if xPtr.Elem().Interface().(fmt.Stringer).String() != before {
			return fmt.Errorf("not equal after add assigning and then sub assigning")
		}
	}
	return nil
}

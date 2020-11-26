package pack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"reflect"
	"testing/quick"

	"github.com/renproject/surge"
)

// A Type is a concrete type definition for a value.
type Type interface {
	surge.Marshaler

	// Kind returns the abstract "type of the type" for this type.
	Kind() Kind

	// Equals returns a boolean indicating whether or not the given type is
	// identical to the current one.
	Equals(other Type) bool

	// Unmarshal a value of this type from binary.
	UnmarshalValue(buf []byte, rem int) (Value, []byte, int, error)

	// Unmarshal a value of this type from JSON.
	UnmarshalValueJSON(data []byte) (Value, error)
}

type typeBool struct{}

func (typeBool) Kind() Kind {
	return KindBool
}

func (t typeBool) Equals(other Type) bool {
	a, err := surge.ToBinary(t)
	if err != nil {
		return false
	}
	b, err := surge.ToBinary(other)
	if err != nil {
		return false
	}
	return bytes.Equal(a, b)
}

func (typeBool) UnmarshalValue(buf []byte, rem int) (Value, []byte, int, error) {
	value := Bool(false)
	buf, rem, err := value.Unmarshal(buf, rem)
	return value, buf, rem, err
}

func (typeBool) UnmarshalValueJSON(data []byte) (Value, error) {
	value := Bool(false)
	err := value.UnmarshalJSON(data)
	return value, err
}

func (t typeBool) SizeHint() int {
	return t.Kind().SizeHint()
}

func (t typeBool) Marshal(buf []byte, rem int) ([]byte, int, error) {
	return t.Kind().Marshal(buf, rem)
}

func (t typeBool) MarshalText() ([]byte, error) {
	return t.Kind().MarshalText()
}

type typeU8 struct{}

func (typeU8) Kind() Kind {
	return KindU8
}

func (t typeU8) Equals(other Type) bool {
	a, err := surge.ToBinary(t)
	if err != nil {
		return false
	}
	b, err := surge.ToBinary(other)
	if err != nil {
		return false
	}
	return bytes.Equal(a, b)
}

func (typeU8) UnmarshalValue(buf []byte, rem int) (Value, []byte, int, error) {
	value := U8(0)
	buf, rem, err := value.Unmarshal(buf, rem)
	return value, buf, rem, err
}

func (typeU8) UnmarshalValueJSON(data []byte) (Value, error) {
	value := U8(0)
	err := value.UnmarshalJSON(data)
	return value, err
}

func (t typeU8) SizeHint() int {
	return t.Kind().SizeHint()
}

func (t typeU8) Marshal(buf []byte, rem int) ([]byte, int, error) {
	return t.Kind().Marshal(buf, rem)
}

func (t typeU8) MarshalText() ([]byte, error) {
	return t.Kind().MarshalText()
}

type typeU16 struct{}

func (typeU16) Kind() Kind {
	return KindU16
}

func (t typeU16) Equals(other Type) bool {
	a, err := surge.ToBinary(t)
	if err != nil {
		return false
	}
	b, err := surge.ToBinary(other)
	if err != nil {
		return false
	}
	return bytes.Equal(a, b)
}

func (typeU16) UnmarshalValue(buf []byte, rem int) (Value, []byte, int, error) {
	value := U16(0)
	buf, rem, err := value.Unmarshal(buf, rem)
	return value, buf, rem, err
}

func (typeU16) UnmarshalValueJSON(data []byte) (Value, error) {
	value := U16(0)
	err := value.UnmarshalJSON(data)
	return value, err
}

func (t typeU16) SizeHint() int {
	return t.Kind().SizeHint()
}

func (t typeU16) Marshal(buf []byte, rem int) ([]byte, int, error) {
	return t.Kind().Marshal(buf, rem)
}

func (t typeU16) MarshalText() ([]byte, error) {
	return t.Kind().MarshalText()
}

type typeU32 struct{}

func (typeU32) Kind() Kind {
	return KindU32
}

func (t typeU32) Equals(other Type) bool {
	a, err := surge.ToBinary(t)
	if err != nil {
		return false
	}
	b, err := surge.ToBinary(other)
	if err != nil {
		return false
	}
	return bytes.Equal(a, b)
}

func (typeU32) UnmarshalValue(buf []byte, rem int) (Value, []byte, int, error) {
	value := U32(0)
	buf, rem, err := value.Unmarshal(buf, rem)
	return value, buf, rem, err
}

func (typeU32) UnmarshalValueJSON(data []byte) (Value, error) {
	value := U32(0)
	err := value.UnmarshalJSON(data)
	return value, err
}

func (t typeU32) SizeHint() int {
	return t.Kind().SizeHint()
}

func (t typeU32) Marshal(buf []byte, rem int) ([]byte, int, error) {
	return t.Kind().Marshal(buf, rem)
}

func (t typeU32) MarshalText() ([]byte, error) {
	return t.Kind().MarshalText()
}

type typeU64 struct{}

func (typeU64) Kind() Kind {
	return KindU64
}

func (t typeU64) Equals(other Type) bool {
	a, err := surge.ToBinary(t)
	if err != nil {
		return false
	}
	b, err := surge.ToBinary(other)
	if err != nil {
		return false
	}
	return bytes.Equal(a, b)
}

func (typeU64) UnmarshalValue(buf []byte, rem int) (Value, []byte, int, error) {
	value := U64(0)
	buf, rem, err := value.Unmarshal(buf, rem)
	return value, buf, rem, err
}

func (typeU64) UnmarshalValueJSON(data []byte) (Value, error) {
	value := U64(0)
	err := value.UnmarshalJSON(data)
	return value, err
}

func (t typeU64) SizeHint() int {
	return t.Kind().SizeHint()
}

func (t typeU64) Marshal(buf []byte, rem int) ([]byte, int, error) {
	return t.Kind().Marshal(buf, rem)
}

func (t typeU64) MarshalText() ([]byte, error) {
	return t.Kind().MarshalText()
}

type typeU128 struct{}

func (typeU128) Kind() Kind {
	return KindU128
}

func (t typeU128) Equals(other Type) bool {
	a, err := surge.ToBinary(t)
	if err != nil {
		return false
	}
	b, err := surge.ToBinary(other)
	if err != nil {
		return false
	}
	return bytes.Equal(a, b)
}

func (typeU128) UnmarshalValue(buf []byte, rem int) (Value, []byte, int, error) {
	value := U128{}
	buf, rem, err := value.Unmarshal(buf, rem)
	return value, buf, rem, err
}

func (typeU128) UnmarshalValueJSON(data []byte) (Value, error) {
	value := U128{}
	err := value.UnmarshalJSON(data)
	return value, err
}

func (t typeU128) SizeHint() int {
	return t.Kind().SizeHint()
}

func (t typeU128) Marshal(buf []byte, rem int) ([]byte, int, error) {
	return t.Kind().Marshal(buf, rem)
}

func (t typeU128) MarshalText() ([]byte, error) {
	return t.Kind().MarshalText()
}

type typeU256 struct{}

func (typeU256) Kind() Kind {
	return KindU256
}

func (t typeU256) Equals(other Type) bool {
	a, err := surge.ToBinary(t)
	if err != nil {
		return false
	}
	b, err := surge.ToBinary(other)
	if err != nil {
		return false
	}
	return bytes.Equal(a, b)
}

func (typeU256) UnmarshalValue(buf []byte, rem int) (Value, []byte, int, error) {
	value := U256{}
	buf, rem, err := value.Unmarshal(buf, rem)
	return value, buf, rem, err
}

func (typeU256) UnmarshalValueJSON(data []byte) (Value, error) {
	value := U256{}
	err := value.UnmarshalJSON(data)
	return value, err
}

func (t typeU256) SizeHint() int {
	return t.Kind().SizeHint()
}

func (t typeU256) Marshal(buf []byte, rem int) ([]byte, int, error) {
	return t.Kind().Marshal(buf, rem)
}

func (t typeU256) MarshalText() ([]byte, error) {
	return t.Kind().MarshalText()
}

type typeString struct{}

func (typeString) Kind() Kind {
	return KindString
}

func (t typeString) Equals(other Type) bool {
	a, err := surge.ToBinary(t)
	if err != nil {
		return false
	}
	b, err := surge.ToBinary(other)
	if err != nil {
		return false
	}
	return bytes.Equal(a, b)
}

func (typeString) UnmarshalValue(buf []byte, rem int) (Value, []byte, int, error) {
	value := String("")
	buf, rem, err := value.Unmarshal(buf, rem)
	return value, buf, rem, err
}

func (typeString) UnmarshalValueJSON(data []byte) (Value, error) {
	value := String("")
	err := value.UnmarshalJSON(data)
	return value, err
}

func (t typeString) SizeHint() int {
	return t.Kind().SizeHint()
}

func (t typeString) Marshal(buf []byte, rem int) ([]byte, int, error) {
	return t.Kind().Marshal(buf, rem)
}

func (t typeString) MarshalText() ([]byte, error) {
	return t.Kind().MarshalText()
}

type typeBytes struct{}

func (typeBytes) Kind() Kind {
	return KindBytes
}

func (t typeBytes) Equals(other Type) bool {
	a, err := surge.ToBinary(t)
	if err != nil {
		return false
	}
	b, err := surge.ToBinary(other)
	if err != nil {
		return false
	}
	return bytes.Equal(a, b)
}

func (typeBytes) UnmarshalValue(buf []byte, rem int) (Value, []byte, int, error) {
	value := Bytes{}
	buf, rem, err := value.Unmarshal(buf, rem)
	return value, buf, rem, err
}

func (typeBytes) UnmarshalValueJSON(data []byte) (Value, error) {
	value := Bytes{}
	err := value.UnmarshalJSON(data)
	return value, err
}

func (t typeBytes) SizeHint() int {
	return t.Kind().SizeHint()
}

func (t typeBytes) Marshal(buf []byte, rem int) ([]byte, int, error) {
	return t.Kind().Marshal(buf, rem)
}

func (t typeBytes) MarshalText() ([]byte, error) {
	return t.Kind().MarshalText()
}

type typeBytes32 struct{}

func (typeBytes32) Kind() Kind {
	return KindBytes32
}

func (t typeBytes32) Equals(other Type) bool {
	a, err := surge.ToBinary(t)
	if err != nil {
		return false
	}
	b, err := surge.ToBinary(other)
	if err != nil {
		return false
	}
	return bytes.Equal(a, b)
}

func (typeBytes32) UnmarshalValue(buf []byte, rem int) (Value, []byte, int, error) {
	value := Bytes32{}
	buf, rem, err := value.Unmarshal(buf, rem)
	return value, buf, rem, err
}

func (typeBytes32) UnmarshalValueJSON(data []byte) (Value, error) {
	value := Bytes32{}
	err := value.UnmarshalJSON(data)
	return value, err
}

func (t typeBytes32) SizeHint() int {
	return t.Kind().SizeHint()
}

func (t typeBytes32) Marshal(buf []byte, rem int) ([]byte, int, error) {
	return t.Kind().Marshal(buf, rem)
}

func (t typeBytes32) MarshalText() ([]byte, error) {
	return t.Kind().MarshalText()
}

type typeBytes65 struct{}

func (typeBytes65) Kind() Kind {
	return KindBytes65
}

func (t typeBytes65) Equals(other Type) bool {
	a, err := surge.ToBinary(t)
	if err != nil {
		return false
	}
	b, err := surge.ToBinary(other)
	if err != nil {
		return false
	}
	return bytes.Equal(a, b)
}

func (typeBytes65) UnmarshalValue(buf []byte, rem int) (Value, []byte, int, error) {
	value := Bytes65{}
	buf, rem, err := value.Unmarshal(buf, rem)
	return value, buf, rem, err
}

func (typeBytes65) UnmarshalValueJSON(data []byte) (Value, error) {
	value := Bytes65{}
	err := value.UnmarshalJSON(data)
	return value, err
}

func (t typeBytes65) SizeHint() int {
	return t.Kind().SizeHint()
}

func (t typeBytes65) Marshal(buf []byte, rem int) ([]byte, int, error) {
	return t.Kind().Marshal(buf, rem)
}

func (t typeBytes65) MarshalText() ([]byte, error) {
	return t.Kind().MarshalText()
}

type typeStructField struct {
	Name string
	Type Type
}

func (field typeStructField) SizeHint() int {
	return surge.SizeHintString(field.Name) + SizeHintType(field.Type)
}

func (field typeStructField) Marshal(buf []byte, rem int) ([]byte, int, error) {
	var err error
	if buf, rem, err = surge.MarshalString(field.Name, buf, rem); err != nil {
		return buf, rem, err
	}
	if buf, rem, err = MarshalType(field.Type, buf, rem); err != nil {
		return buf, rem, err
	}
	return buf, rem, nil
}

func (field *typeStructField) Unmarshal(buf []byte, rem int) ([]byte, int, error) {
	var err error
	if buf, rem, err = surge.UnmarshalString(&field.Name, buf, rem); err != nil {
		return buf, rem, err
	}
	if buf, rem, err = UnmarshalType(&field.Type, buf, rem); err != nil {
		return buf, rem, err
	}
	return buf, rem, err
}

func (field typeStructField) MarshalJSON() ([]byte, error) {
	raw, err := marshalTypeJSON(field.Type)
	if err != nil {
		return raw, err
	}
	return json.Marshal(map[string]json.RawMessage{field.Name: raw})
}

func (field *typeStructField) UnmarshalJSON(data []byte) error {
	raw := map[string]json.RawMessage{}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	if len(raw) != 1 {
		return fmt.Errorf("expected len=1, got len=%v", len(raw))
	}
	for name, data := range raw {
		field.Name = name
		innerType, err := unmarshalTypeJSON(data)
		if err != nil {
			return err
		}
		field.Type = innerType
		return nil
	}
	panic("unreachable")
}

func (typeStructField) Generate(r *rand.Rand, size int) reflect.Value {
	name, _ := quick.Value(reflect.TypeOf(""), r)
	for {
		return reflect.ValueOf(typeStructField{
			Name: name.String(),
			Type: typeU8{}, // FIXME: Generate actual type (no struct).
		})
	}
}

type typeStruct []typeStructField

func (typeStruct) Kind() Kind {
	return KindStruct
}

func (t typeStruct) Equals(other Type) bool {
	a, err := surge.ToBinary(t)
	if err != nil {
		return false
	}
	b, err := surge.ToBinary(other)
	if err != nil {
		return false
	}
	return bytes.Equal(a, b)
}

func (t typeStruct) UnmarshalValue(buf []byte, rem int) (Value, []byte, int, error) {
	v := Struct{}
	for _, field := range t {
		var err error
		var value Value
		if value, buf, rem, err = field.Type.UnmarshalValue(buf, rem); err != nil {
			return nil, buf, rem, fmt.Errorf("unmarshaling value \"%v\": %v", field.Name, err)
		}
		v = append(v, StructField{Name: field.Name, Value: value})
	}
	return v, buf, rem, nil
}

func (t typeStruct) UnmarshalValueJSON(data []byte) (Value, error) {
	raw := map[string]json.RawMessage{}
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}
	v := Struct{}
	for _, field := range t {
		rawValue, ok := raw[field.Name]
		if !ok {
			return nil, fmt.Errorf("unmarshaling value \"%v\": not found", field.Name)
		}
		value, err := field.Type.UnmarshalValueJSON(rawValue)
		if err != nil {
			return nil, fmt.Errorf("unmarshaling value \"%v\": %v", field.Name, err)
		}
		v = append(v, StructField{Name: field.Name, Value: value})
	}
	return v, nil
}

func (t typeStruct) SizeHint() int {
	total := 4
	for _, field := range t {
		total += field.SizeHint()
	}
	return total
}

func (t typeStruct) Marshal(buf []byte, rem int) ([]byte, int, error) {
	var err error
	buf, rem, err = surge.MarshalU32(uint32(len(t)), buf, rem)
	if err != nil {
		return buf, rem, err
	}
	for _, field := range t {
		buf, rem, err = field.Marshal(buf, rem)
		if err != nil {
			return buf, rem, err
		}
	}
	return buf, rem, nil
}

func (t *typeStruct) Unmarshal(buf []byte, rem int) ([]byte, int, error) {
	var err error
	var numFields uint32
	buf, rem, err = surge.UnmarshalU32(&numFields, buf, rem)
	if err != nil {
		return buf, rem, err
	}
	for i := uint32(0); i < numFields; i++ {
		field := typeStructField{}
		buf, rem, err = field.Unmarshal(buf, rem)
		if err != nil {
			return buf, rem, err
		}
		*t = append(*t, field)
	}
	return buf, rem, nil
}

func (t typeStruct) MarshalJSON() ([]byte, error) {
	raw := make([]json.RawMessage, len(t))
	for i, field := range t {
		rawField, err := field.MarshalJSON()
		if err != nil {
			return nil, fmt.Errorf("cannot marshal \"%v\": %v", field.Name, err)
		}
		raw[i] = rawField
	}
	return json.Marshal(raw)
}

func (t *typeStruct) UnmarshalJSON(data []byte) error {
	raw := []json.RawMessage{}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	*t = make(typeStruct, len(raw))
	for i, rawField := range raw {
		field := typeStructField{}
		if err := field.UnmarshalJSON(rawField); err != nil {
			return fmt.Errorf("cannot unmarshal field=%v: %v", i, err)
		}
		(*t)[i] = field
	}
	return nil
}

func (typeStruct) Generate(r *rand.Rand, size int) reflect.Value {
	typeStructFields := make([]typeStructField, size)
	for i := 0; i < size; i++ {
		typeStructFields[i] = typeStructField{}.Generate(r, size).Interface().(typeStructField)
	}
	return reflect.ValueOf(typeStruct(typeStructFields))
}

type typeList struct {
	Type Type
}

func (typeList) Kind() Kind {
	return KindList
}

func (t typeList) Equals(other Type) bool {
	a, err := surge.ToBinary(t)
	if err != nil {
		return false
	}
	b, err := surge.ToBinary(other)
	if err != nil {
		return false
	}
	return bytes.Equal(a, b)
}

func (t typeList) UnmarshalValue(buf []byte, rem int) (Value, []byte, int, error) {
	var err error
	var numElems uint32
	if buf, rem, err = surge.UnmarshalU32(&numElems, buf, rem); err != nil {
		return nil, buf, rem, fmt.Errorf("unmarshaling list length: %v", err)
	}
	v := List{
		T:     t.Type,
		Elems: make([]Value, numElems),
	}
	for i := range v.Elems {
		var value Value
		if value, buf, rem, err = v.T.UnmarshalValue(buf, rem); err != nil {
			return nil, buf, rem, fmt.Errorf("unmarshaling list value: %v", err)
		}
		v.Elems[i] = value
	}
	return v, buf, rem, nil
}

func (t typeList) UnmarshalValueJSON(data []byte) (Value, error) {
	raw := []json.RawMessage{}
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}
	v := List{
		T:     t.Type,
		Elems: make([]Value, len(raw)),
	}
	for i := range v.Elems {
		value, err := v.T.UnmarshalValueJSON(raw[i])
		if err != nil {
			return nil, fmt.Errorf("unmarshaling list value: %v", err)
		}
		v.Elems[i] = value
	}
	return v, nil
}

func (t typeList) SizeHint() int {
	return t.Type.SizeHint() + 64
}

func (t typeList) Marshal(buf []byte, rem int) ([]byte, int, error) {
	return MarshalType(t.Type, buf, rem)
}

func (t *typeList) Unmarshal(buf []byte, rem int) ([]byte, int, error) {
	return UnmarshalType(&t.Type, buf, rem)
}

// SizeHintType returns the number of bytes requires to represent this type in
// binary.
func SizeHintType(t Type) int {
	switch t.Kind() {
	case KindBool:
		return t.SizeHint()
	case KindU8:
		return t.SizeHint()
	case KindU16:
		return t.SizeHint()
	case KindU32:
		return t.SizeHint()
	case KindU64:
		return t.SizeHint()
	case KindU128:
		return t.SizeHint()
	case KindU256:
		return t.SizeHint()
	case KindString:
		return t.SizeHint()
	case KindBytes:
		return t.SizeHint()
	case KindBytes32:
		return t.SizeHint()
	case KindBytes65:
		return t.SizeHint()
	case KindStruct, KindList:
		return t.Kind().SizeHint() + t.SizeHint()
	default:
		return 0
	}
}

// MarshalType to binary.
func MarshalType(t Type, buf []byte, rem int) ([]byte, int, error) {
	switch t.Kind() {
	case KindBool:
		return t.Marshal(buf, rem)
	case KindU8:
		return t.Marshal(buf, rem)
	case KindU16:
		return t.Marshal(buf, rem)
	case KindU32:
		return t.Marshal(buf, rem)
	case KindU64:
		return t.Marshal(buf, rem)
	case KindU128:
		return t.Marshal(buf, rem)
	case KindU256:
		return t.Marshal(buf, rem)
	case KindString:
		return t.Marshal(buf, rem)
	case KindBytes:
		return t.Marshal(buf, rem)
	case KindBytes32:
		return t.Marshal(buf, rem)
	case KindBytes65:
		return t.Marshal(buf, rem)
	case KindStruct, KindList:
		var err error
		if buf, rem, err = t.Kind().Marshal(buf, rem); err != nil {
			return buf, rem, err
		}
		if buf, rem, err = t.Marshal(buf, rem); err != nil {
			return buf, rem, err
		}
		return buf, rem, nil
	default:
		return buf, rem, nil
	}
}

// UnmarshalType from binary.
func UnmarshalType(t *Type, buf []byte, rem int) ([]byte, int, error) {
	var err error
	var kind Kind
	if buf, rem, err = kind.Unmarshal(buf, rem); err != nil {
		return buf, rem, err
	}
	switch kind {
	case KindBool:
		*t = typeBool{}
		return buf, rem, nil
	case KindU8:
		*t = typeU8{}
		return buf, rem, nil
	case KindU16:
		*t = typeU16{}
		return buf, rem, nil
	case KindU32:
		*t = typeU32{}
		return buf, rem, nil
	case KindU64:
		*t = typeU64{}
		return buf, rem, nil
	case KindU128:
		*t = typeU128{}
		return buf, rem, nil
	case KindU256:
		*t = typeU256{}
		return buf, rem, nil
	case KindString:
		*t = typeString{}
		return buf, rem, nil
	case KindBytes:
		*t = typeBytes{}
		return buf, rem, nil
	case KindBytes32:
		*t = typeBytes32{}
		return buf, rem, nil
	case KindBytes65:
		*t = typeBytes65{}
		return buf, rem, nil
	case KindStruct:
		ts := typeStruct{}
		if buf, rem, err = ts.Unmarshal(buf, rem); err != nil {
			return buf, rem, err
		}
		*t = ts
		return buf, rem, nil
	case KindList:
		tl := typeList{}
		if buf, rem, err = tl.Unmarshal(buf, rem); err != nil {
			return buf, rem, err
		}
		*t = tl
		return buf, rem, nil
	default:
		return buf, rem, fmt.Errorf("unsupported kind %v", kind)
	}
}

func marshalTypeJSON(t Type) ([]byte, error) {
	raw, err := json.Marshal(t)
	if err != nil {
		return raw, err
	}
	switch t.Kind() {
	case KindStruct:
		return json.Marshal(map[string]interface{}{
			"struct": json.RawMessage(raw),
		})
	case KindList:
		return json.Marshal(map[string]interface{}{
			"list": json.RawMessage(raw),
		})
	default:
		return raw, nil
	}
}

func unmarshalTypeJSON(data []byte) (Type, error) {
	// First attempt to unmarshal the type directly into a kind. If this
	// succeeds, then the type is simple, and we can return it based solely on
	// the kind. Otherwise, we are dealing with an abstract type, and need to
	// unmarshal differently.
	kind := KindNil
	if err := json.Unmarshal(data, &kind); err == nil {
		switch kind {
		case KindBool:
			return typeBool{}, nil
		case KindU8:
			return typeU8{}, nil
		case KindU16:
			return typeU16{}, nil
		case KindU32:
			return typeU32{}, nil
		case KindU64:
			return typeU64{}, nil
		case KindU128:
			return typeU128{}, nil
		case KindU256:
			return typeU256{}, nil
		case KindString:
			return typeString{}, nil
		case KindBytes:
			return typeBytes{}, nil
		case KindBytes32:
			return typeBytes32{}, nil
		case KindBytes65:
			return typeBytes65{}, nil
		default:
			return nil, fmt.Errorf("unexpected kind %v", kind)
		}
	}
	// Unmarshal into a map. We expect that there will be exactly one key: the
	// kind. The associated value will describe the type.
	raw := map[Kind]json.RawMessage{}
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, fmt.Errorf("unmarshaling kind: %v (%v)", err, string(data))
	}
	if len(raw) != 1 {
		return nil, fmt.Errorf("expected 1 kind, got %v kinds", len(raw))
	}
	for kind, data := range raw {
		switch kind {
		case KindStruct:
			t := typeStruct{}
			if err := json.Unmarshal(data, &t); err != nil {
				return nil, fmt.Errorf("unmarshaling struct: %v", err)
			}
			return t, nil
		case KindList:
			t := typeList{}
			if err := json.Unmarshal(data, &t); err != nil {
				return nil, fmt.Errorf("unmarshaling list: %v", err)
			}
			return t, nil
		default:
			return nil, fmt.Errorf("unexpected kind %v", kind)
		}
	}
	panic("unreachable")
}

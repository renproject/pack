package pack

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"reflect"
	"testing/quick"
)

// StructField represents a named field within a struct. It is not, by itself, a
// value or a type. It is only meant to be used to build structs.
type StructField struct {
	Name  string
	Value Value
}

// NewStructField returns a struct field with the given name and value.
func NewStructField(name string, value Value) StructField {
	return StructField{Name: name, Value: value}
}

// SizeHint returns the number of bytes required to represent the struct field
// in binary.
func (x StructField) SizeHint() int {
	return x.Value.SizeHint()
}

// Marshal the struct field into binary.
func (x StructField) Marshal(buf []byte, rem int) ([]byte, int, error) {
	var err error
	if buf, rem, err = x.Value.Marshal(buf, rem); err != nil {
		return buf, rem, err
	}
	return buf, rem, nil
}

// Generate a random struct field. This method is implemented for use in quick
// tests. See https://golang.org/pkg/testing/quick/#Generator for more
// information. Generated struct fields will never have structs as values.
func (x StructField) Generate(r *rand.Rand, size int) reflect.Value {
	name, _ := quick.Value(reflect.TypeOf(""), r)
	for {
		return reflect.ValueOf(StructField{
			Name:  name.String(),
			Value: Generate(r, size, false, false).Interface().(Value),
		})
	}
}

// Struct represents a structured record.
type Struct []StructField

// NewStruct returns a new struct from a slice of variadic arguments. The
// arguments are expected to be of the form ("name", value)* otherwise the
// function will panic. The same field name must not be used more than once.
//
//  x := NewStruct(
//      "foo", NewU64(42),
//      "bar", NewString("pack is awesome"),
//      "baz", NewBool(true),
//  )
//
func NewStruct(vs ...interface{}) Struct {
	structFields := make([]StructField, len(vs)/2)
	for i := range structFields {
		structFields[i] = NewStructField(
			vs[2*i+0].(string),
			vs[2*i+1].(Value),
		)
	}
	return Struct(structFields)
}

// Type returns the structured record type. This method has O(n) complexity,
// where N is the number of fields in the struct.
func (v Struct) Type() Type {
	t := make(typeStruct, 0, len(v))
	for _, field := range v {
		t = append(t, typeStructField{Name: field.Name, Type: field.Value.Type()})
	}
	return t
}

// Get a field value from the struct, given the field name. This method has O(n)
// complexity, where N is the number of fields in the struct.
func (v Struct) Get(name string) Value {
	for _, field := range v {
		if field.Name == name {
			return field.Value
		}
	}
	return nil
}

// Set a field value in the struct, given the field name. This method has O(n)
// complexity, where N is the number of fields in the struct.
func (v *Struct) Set(name string, value Value) Value {
	for i := range *v {
		if (*v)[i].Name == name {
			prev := (*v)[i].Value
			(*v)[i] = StructField{Name: name, Value: value}
			return prev
		}
	}
	return nil
}

// SizeHint returns the number of bytes required to represent the struct in
// binary.
func (v Struct) SizeHint() int {
	total := 0
	for _, field := range v {
		total += field.SizeHint()
	}
	return total
}

// Marshal the struct into binary.
func (v Struct) Marshal(buf []byte, rem int) ([]byte, int, error) {
	var err error
	for _, field := range v {
		buf, rem, err = field.Marshal(buf, rem)
		if err != nil {
			return buf, rem, err
		}
	}
	return buf, rem, nil
}

// MarshalJSON marshals the struct to JSON. This is done by marshaling the
// struct as if it was a JSON object, where each field in the struct is a field
// in the JSON object with the same name.
func (v Struct) MarshalJSON() ([]byte, error) {
	raw := map[string]interface{}{}
	for _, field := range v {
		rawField, err := json.Marshal(field.Value)
		if err != nil {
			return nil, fmt.Errorf("marshaling field \"%v\": %v", field.Name, err)
		}
		raw[field.Name] = json.RawMessage(rawField)
	}
	return json.Marshal(raw)
}

// String returns the struct in its JSON representation.
func (v Struct) String() string {
	data, err := v.MarshalJSON()
	if err != nil {
		return fmt.Sprintf(`{"error": %v}`, err)
	}
	return string(data)
}

// Generate a random struct field. This method is implemented for use in quick
// tests. See https://golang.org/pkg/testing/quick/#Generator for more
// information. Generated structs will never contain embedded structs.
func (Struct) Generate(r *rand.Rand, size int) reflect.Value {
	s := make(Struct, 0, size)
	for i := 0; i < size; i++ {
		structField := StructField{}.Generate(r, size).Interface().(StructField)
		if s.Get(structField.Name) == nil {
			s = append(s, structField)
		}
	}
	return reflect.ValueOf(s)
}

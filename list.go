package pack

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"reflect"

	"github.com/renproject/surge"
)

type List struct {
	T     Type
	Elems []Value
}

func EmptyList(t Type) List {
	return List{
		T:     t,
		Elems: []Value{},
	}
}

func NewList(vs ...Value) (List, error) {
	if len(vs) == 0 {
		return List{}, fmt.Errorf("cannot construct list with no elements")
	}

	elems := make([]Value, len(vs))
	var t Type
	for i := range vs {
		// Verify the list elements have a consistent type.
		if t == nil {
			t = vs[i].Type()
		} else if !vs[i].Type().Equals(t) {
			return List{}, fmt.Errorf("inconsistent list type: expected %v, got %v", t, vs[i].Type())
		}
		elems[i] = vs[i]
	}
	return List{
		T:     t,
		Elems: elems,
	}, nil
}

// Type returns the list type.
func (v List) Type() Type {
	return typeList{
		Type: v.T,
	}
}

// SizeHint returns the number of bytes required to represent the list in
// binary.
func (v List) SizeHint() int {
	return surge.SizeHint(v.Elems)
}

// Marshal the list into binary.
func (v List) Marshal(buf []byte, rem int) ([]byte, int, error) {
	buf, rem, err := surge.MarshalLen(uint32(len(v.Elems)), buf, rem)
	if err != nil {
		return buf, rem, err
	}
	for i := range v.Elems {
		buf, rem, err = v.Elems[i].Marshal(buf, rem)
		if err != nil {
			return buf, rem, err
		}
	}
	return buf, rem, nil
}

// Unmarshal the list from binary.
func (v *List) Unmarshal(buf []byte, rem int) ([]byte, int, error) {
	if v.T == nil {
		return buf, rem, fmt.Errorf("cannot unmarshal into list with unknown type")
	}

	var numElems uint32
	buf, rem, err := surge.UnmarshalLen(&numElems, 0, buf, rem)
	if err != nil {
		return buf, rem, err
	}

	v.Elems = make([]Value, numElems)
	for i := range v.Elems {
		v.Elems[i], buf, rem, err = v.T.UnmarshalValue(buf, rem)
		if err != nil {
			return buf, rem, err
		}

		// Ensure all elements are of the same type.
		if v.Elems[i].Type() != v.T {
			return buf, rem, fmt.Errorf("unexpected type: expected %v, got %v", v.T, v.Elems[i].Type())
		}
	}

	return buf, rem, nil
}

// MarshalJSON marshals the list to JSON.
func (v List) MarshalJSON() ([]byte, error) {
	raw := []interface{}{}
	for _, elem := range v.Elems {
		rawField, err := elem.MarshalJSON()
		if err != nil {
			return nil, fmt.Errorf("marshaling list element: %v", err)
		}
		raw = append(raw, json.RawMessage(rawField))
	}
	return json.Marshal(raw)
}

// String returns the list in its JSON representation.
func (v List) String() string {
	data, err := v.MarshalJSON()
	if err != nil {
		return err.Error()
	}
	return string(data)
}

// Generate a random list. This method is implemented for use in quick tests.
// See https://golang.org/pkg/testing/quick/#Generator for more information.
// Generated lists will never contain embedded lists.
func (List) Generate(r *rand.Rand, size int) reflect.Value {
	l := List{
		T:     Generate(r, size, false, false).Interface().(Value).Type(),
		Elems: make([]Value, 0, size),
	}
	for i := 0; i < size; i++ {
		v := GenerateFromKind(r, size, l.T.Kind(), false, false).Interface().(Value)
		l.Elems = append(l.Elems, v)
	}
	return reflect.ValueOf(l)
}

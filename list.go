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

func NewList(vs ...Value) (List, error) {
	if len(vs) == 0 {
		return List{}, nil
	}

	elems := make([]Value, len(vs))
	var t Type
	for i := range elems {
		// Verify the list elements have a consistent type.
		if t == nil {
			t = vs[i].Type()
		} else if vs[i].Type() != t {
			return List{}, fmt.Errorf("inconsistent list type: expected %v, got %v", t, elems[i].Type())
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
	return surge.Marshal(v.Elems, buf, rem)
}

// MarshalJSON marshals the list to JSON.
func (v List) MarshalJSON() ([]byte, error) {
	raw := []interface{}{}
	for _, elem := range v.Elems {
		rawField, err := json.Marshal(elem)
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

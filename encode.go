package pack

import (
	"fmt"
	"reflect"
	"strings"
)

// Encode a Go interface into a Value interface.
func Encode(v interface{}) (val Value, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("recovered: %v", err)
			return
		}
	}()

	// If the interface is already a value, then immediately return the
	// interface without modification.
	switch v := v.(type) {
	case Bool:
		return v, nil
	case U8:
		return v, nil
	case U16:
		return v, nil
	case U32:
		return v, nil
	case U64:
		return v, nil
	case U128:
		return v, nil
	case U256:
		return v, nil
	case String:
		return v, nil
	case Bytes:
		return v, nil
	case Bytes32:
		return v, nil
	case Bytes65:
		return v, nil
	case Struct:
		return v, nil
	case List:
		return v, nil
	case Typed:
		return Struct(v), nil
	case Value:
		return v, nil
	}

	// Otherwise, reflect on the kind/type of the interface, and convert it into
	// a value.
	valueOf := reflect.ValueOf(v)
	switch valueOf.Kind() {
	case reflect.Bool:
		return NewBool(valueOf.Bool()), nil
	case reflect.Uint8:
		return NewU8(uint8(valueOf.Uint())), nil
	case reflect.Uint16:
		return NewU16(uint16(valueOf.Uint())), nil
	case reflect.Uint32:
		return NewU32(uint32(valueOf.Uint())), nil
	case reflect.Uint64:
		return NewU64(valueOf.Uint()), nil
	case reflect.String:
		return NewString(valueOf.String()), nil
	case reflect.Slice:
		if valueOf.Type().Elem().Kind() == reflect.Uint8 {
			return NewBytes(valueOf.Bytes()), nil
		}
		if valueOf.Len() == 0 {
			elem := reflect.Zero(valueOf.Type().Elem()).Interface()
			val, err := Encode(elem)
			if err != nil {
				return nil, fmt.Errorf("encoding list item: %v", err)
			}
			return EmptyList(val.Type()), nil
		}
		var err error
		elems := make([]Value, valueOf.Len())
		for i := 0; i < valueOf.Len(); i++ {
			elems[i], err = Encode(valueOf.Index(i).Interface())
			if err != nil {
				return nil, fmt.Errorf("encoding list item: %v", err)
			}
		}
		return NewList(elems...)
	case reflect.Array:
		typeOf := valueOf.Type()
		if typeOf.Elem().Kind() == reflect.Uint8 {
			if typeOf.Len() == 32 {
				return valueOf.Convert(reflect.TypeOf(Bytes32{})).Interface().(Bytes32), nil
			}
			if typeOf.Len() == 65 {
				return valueOf.Convert(reflect.TypeOf(Bytes65{})).Interface().(Bytes65), nil
			}
		}
		return nil, fmt.Errorf("non-exhaustive pattern: type %T", v)
	case reflect.Struct:
		typeOf := valueOf.Type()
		n := typeOf.NumField()
		structFields := make([]StructField, 0, n)
		for i := 0; i < n; i++ {
			f := typeOf.Field(i)
			tags := strings.Split(f.Tag.Get("json"), ",")
			name := f.Name
			for _, tag := range tags {
				if tag == "-" {
					name = ""
					break
				}
				if tag != "omitempty" {
					name = tag
					break
				}
			}
			if name != "" {
				value, err := Encode(valueOf.Field(i).Interface())
				if err != nil {
					return nil, fmt.Errorf("encoding \"%v\": %v", f.Name, err)
				}
				structFields = append(structFields, NewStructField(name, value))
			}
		}
		return Struct(structFields), nil
	default:
		return nil, fmt.Errorf("non-exhaustive pattern: type %T", v)
	}
}

// Decode a Value interface into a Go interface. The Go interface must be a
// pointer.
func Decode(interf interface{}, v Value) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("recovered: %v", err)
			return
		}
	}()

	// If the interface-to-be-decoded-into is a value, then check the type of
	// the value-to-be-decoded, and assign.
	switch interf := interf.(type) {
	case *Bool:
		if v, ok := v.(Bool); ok {
			*interf = v
			return nil
		}
		return fmt.Errorf("unexpected value of type %T", v)
	case *U8:
		if v, ok := v.(U8); ok {
			*interf = v
			return nil
		}
		return fmt.Errorf("unexpected value of type %T", v)
	case *U16:
		if v, ok := v.(U16); ok {
			*interf = v
			return nil
		}
		return fmt.Errorf("unexpected value of type %T", v)
	case *U32:
		if v, ok := v.(U32); ok {
			*interf = v
			return nil
		}
		return fmt.Errorf("unexpected value of type %T", v)
	case *U64:
		if v, ok := v.(U64); ok {
			*interf = v
			return nil
		}
		return fmt.Errorf("unexpected value of type %T", v)
	case *U128:
		if v, ok := v.(U128); ok {
			*interf = v
			return nil
		}
		return fmt.Errorf("unexpected value of type %T", v)
	case *U256:
		if v, ok := v.(U256); ok {
			*interf = v
			return nil
		}
		return fmt.Errorf("unexpected value of type %T", v)
	case *String:
		if v, ok := v.(String); ok {
			*interf = v
			return nil
		}
		return fmt.Errorf("unexpected value of type %T", v)
	case *Bytes:
		if v, ok := v.(Bytes); ok {
			*interf = v
			return nil
		}
		return fmt.Errorf("unexpected value of type %T", v)
	case *Bytes32:
		if v, ok := v.(Bytes32); ok {
			*interf = v
			return nil
		}
		return fmt.Errorf("unexpected value of type %T", v)
	case *Bytes65:
		if v, ok := v.(Bytes65); ok {
			*interf = v
			return nil
		}
		return fmt.Errorf("unexpected value of type %T", v)
	case *Struct:
		if v, ok := v.(Struct); ok {
			*interf = v
			return nil
		}
		return fmt.Errorf("unexpected value of type %T", v)
	case *List:
		if v, ok := v.(List); ok {
			*interf = v
			return nil
		}
		return fmt.Errorf("unexpected value of type %T", v)
	case *Typed:
		if v, ok := v.(Typed); ok {
			*interf = v
			return nil
		}
		if v, ok := v.(Struct); ok {
			*interf = Typed(v)
			return nil
		}
		return fmt.Errorf("unexpected value of type %T", v)
	case *Value:
		if v, ok := v.(Value); ok {
			*interf = v
			return nil
		}
		return fmt.Errorf("unexpected value of type %T", v)
	}

	// Otherwise, reflect on the kind/type of the interface, and attempt to
	// convert the value-to-be-decoded into the interface.
	valueOf := reflect.ValueOf(interf)
	if valueOf.Kind() != reflect.Ptr {
		return fmt.Errorf("expected %v, got %v", reflect.Ptr, valueOf.Kind())
	}
	elem := valueOf.Elem()
	switch elem.Kind() {
	case reflect.Bool:
		if v, ok := v.(Bool); ok {
			elem.SetBool(bool(v))
			return nil
		}
		return fmt.Errorf("unexpected value of type %T", v)
	case reflect.Uint8:
		if v, ok := v.(U8); ok {
			elem.SetUint(uint64(v.Uint8()))
			return nil
		}
		return fmt.Errorf("unexpected value of type %T", v)
	case reflect.Uint16:
		if v, ok := v.(U16); ok {
			elem.SetUint(uint64(v.Uint16()))
			return nil
		}
		return fmt.Errorf("unexpected value of type %T", v)
	case reflect.Uint32:
		if v, ok := v.(U32); ok {
			elem.SetUint(uint64(v.Uint32()))
			return nil
		}
		return fmt.Errorf("unexpected value of type %T", v)
	case reflect.Uint64:
		if v, ok := v.(U64); ok {
			elem.SetUint(v.Uint64())
			return nil
		}
		return fmt.Errorf("unexpected value of type %T", v)
	case reflect.String:
		if v, ok := v.(String); ok {
			elem.SetString(string(v))
			return nil
		}
		return fmt.Errorf("unexpected value of type %T", v)
	case reflect.Slice:
		typeOf := elem.Type()
		if typeOf.Elem().Kind() == reflect.Uint8 {
			if v, ok := v.(Bytes); ok {
				elem.SetBytes([]byte(v))
				return nil
			}
			return fmt.Errorf("unexpected value of type %T", v)
		}
		elem.Set(reflect.MakeSlice(typeOf, len(v.(List).Elems), len(v.(List).Elems)))
		for i := 0; i < len(v.(List).Elems); i++ {
			if err := Decode(elem.Index(i).Addr().Interface(), v.(List).Elems[i]); err != nil {
				return fmt.Errorf("decoding list item: %v", err)
			}
		}
		return nil
	case reflect.Array:
		typeOf := elem.Type()
		if typeOf.Elem().Kind() == reflect.Uint8 {
			if typeOf.Len() == 32 {
				if v, ok := v.(Bytes32); ok {
					elem.Set(reflect.ValueOf(v).Convert(typeOf))
					return nil
				}
				return fmt.Errorf("unexpected value of type %T", v)
			}
			if typeOf.Len() == 65 {
				if v, ok := v.(Bytes65); ok {
					elem.Set(reflect.ValueOf(v).Convert(typeOf))
					return nil
				}
				return fmt.Errorf("unexpected value of type %T", v)
			}
		}
		return fmt.Errorf("non-exhaustive pattern: type %T", v)
	case reflect.Struct:
		var structOrTyped Struct
		if s, ok := v.(Struct); ok {
			structOrTyped = s
		} else if t, ok := v.(Typed); ok {
			structOrTyped = Struct(t)
		} else {
			return fmt.Errorf("non-exhaustive pattern: type %T", v)
		}

		typeOf := elem.Type()
		n := typeOf.NumField()
		for i := 0; i < n; i++ {
			f := typeOf.Field(i)
			tags := strings.Split(f.Tag.Get("json"), ",")
			name := f.Name
			for _, tag := range tags {
				if tag == "-" {
					name = ""
					break
				}
				if tag != "omitempty" {
					name = tag
					break
				}
			}
			if name != "" {
				// If the struct value is nil, do not decode it.
				if structOrTyped.Get(name) == nil {
					continue
				}
				if err := Decode(elem.Field(i).Addr().Interface(), structOrTyped.Get(name)); err != nil {
					return fmt.Errorf("decoding \"%v\": %v", f.Name, err)
				}
			}
		}
		return nil
	default:
		return fmt.Errorf("non-exhaustive pattern: type %T", v)
	}
}

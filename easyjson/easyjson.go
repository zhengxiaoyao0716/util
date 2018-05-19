package easyjson

import (
	"encoding/json"
	"unicode"
)

// JSON .
type JSON interface {
	String() string
	JSONString() string
}

// Object .
type Object map[string]interface{}

// Array .
type Array []interface{}

// String .
type String string

// Boolean .
type Boolean bool

// Number .
type Number float64

type container interface{ IsEmpty() bool }

// IsEmpty .
func (json *Object) IsEmpty() bool { return len(*json) == 0 }

// IsEmpty .
func (json *Array) IsEmpty() bool { return len(*json) == 0 }

// ValueAt .
func (json *Object) ValueAt(name string, defaults ...interface{}) (interface{}, error) {
	value, ok := (*json)[name]
	if !ok {
		if len(defaults) > 0 {
			for _, v := range defaults {
				if v != nil {
					return v, nil
				}
			}
		}
		return nil, &ValueNotFoundError{refContainer{json}, name}
	}
	return value, nil
}

// MustValueAt .
func (json *Object) MustValueAt(name string, defaults ...interface{}) interface{} {
	r, err := json.ValueAt(name, defaults...)
	if err != nil {
		panic(err)
	}
	return r
}

// ValueAt .
func (json *Array) ValueAt(index int, defaults ...interface{}) (interface{}, error) {
	if index >= len(*json) {
		if len(defaults) > 0 {
			for _, v := range defaults {
				if v != nil {
					return v, nil
				}
			}
		}
		return nil, &ValueNotFoundError{refContainer{json}, index}
	}
	return (*json)[index], nil
}

// MustValueAt .
func (json *Array) MustValueAt(index int, defaults ...interface{}) interface{} {
	r, err := json.ValueAt(index, defaults...)
	if err != nil {
		panic(err)
	}
	return r
}

//go:generate go run gen.go

// Parse .
func Parse(s string) (JSON, error) {
	var j JSON
	var err error
	for _, r := range s {
		if unicode.IsSpace(r) {
			continue
		}
		switch r {
		case '{':
			data := Object{}
			err = json.Unmarshal([]byte(s), &data)
			j = data
		case '[':
			data := Array{}
			err = json.Unmarshal([]byte(s), &data)
			j = data
		case '"':
			data := String("")
			err = json.Unmarshal([]byte(s), &data)
			j = data
		case 't':
			fallthrough
		case 'f':
			data := Boolean(false)
			err = json.Unmarshal([]byte(s), &data)
			j = data
		default:
			data := Number(0)
			err = json.Unmarshal([]byte(s), &data)
			j = data
		}
		break
	}
	if err != nil {
		return nil, err
	}
	return j, nil
}

// MustParse .
func MustParse(s string) JSON {
	json, err := Parse(s)
	if err != nil {
		panic(err)
	}
	return json
}

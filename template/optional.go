package template

import (
	"encoding/json"
	"fmt"
	"time"
)

var (
	_ = time.Time{}
)

// template type Optional(T)

type T string

// Optional wraps a value that may or may not be nil.
// If a value is present, it may be unwrapped to expose the underlying value.
type Optional optional

type optional struct {
	V   T
	Set bool
}

// Of wraps the value in an optional.
func Of(value T) Optional {
	return Optional{Val: value, Set: true}
}

func OfOptionalPtr(ptr *T) Optional {
	if ptr == nil {
		return Empty()
	} else {
		return Of(*ptr)
	}
}

// Empty returns an empty optional.
func Empty() Optional {
	return Optional{}
}

// Get returns the value wrapped by this optional, and an ok signal for whether a value was wrapped.
func (o Optional) Get() (value T, ok bool) {
	o.If(func(v T) {
		value = v
		ok = true
	})
	return
}

func (o Optional) Value() (value interface{}, ok bool) {
	o.If(func(v T) {
		value = v
		ok = true
	})
	return
}

func (o Optional) Ptr() (value *T) {
	if o.Set {
		return &o.Val
	}
	return nil
}

// IsPresent returns true if there is a value wrapped by this optional.
func (o Optional) IsPresent() bool {
	return o.Set
}

// If calls the function if there is a value wrapped by this optional.
func (o Optional) If(f func(value T)) {
	if o.IsPresent() {
		f(o.V)
	}
}

func (o Optional) ElseFunc(f func() T) (value T) {
	if o.IsPresent() {
		o.If(func(v T) { value = v })
		return
	} else {
		return f()
	}
}

// Else returns the value wrapped by this optional, or the value passed in if
// there is no value wrapped by this optional.
func (o Optional) Else(elseValue T) (value T) {
	return o.ElseFunc(func() T { return elseValue })
}

// ElseZero returns the value wrapped by this optional, or the zero value of
// the type wrapped if there is no value wrapped by this optional.
func (o Optional) ElseZero() (value T) {
	var zero T
	return o.Else(zero)
}

// String returns the string representation of the wrapped value, or the string
// representation of the zero value of the type wrapped if there is no value
// wrapped by this optional.
func (o Optional) String() string {
	return fmt.Sprintf("%+v", o.ElseZero())
}

// MarshalJSON marshals the value being wrapped to JSON. If there is no vale
// being wrapped, the zero value of its type is marshaled.
func (o Optional) MarshalJSON() (data []byte, err error) {
	if o.Set {
		return json.Marshal(o.V)
	}
	return []byte("null"), nil
}

// UnmarshalJSON unmarshals the JSON into a value wrapped by this optional.
func (o *Optional) UnmarshalJSON(data []byte) error {
	o.Set = true
	if string(data) == "null" {
		return nil
	}
	var v T
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	*o = Of(v)
	return nil
}

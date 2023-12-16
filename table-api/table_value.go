package tableapi

import (
	"fmt"
	"reflect"
)

type value interface{}

// TableValue is the reflection interface to a table value.
type TableValue struct {
	value        value
	displayValue value
	link         string
}

func (tV *TableValue) preferredValue() value {
	if tV.displayValue != nil {
		return tV.displayValue
	}
	return tV.value

// Deprecated: deprecated as of {version} please utilize `ToInt64`
//
// ToInt64 returns tV's underlying value, as an int64.
func (tV *TableValue) ToInt64() (int64, error) {

	return tV.Int()
}

// Int returns tV's underlying value, as an int64.
func (tv *TableValue) Int() (int64, error) {
	switch v := tv.preferredValue().(type) {
	case int:
		return int64(v), nil
	case int8:
		return int64(v), nil
	case int16:
		return int64(v), nil
	case int32:
		return int64(v), nil
	case int64:
		return v, nil
	default:
		return 0, fmt.Errorf("unable to convert %T to int64", tv.value)
	}
}

// Deprecated: deprecated as of {version} please utilize `ToFloat64`
//
// ToFloat64 returns tV's underlying value, as a float64.
func (tV *TableValue) ToFloat64() (float64, error) {
	return tV.Float()
}

// Float returns tV's underlying value, as a float64.
func (tV *TableValue) Float() (float64, error) {
	switch v := tV.preferredValue().(type) {
	case float32:
		return float64(v), nil
	case float64:
		return v, nil
	default:
		return 0, fmt.Errorf("unable to convert %T to float64", tV.value)
	}
}

// Deprecated: deprecated as of {version} please utilize `String`
//
// ToString returns tV's underlying value, as a string.
func (tV *TableValue) ToString() (string, error) {
	return tV.String()
}

// String returns tV's underlying value, as a string.
func (tV *TableValue) String() (string, error) {
	return convertType[string](tV.preferredValue())
}

// Deprecated: deprecated as of {version} please utilize `Bool`
//
// ToBool returns tV's underlying value, as a bool.
func (tV *TableValue) ToBool() (bool, error) {
	return tV.Bool()
}

// Bool returns tV's underlying value, as a bool.
func (tV *TableValue) Bool() (bool, error) {
	return convertType[bool](tV.preferredValue())
}

// Deprecated: deprecated as of {version} please utilize `Type`
//
// GetType returns tV's underlying value type.
func (tV *TableValue) GetType() reflect.Type {
	return tV.Type()
}

// Type returns tV's underlying value type.
func (tV *TableValue) Type() reflect.Type {
	return reflect.TypeOf(tV.preferredValue())
}

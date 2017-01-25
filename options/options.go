/*
 little helper to get "options" from a map[string]interface{}
 but more in a typed fashion
*/

package options

import (
	"fmt"
	"time"
)

// Options type
type Options map[string]interface{}

// New  Options object
func New() Options {
	return Options(make(map[string]interface{}))
}

func (o *Options) get(name string) (interface{}, bool) {
	c := map[string]interface{}(*o)
	gots, ok := c[name]
	return gots, ok
}

// Set
func (o *Options) Set(name string, val interface{}) {
	c := map[string]interface{}(*o)
	c[name] = val
}

// String get a string or the default
func (o *Options) String(name, def string) string {
	got, ok := o.get(name)
	if ok {
		return got.(string)
	}
	return def
}

// StringRequired get a string or return an error if option is not there
func (o *Options) StringRequired(name string) (string, error) {
	got, ok := o.get(name)
	if ok {
		return got.(string), nil
	}
	return "", fmt.Errorf("%s is required", name)
}

// Object get an object (it is up the called to type check the object)
func (o *Options) Object(name, def string) interface{} {
	got, ok := o.get(name)
	if ok {
		return got
	}
	return def
}

// ObjectRequired get an object (it is up the called to type check the object) or return there will be an
// error if not present
func (o *Options) ObjectRequired(name string) (interface{}, error) {
	got, ok := o.get(name)
	if ok {
		return got, nil
	}
	return nil, fmt.Errorf("%s is required", name)
}

// Int64 get an int64 or the default
// this will cast any other int/uint type into an int64
// and will panic if things cannot be converted to an int64
func (o *Options) Int64(name string, def int64) int64 {
	got, ok := o.get(name)
	if ok {
		switch got.(type) {
		case int64:
			return got.(int64)
		case uint8:
			return int64(got.(uint8))
		case uint16:
			return int64(got.(uint16))
		case uint32:
			return int64(got.(uint32))
		case uint64:
			return int64(got.(uint64))
		case int8:
			return int64(got.(int8))
		case int16:
			return int64(got.(int16))
		case int32:
			return int64(got.(int32))
		case int:
			return int64(got.(int))
		case float64:
			return int64(got.(float64))
		case float32:
			return int64(got.(float32))
		default:
			panic(fmt.Errorf("Cannot convert int type"))
		}
	}
	return def
}

// Int64Required get an int64 or error if it does not exist
// this will cast any other int/uint type into an int64
// and will panic if things cannot be converted to an int64
func (o *Options) Int64Required(name string) (int64, error) {
	got, ok := o.get(name)
	if ok {
		switch got.(type) {
		case int64:
			return got.(int64), nil
		case uint8:
			return int64(got.(uint8)), nil
		case uint16:
			return int64(got.(uint16)), nil
		case uint32:
			return int64(got.(uint32)), nil
		case uint64:
			return int64(got.(uint64)), nil
		case int8:
			return int64(got.(int8)), nil
		case int16:
			return int64(got.(int16)), nil
		case int32:
			return int64(got.(int32)), nil
		case int:
			return int64(got.(int)), nil
		case float64:
			return int64(got.(float64)), nil
		case float32:
			return int64(got.(float32)), nil
		default:
			return 0, fmt.Errorf("%s is not an int (it's a `%T`)", name, got)
		}
	}
	return 0, fmt.Errorf("%s is required", name)
}

// Float64 get an float64 or the default
// this will cast any other numeric types into an float64
// and will panic if things cannot be converted to an float64
func (o *Options) Float64(name string, def float64) float64 {
	got, ok := o.get(name)
	if ok {
		switch got.(type) {
		case int64:
			return float64(got.(int64))
		case uint8:
			return float64(got.(uint8))
		case uint16:
			return float64(got.(uint16))
		case uint32:
			return float64(got.(uint32))
		case uint64:
			return float64(got.(uint64))
		case int8:
			return float64(got.(int8))
		case int16:
			return float64(got.(int16))
		case int32:
			return float64(got.(int32))
		case int:
			return float64(got.(int))
		case float32:
			return float64(got.(float32))
		case float64:
			return got.(float64)
		default:
			panic(fmt.Errorf("%s is not a float", name))
		}
	}
	return def
}

// Float64 get an float64 or an error
// this will cast any other numeric types into a float64
// and will panic if things cannot be converted to an float64
func (o *Options) Float64Required(name string) (float64, error) {
	got, ok := o.get(name)
	if ok {
		switch got.(type) {
		case int64:
			return float64(got.(int64)), nil
		case uint8:
			return float64(got.(uint8)), nil
		case uint16:
			return float64(got.(uint16)), nil
		case uint32:
			return float64(got.(uint32)), nil
		case uint64:
			return float64(got.(uint64)), nil
		case int8:
			return float64(got.(int8)), nil
		case int16:
			return float64(got.(int16)), nil
		case int32:
			return float64(got.(int32)), nil
		case int:
			return float64(got.(int)), nil
		case float32:
			return float64(got.(float32)), nil
		case float64:
			return got.(float64), nil
		default:
			return 0, fmt.Errorf("%s is not a float", name)

		}
	}
	return 0, fmt.Errorf("%s is required", name)
}

// Bool get a bool or the default
func (o *Options) Bool(name string, def bool) bool {
	got, ok := o.get(name)
	if ok {
		return got.(bool)
	}
	return def
}

// BoolRequired get a bool or an error
func (o *Options) BoolRequired(name string) (bool, error) {
	got, ok := o.get(name)
	if ok {
		return got.(bool), nil
	}
	return false, fmt.Errorf("%s is required", name)
}

// Duration get a duration or a default.
// the entry can be a string or another time.Duration object and will panic otherwise
// If the option is a string, it will attempt to parse the duration
// If this fails it will panic.
func (o *Options) Duration(name string, def time.Duration) time.Duration {
	got, ok := o.get(name)
	if ok {
		switch got.(type) {
		case string:
			rdur, err := time.ParseDuration(got.(string))
			if err != nil {
				panic(err)
			}
			return rdur
		case time.Duration:
			return got.(time.Duration)
		default:
			panic(fmt.Sprintf("%s is not a string or duration (it's a `%T`)", name, got))
		}

	}
	return def
}

// DurationRequired get a duration or an error if not found.
// the entry can be a string or another time.Duration object and will panic otherwise
// If the option is a string, it will attempt to parse the duration
// If this fails it will panic.
func (o *Options) DurationRequired(name string) (time.Duration, error) {
	got, ok := o.get(name)
	if ok {
		rdur, err := time.ParseDuration(got.(string))
		if err != nil {
			return time.Duration(0), err
		}
		return rdur, nil
	}
	return time.Duration(0), fmt.Errorf("%s is required", name)
}

// ToString a pretty print function
func (o *Options) ToString() string {
	out := "Options("
	for k, v := range *o {
		out += fmt.Sprintf("%s=%v, ", k, v)
	}
	out += ")"
	return out
}

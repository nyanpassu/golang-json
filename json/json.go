package json

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// Any .
type Any interface {
	// true if node is null or undefined
	Null() bool
	// return string value if is string
	// return "" if null
	// return error otherwise
	StringValue() (string, error)
	// will return the json format of node
	AsText() (string, error)
	// if node is int, then return int value
	// nullnode will return 0
	// otherwise return error
	IntValue() (int64, error)
	// int node will return int value
	// if is float, will convert to int by int64(value)
	// if node is string, will convert to int64 by strconv.formatInt
	// otherwise return error
	AsInt() (int64, error)
	// return float value if float
	// return 0 if null | undefined
	// return error otherwise
	FloatValue() (float64, error)
	// if node is string, will try convert to float
	// return float64(value) when int
	// return error otherwise
	AsFloat() (float64, error)
	// return converted value by strconv.formatBool if is string
	// return true if non zero
	// return false if is zero
	// return false if null
	// return error otherwise
	BoolValue() (bool, error)
	// return bool value if node is bool
	// return false if null | undefined
	// otherwise return error
	AsBool() (bool, error)
	// return node array if is array
	// return nil if null | undefined
	// otherwise return error
	ArrayValue() ([]Any, error)
	// return node map if node is object
	// return nil if null | undefined
	// otherwise return error
	ObjectValue() (map[string]Any, error)
	// return raw value
	Value() interface{}
}

// Unmarshal .
func Unmarshal(data []byte) (Any, error) {
	var src interface{}
	if err := json.Unmarshal(data, &src); err != nil {
		return nil, err
	}
	return newNode(src)
}

// Marshal .
func Marshal(obj Any) ([]byte, error) {
	return json.Marshal(obj.Value())
}

func newNode(src interface{}) (Any, error) {
	if src == nil {
		return newNullNode(), nil
	}
	if str, ok := src.(string); ok {
		return newStringNode(str), nil
	}
	if f, ok := src.(float64); ok {
		return newFloatNode(f), nil
	}
	if array, ok := src.([]interface{}); ok {
		return newArrayNode(array)
	}
	if objectMapping, ok := src.(map[string]interface{}); ok {
		return newObjectNode(objectMapping)
	}
	if boolean, ok := src.(bool); ok {
		return newBoolNode(boolean), nil
	}
	return newNode2(src)
}

func newNode2(src interface{}) (Any, error) {
	if f, ok := src.(float32); ok {
		return newFloatNode(float64(f)), nil
	}
	if i, ok := src.(int); ok {
		return newIntNode(int64(i)), nil
	}
	if i, ok := src.(int8); ok {
		return newIntNode(int64(i)), nil
	}
	if i, ok := src.(int16); ok {
		return newIntNode(int64(i)), nil
	}
	if i, ok := src.(int32); ok {
		return newIntNode(int64(i)), nil
	}
	if i, ok := src.(int64); ok {
		return newIntNode(int64(i)), nil
	}
	if i, ok := src.(uint); ok {
		return newIntNode(int64(i)), nil
	}
	if i, ok := src.(uint8); ok {
		return newIntNode(int64(i)), nil
	}
	if i, ok := src.(uint16); ok {
		return newIntNode(int64(i)), nil
	}
	if i, ok := src.(uint32); ok {
		return newIntNode(int64(i)), nil
	}
	if i, ok := src.(uint64); ok {
		return newIntNode(int64(i)), nil
	}
	if b, ok := src.(byte); ok {
		return newIntNode(int64(b)), nil
	}
	if r, ok := src.(rune); ok {
		return newStringNode(string(r)), nil
	}
	return nil, fmt.Errorf("Unsupport type conversion: %v", reflect.TypeOf(src))
}

func asText(any interface{}) (string, error) {
	bytes, err := json.Marshal(any)
	return string(bytes), err
}

func errorf(format string, obj ...interface{}) error {
	return fmt.Errorf(format, obj...)
}

func notString(any interface{}) (string, error) {
	return "", errorf("not a string: %v", any)
}

func notInt(any interface{}) (int64, error) {
	return 0, errorf("not an int: %v", any)
}

func notFloat(any interface{}) (float64, error) {
	return 0, errorf("not a float: %v", any)
}

func notBool(any interface{}) (bool, error) {
	return false, errorf("not a bool: %v", any)
}

func notArray(any interface{}) ([]Any, error) {
	return nil, errorf("not an array: %v", any)
}

func notObject(any interface{}) (map[string]Any, error) {
	return nil, errorf("not an object: %v", any)
}

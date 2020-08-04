package json

import (
	"strconv"
)

type numberNode struct {
	floatValue float64
	intValue   int64
	isInt      bool
}

// Int .
func Int(i int64) Any {
	return numberNode{
		intValue: i,
		isInt:    true,
	}
}

// Float .
func Float(f float64) Any {
	return numberNode{
		floatValue: f,
		isInt:      false,
	}
}

func newFloatNode(value float64) Any {
	intValue := int64(value)
	if value == float64(intValue) {
		return numberNode{
			floatValue: value,
			intValue:   intValue,
			isInt:      true,
		}
	}
	return numberNode{
		floatValue: value,
		isInt:      false,
	}
}

func newIntNode(value int64) Any {
	return numberNode{
		floatValue: float64(value),
		intValue:   value,
		isInt:      true,
	}
}

// always return false
func (node numberNode) Null() bool {
	return false
}

func (node numberNode) StringValue() (string, error) {
	return notString(node.floatValue)
}

func (node numberNode) AsText() (string, error) {
	if node.isInt {
		return strconv.FormatInt(node.intValue, 64), nil
	} else {
		return strconv.FormatFloat(node.floatValue, 'b', -1, 64), nil
	}
}

func (node numberNode) IntValue() (int64, error) {
	if node.isInt {
		return node.intValue, nil
	}
	return notInt(node.floatValue)
}

func (node numberNode) AsInt() (int64, error) {
	if node.isInt {
		return node.intValue, nil
	}
	return int64(node.floatValue), nil
}

func (node numberNode) FloatValue() (float64, error) {
	return node.floatValue, nil
}

func (node numberNode) AsFloat() (float64, error) {
	return node.floatValue, nil
}

func (node numberNode) BoolValue() (bool, error) {
	return notBool(node.floatValue)
}

func (node numberNode) AsBool() (bool, error) {
	return node.floatValue != 0, nil
}

func (node numberNode) ArrayValue() ([]Any, error) {
	return notArray(node.floatValue)
}

func (node numberNode) ObjectValue() (map[string]Any, error) {
	return notObject(node.floatValue)
}

func (node numberNode) Value() interface{} {
	if node.isInt {
		return node.intValue
	}
	return node.floatValue
}

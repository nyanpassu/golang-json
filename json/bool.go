package json

import (
	"strconv"
)

type boolNode struct {
	value bool
}

// Bool .
func Bool(value bool) Any {
	return boolNode{value}
}

func newBoolNode(value bool) Any {
	return boolNode{value}
}

// always return false
func (node boolNode) Null() bool {
	return false
}

func (node boolNode) StringValue() (string, error) {
	return notString(node.value)
}

func (node boolNode) AsText() (string, error) {
	return strconv.FormatBool(node.value), nil
}

func (node boolNode) IntValue() (int64, error) {
	return notInt(node.value)
}

func (node boolNode) AsInt() (int64, error) {
	return notInt(node.value)
}

func (node boolNode) FloatValue() (float64, error) {
	return notFloat(node.value)
}

func (node boolNode) AsFloat() (float64, error) {
	return notFloat(node.value)
}

func (node boolNode) BoolValue() (bool, error) {
	return node.value, nil
}

func (node boolNode) AsBool() (bool, error) {
	return node.value, nil
}

func (node boolNode) ArrayValue() ([]Any, error) {
	return notArray(node.value)
}

func (node boolNode) ObjectValue() (map[string]Any, error) {
	return notObject(node.value)
}

func (node boolNode) Value() interface{} {
	return node.value
}

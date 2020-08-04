package json

import (
	"strconv"
)

type stringNode struct {
	value string
}

// String .
func String(str string) Any {
	return stringNode{value: str}
}

func newStringNode(value string) Any {
	return stringNode{value}
}

// always return false
func (node stringNode) Null() bool {
	return false
}

func (node stringNode) StringValue() (string, error) {
	return node.value, nil
}

func (node stringNode) AsText() (string, error) {
	return node.value, nil
}

func (node stringNode) IntValue() (int64, error) {
	return notInt(node.value)
}

func (node stringNode) AsInt() (int64, error) {
	return strconv.ParseInt(node.value, 10, 64)
}

func (node stringNode) FloatValue() (float64, error) {
	return notFloat(node.value)
}

func (node stringNode) AsFloat() (float64, error) {
	return strconv.ParseFloat(node.value, 64)
}

func (node stringNode) BoolValue() (bool, error) {
	return notBool(node.value)
}

func (node stringNode) AsBool() (bool, error) {
	return strconv.ParseBool(node.value)
}

func (node stringNode) ArrayValue() ([]Any, error) {
	return notArray(node.value)
}

func (node stringNode) ObjectValue() (map[string]Any, error) {
	return notObject(node.value)
}

func (node stringNode) Value() interface{} {
	return node.value
}

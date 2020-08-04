package json

type nullNode struct{}

var null = nullNode{}

// Null .
func Null() Any {
	return null
}

func newNullNode() Any {
	return nullNode{}
}

// always return false
func (node nullNode) Null() bool {
	return true
}

func (node nullNode) StringValue() (string, error) {
	return "", nil
}

func (node nullNode) AsText() (string, error) {
	return "", nil
}

func (node nullNode) IntValue() (int64, error) {
	return 0, nil
}

func (node nullNode) AsInt() (int64, error) {
	return 0, nil
}

func (node nullNode) FloatValue() (float64, error) {
	return 0, nil
}

func (node nullNode) AsFloat() (float64, error) {
	return 0, nil
}

func (node nullNode) BoolValue() (bool, error) {
	return false, nil
}

func (node nullNode) AsBool() (bool, error) {
	return false, nil
}

func (node nullNode) ArrayValue() ([]Any, error) {
	return nil, nil
}

func (node nullNode) ObjectValue() (map[string]Any, error) {
	return nil, nil
}

func (node nullNode) Value() interface{} {
	return nil
}

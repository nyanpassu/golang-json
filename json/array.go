package json

type arrayNode struct {
	array []Any
}

func newArrayNode(src []interface{}) (Any, error) {
	size := len(src)
	array := make([]Any, size)
	for i := 0; i < size; i++ {
		var err error
		if array[i], err = newNode(src[i]); err != nil {
			return nil, err
		}
	}
	return arrayNode{array}, nil
}

func (node arrayNode) Null() bool {
	return false
}

func (node arrayNode) StringValue() (string, error) {
	return notString(node.array)
}

func (node arrayNode) AsText() (string, error) {
	return asText(node.array)
}

func (node arrayNode) IntValue() (int64, error) {
	return notInt(node.array)
}

func (node arrayNode) AsInt() (int64, error) {
	return notInt(node.array)
}

func (node arrayNode) FloatValue() (float64, error) {
	return notFloat(node.array)
}

func (node arrayNode) AsFloat() (float64, error) {
	return notFloat(node.array)
}

func (node arrayNode) AsBool() (bool, error) {
	return notBool(node.array)
}

func (node arrayNode) BoolValue() (bool, error) {
	return notBool(node.array)
}

func (node arrayNode) ArrayValue() ([]Any, error) {
	return node.array, nil
}

func (node arrayNode) ObjectValue() (map[string]Any, error) {
	return notObject(node.array)
}

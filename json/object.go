package json

type objectNode struct {
	objectMapping map[string]Any
}

// Object .
func Object(mapping map[string]Any) Any {
	return objectNode{objectMapping: mapping}
}

func newObjectNode(src map[string]interface{}) (Any, error) {
	mapping := make(map[string]Any)
	for key, value := range src {
		var err error
		if mapping[key], err = newNode(value); err != nil {
			return nil, err
		}
	}
	return objectNode{mapping}, nil
}

// always return false
func (node objectNode) Null() bool {
	return false
}

func (node objectNode) StringValue() (string, error) {
	return notString(node.objectMapping)
}

func (node objectNode) AsText() (string, error) {
	return asText(node.objectMapping)
}

func (node objectNode) IntValue() (int64, error) {
	return notInt(node.objectMapping)
}

func (node objectNode) AsInt() (int64, error) {
	return notInt(node.objectMapping)
}

func (node objectNode) FloatValue() (float64, error) {
	return notFloat(node.objectMapping)
}

func (node objectNode) AsFloat() (float64, error) {
	return notFloat(node.objectMapping)
}

func (node objectNode) BoolValue() (bool, error) {
	return notBool(node.objectMapping)
}

func (node objectNode) AsBool() (bool, error) {
	return notBool(node.objectMapping)
}

func (node objectNode) ArrayValue() ([]Any, error) {
	return notArray(node.objectMapping)
}

func (node objectNode) ObjectValue() (map[string]Any, error) {
	return node.objectMapping, nil
}

func (node objectNode) Value() interface{} {
	mapping := make(map[string]interface{})
	for key, value := range node.objectMapping {
		mapping[key] = value.Value()
	}
	return mapping
}

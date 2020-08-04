package json

type lazyNode struct {
	lazy func() Any
	node Any
}

func newLazyNode(lazy func() Any) Any {
	return lazyNode{lazy: lazy}
}

func (node lazyNode) getNode() Any {
	if node.node != nil {
		return node.node
	}
	node.node = node.lazy()
	node.lazy = nil
	return node.node
}

func (node lazyNode) Null() bool {
	return node.getNode().Null()
}

func (node lazyNode) StringValue() (string, error) {
	return node.getNode().StringValue()
}

func (node lazyNode) AsText() (string, error) {
	return node.getNode().AsText()
}

func (node lazyNode) IntValue() (int64, error) {
	return node.getNode().IntValue()
}

func (node lazyNode) AsInt() (int64, error) {
	return node.getNode().AsInt()
}

func (node lazyNode) FloatValue() (float64, error) {
	return node.getNode().FloatValue()
}

func (node lazyNode) AsFloat() (float64, error) {
	return node.getNode().AsFloat()
}

func (node lazyNode) BoolValue() (bool, error) {
	return node.getNode().BoolValue()
}

func (node lazyNode) AsBool() (bool, error) {
	return node.getNode().AsBool()
}

func (node lazyNode) ArrayValue() ([]Any, error) {
	return node.getNode().ArrayValue()
}

func (node lazyNode) ObjectValue() (map[string]Any, error) {
	return node.getNode().ObjectValue()
}

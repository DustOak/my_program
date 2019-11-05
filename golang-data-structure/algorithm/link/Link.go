package link

type Api interface {
	AddHead(value interface{}) *Node
	DeleteFirst() *Node
	Append(value interface{})
	HasNext() bool
	NextNode() *Node
}

type Node struct {
	Value interface{}
	Next  *Node
}

func (n *Node) HasNext() bool {
	return n != nil
}

func (n *Node) NextNode() *Node {
	return n.Next
}

func (n *Node) AddHead(value interface{}) *Node {
	temp := NewLinkedList(value)
	temp.Next = n
	return temp

}

func (n *Node) DeleteFirst() *Node {
	return n.Next
}

func (n *Node) Append(value interface{}) {
	for current := n; ; current = current.Next {
		if current.Next == nil {
			current.Next = NewLinkedList(value)
			break
		}
	}
}

func NewLinkedList(value interface{}) *Node {
	return &Node{
		Value: value,
		Next:  nil,
	}
}

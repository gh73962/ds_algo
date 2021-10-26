package list

// reverse single list
// list is a ring
// merge two sorted list
// delete node which last n
// get mid node of list

type List struct {
	root   *Node
	length int
}

func (l *List) Insert(v interface{}) {

}

type Node struct {
	next  *Node
	value interface{}
}

func (n *Node) Value() interface{} {
	return n.value
}

func (n *Node) GetNext() *Node {
	return n.next
}

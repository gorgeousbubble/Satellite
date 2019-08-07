package ds

type Object interface {}

type Node struct {
	data Object
	next *Node
}

type List struct {
	head *Node
	tail *Node
	size uint64
}

func (list *List) Init() {
	list.head = nil
	list.tail = nil
	list.size = 0
}



package ds

type Object interface {}

type Node struct {
	data Object
	next *Node
}

type List struct {
	head *Node
	tail *Node
	size int
}

func (list *List) Init() {
	list.head = nil
	list.tail = nil
	list.size = 0
}

func (list *List) Append(node *Node) bool {
	if node == nil {
		return false
	}
	(*node).next = nil
	if (*list).size == 0 {
		(*list).head = node
	} else {
		tail := (*list).tail
		(*tail).next = node
	}
	(*list).tail = node
	(*list).size += 1
	return true
}

func (list *List) Insert(i int, node *Node) bool {
	if node == nil || i > (*list).size || (*list).size == 0 {
		return false
	}
	if i == 0 {
		(*node).next = (*list).head
		(*list).head = node
	} else {
		pre := (*list).head
		for j := 1; j < i; i++ {
			pre = (*pre).next
		}
		(*node).next = (*pre).next
		(*pre).next = pre
	}
	(*list).size++
	return true
}

func (list *List) Remove(i int, node *Node) bool {
	if i >= (*list).size {
		return false
	}
	if i == 0 {
		node = (*list).head
		(*list).head = (*node).next
		if (*list).size == 1 {
			(*list).tail = nil
		}
	} else {
		pre := (*list).head
		for j := 1; j < i; j++ {
			pre = (*pre).next
		}
		node = (*pre).next
		(*pre).next = (*node).next
		if i == ((*list).size - 1) {
			(*list).tail = pre
		}
	}
	(*list).size--
	return true
}

func (list *List) Get(i int) *Node {
	if i >= (*list).size {
		return nil
	}
	item := (*list).head
	for j := 0; j < i ; j++ {
		item = (*item).next
	}
	return item
}

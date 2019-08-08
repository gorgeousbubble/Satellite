package ds

type DNode struct {
	data Object
	prev *DNode
	next *DNode
}

type DList struct {
	head *DNode
	tail *DNode
	size int
}

func (dlist *DList) Init() {
	(*dlist).size = 0
	(*dlist).head = nil
	(*dlist).tail = nil
}

func (dlist *DList) Append(data Object) {
	newNode := new(DNode)
	(*newNode).data =  data

	if (*dlist).GetSize() == 0 {
		(*dlist).head = newNode
		(*dlist).tail = newNode
		(*newNode).prev = nil
		(*newNode).next = nil
	} else {
		(*newNode).prev = (*dlist).tail
		(*newNode).next = nil
		(*((*dlist).tail)).next = newNode
		(*dlist).tail = newNode
	}
	(*dlist).size++
}

func (dlist *DList) InsertNext(node *DNode, data Object) bool {
	if node == nil {
		return false
	}
	if dlist.isTail(node) {
		dlist.Append(data)
	} else {
		newNode := new(DNode)
		(*newNode).data =  data
		(*newNode).prev = node
		(*newNode).next = (*node).next

		(*node).next = newNode
		(*((*newNode).next)).prev = newNode
		(*dlist).size++
	}
	return true
}

func (dlist *DList) InsertPrev(node *DNode, data Object) bool {
	if node == nil {
		return false
	}
	if dlist.isHead(node) {
		newNode := new(DNode)
		(*newNode).data = data
		(*newNode).next = dlist.GetHead()
		(*newNode).prev = nil

		(*(dlist.head)).prev = newNode
		dlist.head = newNode
		dlist.size++
		return true
	} else {
		prev := (*node).prev
		return dlist.InsertNext(prev, data)
	}
}

func (dlist *DList) Remove(node *DNode) Object {
	if node == nil {
		return false
	}

	prev := (*node).prev
	next := (*node).next

	if dlist.isHead(node) {
		dlist.head = next
	} else {
		(*prev).next = next
	}

	if dlist.isTail(node) {
		dlist.tail = prev
	} else {
		(*next).prev = prev
	}

	dlist.size--
	return (*node).GetData()
}

func (dlist *DList) Search(data Object) *DNode {
	if dlist.GetSize() == 0 {
		return nil
	}

	node := dlist.GetHead()
	for ; node != nil; node = node.GetNext() {
		if node.GetData() == data {
			break
		}
	}

	return node
}

func (dlist *DList) GetSize() int {
	return (*dlist).size
}

func (dlist *DList) GetHead() *DNode {
	return (*dlist).head
}

func (dlist *DList) GetTail() *DNode {
	return (*dlist).tail
}

func (dlist *DList) isHead(node *DNode) bool {
	return dlist.GetHead() == node
}

func (dlist *DList) isTail(node *DNode) bool {
	return dlist.GetTail() == node
}

func (node *DNode) GetData() Object {
	return (*node).data
}

func (node *DNode) GetNext() *DNode {
	return (*node).next
}

func (node *DNode) GetPrev() *DNode {
	return (*node).prev
}

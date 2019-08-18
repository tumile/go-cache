package lru

type node struct {
	key, val interface{}
	prev     *node
	next     *node
}

type doublyLinkedList struct {
	head *node
	tail *node
	size int
}

func newDoublyLinkedList() *doublyLinkedList {
	list := doublyLinkedList{
		head: &node{},
		tail: &node{},
		size: 0,
	}
	list.head.next = list.tail
	list.tail.prev = list.head
	return &list
}

func (list *doublyLinkedList) add(node *node) {
	node.prev = list.head
	node.next = list.head.next
	list.head.next.prev = node
	list.head.next = node
	list.size++
}

func (list *doublyLinkedList) remove(node *node) {
	node.next.prev = node.prev
	node.prev.next = node.next
	list.size--
}

func (list *doublyLinkedList) removeTail() *node {
	node := list.tail.prev
	list.remove(node)
	return node
}

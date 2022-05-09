package gfl

import (
	"fmt"
)

type Node struct {
	next *Node
	prev *Node
	key  interface{}
}

type List struct {
	tail *Node
	head *Node
}

func (LinkedList *List) Insert(key interface{}) {
	new_head := &Node{
		prev: LinkedList.head,
		key:  key,
	}
	if LinkedList.head != nil {
		LinkedList.head.next = new_head
	}
	LinkedList.head = new_head

	if LinkedList.tail == nil {
		LinkedList.tail = LinkedList.head
	}
}

func (LinkedList *List) Display() {
	node := LinkedList.tail
	for node != nil {
		fmt.Printf("%+v ->", node.key)
		node = node.next
	}
	fmt.Println()
}

package linked_list

import (
	"fmt"
	"testing"
)

func TestLinkedList_PushFront(t *testing.T) {
	list := NewLinkedList()
	list.PushFront("1")
	list.PushFront("2")
	list.PushFront("3")
	list.PushFront("4")

	list.PrintList()

	fmt.Println("dropping")
	list.Pop()
	list.Pop()
	list.Pop()
	list.Pop()
	list.Pop()

	list.PrintList()
}

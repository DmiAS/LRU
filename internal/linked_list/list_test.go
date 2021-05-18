package linked_list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedList_PushFront_One(t *testing.T) {
	list := NewLinkedList()
	var key uint32 = 10
	val := "something"
	list.PushFront(key, val)

	assert.NotNil(t, list.head)
	assert.NotNil(t, list.tail)

	assert.Equal(t, list.head, list.tail)
	assert.Equal(t, list.head.key, key)
	assert.Equal(t, list.head.value, val)
}

func TestLinkedList_PushFront_Two(t *testing.T) {
	list := NewLinkedList()
	var key1, key2 uint32 = 10, 20
	val1, val2 := "something", "ex"
	list.PushFront(key1, val1)
	list.PushFront(key2, val2)

	assert.NotNil(t, list.head)
	assert.Nil(t, list.head.prev)
	assert.NotNil(t, list.head.next)

	assert.NotNil(t, list.tail)
	assert.Nil(t, list.tail.next)
	assert.NotNil(t, list.tail.prev)

	assert.NotEqual(t, list.head, list.tail)

	first := list.head
	second := list.tail

	assert.Equal(t, first.key, key2)
	assert.Equal(t, first.value, val2)

	assert.Equal(t, second.key, key1)
	assert.Equal(t, second.value, val1)

	assert.Equal(t, first.next, second)
	assert.Equal(t, second.prev, first)
}

func TestLinkedList_PushNodeFront_Nil(t *testing.T) {
	l := NewLinkedList()
	assert.NotPanics(t, func() {
		l.PushNodeFront(nil)
	})
}

func TestLinkedList_PushNodeFront_One(t *testing.T) {
	l := NewLinkedList()
	var key uint32 = 10
	val := "msg"
	node := NewNode(key, val)
	l.PushNodeFront(node)

	assert.NotNil(t, l.head)
	assert.NotNil(t, l.tail)
	assert.Equal(t, l.head, l.tail)

	listNode := l.head
	assert.Equal(t, listNode, node)
}

func TestLinkedList_PushNodeFront_Two(t *testing.T) {
	l := NewLinkedList()
	var key1, key2 uint32 = 10, 20
	val1, val2 := "msg1", "msg2"
	node1 := NewNode(key1, val1)
	node2 := NewNode(key2, val2)

	l.PushNodeFront(node1)
	l.PushNodeFront(node2)

	assert.NotNil(t, l.head)
	assert.Nil(t, l.head.prev)
	assert.NotNil(t, l.head.next)

	assert.NotNil(t, l.tail)
	assert.Nil(t, l.tail.next)
	assert.NotNil(t, l.tail.prev)

	listNode1 := l.tail
	listNode2 := l.head

	assert.Equal(t, listNode1, node1)
	assert.Equal(t, listNode2, node2)

	assert.Equal(t, listNode2.next, listNode1)
	assert.Equal(t, listNode1.prev, listNode2)
}

func TestLinkedList_Pop_Zero(t *testing.T) {
	l := NewLinkedList()
	node := l.Pop()
	assert.Nil(t, node)
	assert.Nil(t, l.head)
	assert.Nil(t, l.tail)
}

func TestLinkedList_Pop_One(t *testing.T) {
	l := NewLinkedList()
	node := NewNode(1, "msg")
	l.PushNodeFront(node)
	listNode := l.Pop()

	assert.NotNil(t, listNode)
	assert.Nil(t, l.head)
	assert.Nil(t, l.tail)

	assert.Equal(t, node, listNode)
}

func TestLinkedList_Pop_Two(t *testing.T) {
	l := NewLinkedList()
	node1 := NewNode(1, "msg1")
	node2 := NewNode(2, "msg2")
	l.PushNodeFront(node1)
	l.PushNodeFront(node2)
	listNode1 := l.Pop()

	assert.NotNil(t, listNode1)
	assert.NotNil(t, l.head)
	assert.NotNil(t, l.tail)

	assert.Equal(t, node1, listNode1)
	assert.Equal(t, l.head, node2)
	assert.Equal(t, l.tail, node2)

	listNode2 := l.Pop()

	assert.NotNil(t, listNode2)
	assert.Nil(t, l.head)
	assert.Nil(t, l.tail)

	assert.NotEqual(t, listNode1, listNode2)
}

func TestLinkedList_Unlink_Nil(t *testing.T) {
	l := NewLinkedList()
	assert.NotPanics(t, func() {
		l.Unlink(nil)
	})
}

func TestLinkedList_Unlink_One(t *testing.T) {
	l := NewLinkedList()
	node := NewNode(1, "1")
	l.PushNodeFront(node)

	l.Unlink(node)

	assert.Nil(t, l.head)
	assert.Nil(t, l.tail)

	assert.Nil(t, node.prev)
	assert.Nil(t, node.next)
}

func TestLinkedList_Unlink_Right(t *testing.T) {
	l := NewLinkedList()
	node1 := NewNode(1, "1")
	node2 := NewNode(2, "2")
	l.PushNodeFront(node1)
	l.PushNodeFront(node2)

	l.Unlink(node1)

	assert.NotNil(t, l.head)
	assert.NotNil(t, l.tail)

	assert.Nil(t, node1.prev)
	assert.Nil(t, node1.next)

	assert.Equal(t, l.head, node2)
	assert.Equal(t, l.tail, node2)

	assert.Nil(t, node2.next)
	assert.Nil(t, node2.prev)
}

func TestLinkedList_Unlink_Left(t *testing.T) {
	l := NewLinkedList()
	node1 := NewNode(1, "1")
	node2 := NewNode(2, "2")
	l.PushNodeFront(node1)
	l.PushNodeFront(node2)

	l.Unlink(node2)

	assert.NotNil(t, l.head)
	assert.NotNil(t, l.tail)

	assert.Nil(t, node2.prev)
	assert.Nil(t, node2.next)

	assert.Equal(t, l.head, node1)
	assert.Equal(t, l.tail, node1)

	assert.Nil(t, node1.next)
	assert.Nil(t, node1.prev)
}

func TestLinkedList_Unlink_Middle(t *testing.T) {
	l := NewLinkedList()
	node1 := NewNode(1, "1")
	node2 := NewNode(2, "2")
	node3 := NewNode(3, "3")
	l.PushNodeFront(node1)
	l.PushNodeFront(node2)
	l.PushNodeFront(node3)

	l.Unlink(node2)

	assert.NotNil(t, l.head)
	assert.NotNil(t, l.tail)

	assert.Nil(t, node2.next)
	assert.Nil(t, node2.prev)

	assert.Equal(t, l.head, node3)
	assert.Equal(t, l.tail, node1)

	assert.Nil(t, node1.next)
	assert.Equal(t, node1.prev, node3)

	assert.Nil(t, node3.prev)
	assert.Equal(t, node3.next, node1)
}

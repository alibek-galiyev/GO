package linkedlists

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head   *Node
	tail   *Node
	length int
}

func (l *LinkedList) InsertAtBeginning(data int) {
	newNode := &Node{data: data}
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
		l.length++
	} else {
		newNode.next = l.head
		l.head = newNode
		l.length++
	}
}

func (l *LinkedList) InsertAtEnding(data int) {
	newNode := &Node{data: data}
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
		l.length++
	} else {
		l.tail.next = newNode
		l.tail = newNode
		l.length++
	}
}

func (l *LinkedList) Display() {
	current := l.head
	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}
	fmt.Println("nil")
}

func (l *LinkedList) DeleteNode(data int) {
	if l.head == nil {
		return
	}
	if l.head.data == data {
		l.head = l.head.next
		if l.head == nil {
			l.tail = nil
		}
		return
	}

	current := l.head
	for current.next != nil {
		if current.next.data == data {
			current.next = current.next.next
			if current.next == nil {
				l.tail = current
			}
			return
		}
		current = current.next
	}
}

func MergeSortedLists(l1, l2 *LinkedList) *LinkedList {
	dummy := &Node{}
	tail := dummy
	n1 := l1.head
	n2 := l2.head
	for n1 != nil && n2 != nil {
		if n1.data <= n2.data {
			tail.next = n1
			n1 = n1.next
		} else {
			tail.next = n2
			n2 = n2.next
		}
		tail = tail.next
	}
	if n1 != nil {
		tail.next = n1
	} else {
		tail.next = n2
	}
	result := &LinkedList{}
	result.head = dummy.next
	result.length = l1.length + l2.length
	if result.head == nil {
		result.tail = nil
	} else {
		t := result.head
		for t.next != nil {
			t = t.next
		}
		result.tail = t
	}
	return result
}

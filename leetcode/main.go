package main

import (
	// mergetwosortedlists "main.go/021mergetwosortedlists"
	linkedlists "main.go/linkedlists"
)

func main() {
	l := linkedlists.LinkedList{}
	l.InsertAtBeginning(10)
	l.InsertAtBeginning(20)
	l.InsertAtEnding(30)
	l.InsertAtBeginning(99)
	l.Display()

	l.DeleteNode(20)
	l.Display()
}

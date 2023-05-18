package main

import "fmt"

type SLLNode struct {
	Next  *SLLNode
	Value int
}

func (sNode *SLLNode) SetValue(v int) {
	sNode.Value = v
}

func (sNode *SLLNode) GetValue() int {
	return sNode.Value
}

func NewSLLNode() *SLLNode {
	return new(SLLNode)
}

type SingleLinkedList struct {
	Head *SLLNode
	Tail *SLLNode
}

func NewSingleLinkedList() *SingleLinkedList {
	return new(SingleLinkedList)
}

func (list *SingleLinkedList) Add(v int) {

	newNode := &SLLNode{Value: v}
	if list.Head == nil {
		list.Head = newNode
	} else if list.Tail == list.Head {
		list.Head.Next = newNode
	} else if list.Tail != nil {
		list.Tail.Next = newNode
	}
	list.Tail = newNode
}

func (list *SingleLinkedList) String() string {
	s := ""
	for n := list.Head; n != nil; n = n.Next {
		s += fmt.Sprintf(" %d ", n.GetValue())
	}
	return s
}

func main() {
	list := NewSingleLinkedList()

	list.Add(3)
	list.Add(4)
	list.Add(5)
	list.Add(6)

	fmt.Println("Lista:", list)

}

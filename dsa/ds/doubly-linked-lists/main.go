package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
	Prev *Node
}

type DLinkedList struct {
	head *Node
}

func (dll *DLinkedList) Display() {
	temp := dll.head
	for temp != nil {
		fmt.Printf("%d -> ", temp.Val)
		temp = temp.Next
	}
	fmt.Println()
}

func (dll *DLinkedList) InsertFirst(val int) {
	node := Node{Val: val}
	node.Next = dll.head
	if dll.head != nil {
		dll.head.Prev = &node
	}
	dll.head = &node
}

func (dll *DLinkedList) InsertLast(val int) {
	node := Node{Val: val}
	temp := dll.head
	for temp.Next != nil {
		temp = temp.Next
	}
	temp.Next = &node
	node.Prev = temp
}

func (dll *DLinkedList) Insert(val, index int) {
	node := Node{Val: val}
	temp := dll.head
	for i := 1; i < index; i++ {
		temp = temp.Next
	}
	node.Next = temp.Next
	temp.Next = &node
	if temp.Next.Next != nil {
		temp.Next.Next.Prev = &node
	}
	node.Prev = temp
}

func main() {
	dll := DLinkedList{}
	dll.InsertFirst(10)
	dll.InsertFirst(20)
	dll.Display()
	dll.InsertLast(30)
	dll.Display()
	dll.Insert(40, 1)
	dll.Display()
}

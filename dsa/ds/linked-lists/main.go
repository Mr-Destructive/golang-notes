package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

type LinkedList struct {
	head *Node
	//tail *Node
	size int
}

func (ll *LinkedList) Display() {
	temp := ll.head
	for temp != nil {
		fmt.Printf("%d - > ", temp.Val)
		temp = temp.Next
	}
	fmt.Println()
}

func (ll *LinkedList) Get(index int) *Node {
	temp := ll.head
	for i := 1; i < index; i++ {
		temp = temp.Next
	}
	return temp
}

func (ll *LinkedList) InsertFirst(val int) {
	node := Node{Val: val}
	node.Next = ll.head
	ll.head = &node
	ll.size++
}

func (ll *LinkedList) InsertLast(val int) {
	node := Node{Val: val}
	temp := ll.head
	if temp == nil {
		ll.head = &node
		return
	}
	for temp.Next != nil {
		temp = temp.Next
	}
	temp.Next = &node
	ll.size++
}

func (ll *LinkedList) InsertIndex(val int, index int) {
	node := Node{Val: val}
	temp := ll.head
	if index > ll.size {
		return
	}
	for i := 1; i < index; i++ {
		temp = temp.Next
	}
	node.Next = temp.Next
	temp.Next = &node
	ll.size++
}

func (ll *LinkedList) DeleteFirst() int {
	val := ll.head.Val
	if ll.head == nil {
	}
	ll.head = ll.head.Next
	return val
}

func (ll *LinkedList) DeleteLast() int {
	node := ll.Get(ll.size - 2)
	if ll.size == 1 {
		return ll.DeleteFirst()
	}
	node.Next = nil
	return node.Val
}

func (ll *LinkedList) DeleteIndex(index int) int {
	node := ll.Get(index-1)
	val := node.Next.Val
	node.Next = node.Next.Next
	return val
}

func main() {
	ll := LinkedList{}
	ll.InsertFirst(10)
	ll.InsertLast(30)
	ll.InsertFirst(20)
	ll.InsertIndex(40, 2)
	ll.InsertIndex(50, 5)
	ll.Display()
	ll.DeleteFirst()
	ll.Display()
	ll.DeleteLast()
	ll.Display()
	ll.InsertLast(50)
	ll.InsertLast(60)
	ll.Display()
	ll.DeleteIndex(2)
	ll.Display()
}

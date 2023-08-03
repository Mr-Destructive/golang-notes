package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

type List struct {
	head *Node
}

type Stack struct {
	data List
}

func (s *Stack) Display() {
	temp := s.data.head
	for temp != nil {
		fmt.Printf("%d ", temp.Val)
		temp = temp.Next
	}
	fmt.Println()
}

func (s *Stack) IsEmpty() bool {
	return s.data.head == nil
}

func (s *Stack) Push(val int) {
	node := Node{Val: val}
	temp := s.data.head
	node.Next = temp
	s.data.head = &node
}

func (s *Stack) Pop() int {
	temp := s.data.head
	s.data.head = temp.Next
	return temp.Val
}

func main() {
	s := Stack{}
	s.Display()
	s.Push(10)
	s.Push(20)
	s.Display()
	s.Pop()
	s.Pop()
	fmt.Println(s.IsEmpty())
	s.Display()
}

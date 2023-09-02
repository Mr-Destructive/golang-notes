package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

type Stack struct {
	head *Node
}

func (s *Stack) Display() {
	temp := s.head
	for temp != nil {
		fmt.Printf("%d ", temp.Val)
		temp = temp.Next
	}
	fmt.Println()
}

func (s *Stack) IsEmpty() bool {
	return s.head == nil
}

func (s *Stack) Push(val int) {
	node := Node{Val: val}
	temp := s.head
	node.Next = temp
	s.head = &node
}

func (s *Stack) Pop() (int, error) {
	temp := s.head
	s.head = temp.Next
	if s.head == nil {
		err := fmt.Errorf("Stack is empty")
		return -1, err
	}
	return temp.Val, nil
}

func (s *Stack) Peek() (int, error) {
	temp := s.head
	if temp == nil {
		err := fmt.Errorf("Stack is empty")
		return -1, err
	}
	return temp.Val, nil
}

func main() {
	s := Stack{}
	s.Push(10)
	s.Display()
	s.Push(20)
	s.Display()
	s.Pop()
	s.Display()
	s.Pop()
	s.Display()
}

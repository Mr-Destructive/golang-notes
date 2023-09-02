package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

type List struct {
	head *Node
	tail *Node
}

type Queue struct {
	list List
}

func (s *Queue) Display() {
	temp := s.list.head
	for temp != nil {
		fmt.Printf("%d -> ", temp.Val)
		temp = temp.Next
	}
	fmt.Println()
}

func (s *Queue) IsEmpty() bool {
	return s.list.head == nil
}

func (s *Queue) Push(val int) {
	temp := &Node{Val: val}
	if s.list.head == nil {
		s.list.head = temp
		s.list.tail = temp
	} else {
		s.list.tail.Next = temp
		s.list.tail = temp
	}
}

func (s *Queue) Pop() (int, error) {
	temp := s.list.head
	if s.list.head == nil {
		err := fmt.Errorf("Queue is empty")
		return -1, err
	}
	s.list.head = temp.Next
	return temp.Val, nil
}

func (s *Queue) Peek() (int, error) {
	temp := s.list.head
	if temp == nil {
		err := fmt.Errorf("Queue is empty")
		return -1, err
	}
	return temp.Val, nil
}

func main() {
	s := Queue{}
	s.Display()
	fmt.Println(s.Peek())
	s.Push(10)
	fmt.Println(s.Peek())
	s.Pop()
	s.Push(20)
	s.Display()
	fmt.Println(s.Peek())
	s.Pop()
	s.Pop()
	s.Display()
}

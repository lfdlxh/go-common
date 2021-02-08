package dataStructural

import "sync"

type Node struct{
	Pre *Node
	Value interface{}
}

type Stack struct{
	Top *Node
	Length int
	Lock *sync.RWMutex
}
//create a new stack
func NewStack() *Stack{
	return &Stack{
		Top:    nil,
		Length: 0,
		Lock:   nil,
	}
}

//return the number of items in the stack
func (s *Stack)GetLength() int{
	return s.Length
}

//return the value of the top item in the stack
func (s *Stack)GetPeekItem() interface{}{
	if s.Length == 0{
		return nil
	}
	return s.Top.Value
}

//push a node upon the stack
func (s *Stack)Push(value interface{}){
	s.Lock.Lock()
	defer s.Lock.Unlock()
	node :=&Node{
		Pre:   s.Top,
		Value: value,
	}
	s.Top = node
	s.Length++
}

//pop the top item of the stack
func (s *Stack)Pop()*Node{
	s.Lock.Lock()
	defer s.Lock.Unlock()
	if s.Length==0{
		return nil
	}
	n :=s.Top
	s.Top = n.Pre
	s.Length--
	return n
}


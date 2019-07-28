package dataStruct

import (
	"fmt"
)

//栈的接口
type IStack interface {

	Push(Data interface{})error
	Pop()(interface{},error)
	//取出栈顶值，并不弹出
	Peek()interface{}
	Length()int
	traverse()
}

type StackError struct {
	msg string
}

func (rcv StackError)Error() string{
    return fmt.Sprintf("%s",rcv.msg)
}

//链栈的实现
type Node struct {
	Data interface{}
	Next *Node
}

type Stack struct {
	length int
	head *Node
}

//对栈的初始化
func CreateStack()*Stack{
	stack := new(Stack)
	stack.head = new(Node)
	return stack
}

func (s *Stack)Push(Data interface{})error{
	node := new(Node)
	node.Data = Data

	//头插来模拟后进先出
	node.Next = s.head.Next
	s.head.Next = node
	s.length++
	return nil
}

func (s *Stack)Pop()(interface{},error){
	var ele interface{}

	if s.length ==0{
		return nil,StackError{msg:"栈为空了"}
	}

	ele = s.head.Next.Data
	s.head.Next = s.head.Next.Next
	s.length--
	return ele,nil
}

func (s *Stack)Length()int{
	return s.length
}

func (s *Stack)Peek()interface{}{
	return s.head.Next.Data
}

//栈的遍历
func (s *Stack)Traverse(){

	node := s.head.Next
	for node !=nil{
		fmt.Println(node.Data)
		node = node.Next
	}
}



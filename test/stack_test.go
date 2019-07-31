package test

import (
	"algorithm/dataStruct"
	"testing"
)

func TestStackPush(t *testing.T) {
	stack := dataStruct.CreateStack()

	stack.Push("a")
	stack.Push("b")
	stack.Push("c")
	stack.Push("d")

	stack.Traverse()
}

func TestStackPop(t *testing.T) {
	stack := dataStruct.CreateStack()

	stack.Push("a")
	stack.Push("b")
	stack.Push("c")
	stack.Push("d")

	//弹出的元素为
	t.Log(stack.Pop())

	stack.Traverse()
}

func TestStackLength(t *testing.T) {
	stack := dataStruct.CreateStack()

	stack.Push("a")
	stack.Push("b")
	stack.Push("c")
	stack.Push("d")

	t.Log("栈的长度为", stack.Length())
}

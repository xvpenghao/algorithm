package test

import (
	"algorithm/dataStruct"
	"testing"
)

func TestAdd (t *testing.T){
	L := new(dataStruct.LinkList)

	L.Add("a")
	L.Add("b")
	L.Add("c")
}

func TestInsert (t *testing.T){
	L := new(dataStruct.LinkList)

	L.Add("a")
	L.Add("b")
	L.Add("c")

	//插入的到第一个
	L.InsertNode(1,"aa")

	//插入到中间最后
	L.InsertNode(2,"bb")
}

func TestDelete (t *testing.T){
	L := new(dataStruct.LinkList)

	L.Add("a")
	L.Add("b")
	L.Add("c")

	L.Delete(1)
}

func TestGetValueLocation(t *testing.T){
	L := new(dataStruct.LinkList)

	L.Add("a")
	L.Add("b")
	L.Add("c")

	index,_ := L.GetValueLocation("b")
	t.Logf("%v",index)
}

func TestReverse(t *testing.T){
	L := new(dataStruct.LinkList)
	L.Add("a")
	L.Add("b")
	L.Add("c")
	L.Add("d")

	L2,_ := L.Reverse()
	node,_ := L2.Get(1)
	t.Logf("%v",node.Data)
}
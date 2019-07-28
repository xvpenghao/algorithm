package test

import (
	"algorithm/dataStruct"
	"testing"
)

func TestReverseBinaryTree(t *testing.T){
    Tree := dataStruct.CreateBinaryTree()
    dataStruct.ReverseBinaryTree(Tree)
}

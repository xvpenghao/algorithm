package main

import (
	"algorithm/dataStruct"
	"fmt"
)

func main() {
	Tree := dataStruct.CreateBinaryTree()
	//dataStruct.PreTraverse(Tree)
	fmt.Println(dataStruct.IsCompleteBinaryTree(Tree))
}

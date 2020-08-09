package main

import "fmt"

func main() {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	root := buildTree(preorder, inorder)
	fmt.Println(root)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//剑指 Offer 07. 重建二叉树
//https://leetcode-cn.com/problems/zhong-jian-er-cha-shu-lcof/
//(参考图解分析)[https://www.cnblogs.com/jiaxin359/p/9512348.html]
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	//中序 位置和元素的映射，根据先序元素，从中序中能得到 根元素 和左子树元素，右子树元素
	//使用map存放 中序下标和元素的映射，便于获取先序元素在 中序中的一些信息
	inorderMap := make(map[int]int, len(inorder))
	for index, val := range inorder {
		inorderMap[val] = index
	}

	return buildBinaryTree(preorder, 0, len(preorder)-1, inorder, 0, len(inorder)-1, inorderMap)
}

func buildBinaryTree(preorder []int, preStartIndex, preEndIndex int, inorder []int, inStartIndex, inEndIndex int, inorderMap map[int]int) *TreeNode {
	if preStartIndex > preEndIndex {
		return nil
	}
	//也有可能取出的一个值
	rootNode := &TreeNode{
		Val: preorder[preStartIndex],
	}
	if preStartIndex == preEndIndex {
		return rootNode
	}

	//根节点 左右子树有多个元素则需要计算
	rootIndex := inorderMap[rootNode.Val]
	//计算的 根节点左右元素个数
	leftNodes, rightNodes := rootIndex-inStartIndex, inEndIndex-rootIndex
	rootNode.Left = buildBinaryTree(preorder, preStartIndex+1, preStartIndex+leftNodes, inorder, inStartIndex, rootIndex-1, inorderMap)
	//preEndIndex-rightNodes+1 根节点已经添加了，+1 跳过
	rootNode.Right = buildBinaryTree(preorder, preEndIndex-rightNodes+1, preEndIndex, inorder, rootIndex+1, inEndIndex, inorderMap)
	return rootNode
}

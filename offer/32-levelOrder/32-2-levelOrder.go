package main

import (
	"container/list"
	"fmt"
)

/**
*@Author:徐鹏豪
*@Time: 2021/2/17 16:29
*@Description:
 */
func main() {
	root := &TreeNode2{
		Val: 3,
		Left: &TreeNode2{
			Val:   9,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode2{
			Val: 20,
			Left: &TreeNode2{
				Val:   15,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode2{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}

	res := ll_levelOrder2(root)
	for _, val := range res {
		fmt.Println(val)
	}
}

type TreeNode2 struct {
	Val   int
	Left  *TreeNode2
	Right *TreeNode2
}

// 剑指 Offer 32 - II. 从上到下打印二叉树 II
// https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-ii-lcof/
func ll_levelOrder2(root *TreeNode2) [][]int {
	if root == nil {
		return nil
	}

	var result [][]int
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		var eleList []int
		// 临时存储每层的节点数
		for i := queue.Len(); i > 0; i-- {
			node := queue.Remove(queue.Front()).(*TreeNode2)
			eleList = append(eleList, node.Val)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		result = append(result, eleList)
	}

	return result
}

func ll_levelOrder(root *TreeNode2) [][]int {
	// 层次遍历，当一层遍历完毕，则将该层添加的元素，添加到 二维数组中
	// 如何确定该层，遍历完毕呢？？
	// 如何才能知道，这一层的最后一个节点呢？
	if root == nil {
		return nil
	}

	queue := list.New()
	queue.PushBack(root)
	var result [][]int
	var eleList []int
	lastNode, curLastNode := root, root
	for queue.Len() > 0 {
		node := queue.Remove(queue.Front()).(*TreeNode2)
		eleList = append(eleList, node.Val)

		if node.Left != nil {
			curLastNode = node.Left
			queue.PushBack(node.Left)
		}

		if node.Right != nil {
			queue.PushBack(node.Right)
			curLastNode = node.Right
		}

		if node == lastNode {
			result = append(result, eleList)
			eleList = nil
			lastNode = curLastNode
			curLastNode = nil
		}
	}
	return result
}

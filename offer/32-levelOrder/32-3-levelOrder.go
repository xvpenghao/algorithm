package main

import (
	"container/list"
	"fmt"
)

/**
*@Author:徐鹏豪
*@Time: 2021/3/7 15:03
*@Description:
 */

// 从上到下打印二叉树 III
// https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-iii-lcof/
func main() {
	root := &TreeNode3{
		Val: 3,
		Left: &TreeNode3{
			Val:   9,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode3{
			Val: 20,
			Left: &TreeNode3{
				Val:   15,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode3{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}

	fmt.Println(levelOrder4(root))
}

type TreeNode3 struct {
	Val   int
	Left  *TreeNode3
	Right *TreeNode3
}

// 左右打印 二叉树，
// 打印每一层，控制结果的存储方式，是左还是右，
// 切换时机，换层时候，默认是左
func levelOrder4(root *TreeNode3) [][]int {
	if root == nil {
		return nil
	}
	// 如果是独苗,则直接返回该对象
	if root.Left == nil && root.Right == nil {
		return [][]int{{root.Val}}
	}

	queue := list.New()
	queue.PushBack(root)
	var result [][]int
	for queue.Len() > 0 {
		// 存放每一层的元素
		tmpQueue := list.New()
		var eleList []int
		for n := queue.Len(); n > 0; n-- {
			// 删除，然后添加节点
			node := queue.Remove(queue.Front()).(*TreeNode3)
			if len(result)%2 != 0 { // 右
				tmpQueue.PushFront(node.Val)
			} else { // 左
				tmpQueue.PushBack(node.Val)
			}
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}

		for tmpQueue.Len() > 0 {
			val := tmpQueue.Remove(tmpQueue.Front()).(int)
			eleList = append(eleList, val)
		}
		result = append(result, eleList)
	}

	return result
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 左右打印 二叉树，
// 打印每一层，控制结果的存储方式，是左还是右，
// 切换时机，换层时候，默认是左
func levelOrder3(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	// 如果是独苗,则直接返回该对象
	if root.Left == nil && root.Right == nil {
		return [][]int{{root.Val}}
	}

	queue := list.New()
	queue.PushBack(root)
	var result [][]int
	direct := 1 // 取反，大于零 左，小于零，右
	for queue.Len() > 0 {
		// 存放每一层的元素
		var eleList []int
		for n := queue.Len(); n > 0; n-- {
			// 删除，然后添加节点
			node := queue.Remove(queue.Front()).(*TreeNode)
			eleList = append(eleList, node.Val)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}

		if direct < 0 {
			eleNewList := make([]int, len(eleList))
			length := len(eleList)
			for index, ele := range eleList {
				eleNewList[length-1-index] = ele
			}
			eleList = eleNewList
		}
		result = append(result, eleList)
		direct = ^direct

	}
	return result
}

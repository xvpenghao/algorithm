package main

import (
	"container/list"
	"fmt"
)

/**
*@Author:徐鹏豪
*@Time: 2021/2/15 10:21
*@Description:
 */

func main() {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   9,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val:   15,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}

	fmt.Println(levelOrder(root))
}

// https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-lcof/
// 剑指 Offer 32 - I. 从上到下打印二叉树

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 二叉树的层次遍历
/*
     3
   / 	\
  9  	20
 / \  	/  \
4  	6	15   7
result [3,9,20,4,6,15,7]
*/
// 为什么 想到了用 队列做着呢？s
// 分析 3，9，20，4，6 节点关系，【4，6】是9的子节点，20遍历完毕才能遍历 【4.6】
// 需要在遍历 9的时候把 【4，6】放在 20的后面，
// 怎样才能放在后面呢？
// 即 9，20先进，然后 4，6再进，输出 9，20，4，6
// 符合队列的特性 先进先出

func levelOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var result []int
	// 使用队列完成 层次的遍历
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		node := queue.Remove(queue.Front()).(*TreeNode)
		result = append(result, node.Val)
		if node.Left != nil {
			queue.PushBack(node.Left)
		}
		if node.Right != nil {
			queue.PushBack(node.Right)
		}
	}
	return result
}

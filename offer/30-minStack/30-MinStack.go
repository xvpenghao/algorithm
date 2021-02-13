package main

import "fmt"

/**
*@Author:徐鹏豪
*@Time: 2021/2/13 10:40
*@Description: Offer 30. 包含min函数的栈
 */
func main() {
	ms := Constructor()
	ms.Push(2)
	ms.Push(0)
	ms.Push(3)
	ms.Push(0)
	fmt.Println(ms.Min())
	ms.Pop()
	fmt.Println(ms.Min())
	ms.Pop()
	fmt.Println(ms.Min())
	ms.Pop()
	fmt.Println(ms.Min())
}

// 剑指 Offer 30. 包含min函数的栈
// https://leetcode-cn.com/problems/bao-han-minhan-shu-de-zhan-lcof/
type MinStack struct {
	A []int // 主栈
	B []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	this.A = append(this.A, x)
	if len(this.A) == 1 {
		this.B = append(this.B, x)
		return
	}
	// 里面可能有重复的值
	if x <= this.B[len(this.B)-1] {
		this.B = append(this.B, x)
	}
}

// 弹出栈顶元素 4 2 7 3 9 1
// 1 -> 2, 9->2, 3->2, 7->2, 2->4
// 4 2 1 // 当弹出的元素和 和B中的元素相等是 则B则弹出
// 如何实现在 弹出 1后 2是最小的元素
// 如何利用 空间 换时间 仔细分析后，你会发现，有些弹出的元素，是不会影响到 最小值的发挥的
func (this *MinStack) Pop() {
	if len(this.A) == 0 {
		return
	}
	aPopVal := this.A[len(this.A)-1]
	bPopVal := this.B[len(this.B)-1]
	// 弹出A元素中的内容
	this.A = this.A[:len(this.A)-1]
	if aPopVal == bPopVal {
		this.B = this.B[:len(this.B)-1]
	}
}

// 获取栈顶元素的值
func (this *MinStack) Top() int {
	if len(this.A) == 0 {
		return 0
	}
	return this.A[len(this.A)-1]
}

// 每次push和pop都要重新计算一下
func (this *MinStack) Min() int {
	if len(this.B) == 0 {
		return 0
	}
	return this.B[len(this.B)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */

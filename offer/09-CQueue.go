package main

import "fmt"

//剑指 Offer 09. 用两个栈实现队列
//https://leetcode-cn.com/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof/
func main() {
	cq := Constructor()
	cq.AppendTail(3)
	fmt.Println(cq.DeleteHead())
	fmt.Println(cq.DeleteHead())
	fmt.Println(cq.DeleteHead())
}

type CQueue struct {
	//当做栈来用
	AppendStack []int
	DeleteStack []int
}

func Constructor() CQueue {
	return CQueue{}
}

func (this *CQueue) AppendTail(value int) {
	this.AppendStack = append(this.AppendStack, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.AppendStack) == 0 && len(this.DeleteStack) == 0 {
		return -1
	}
	//AppendStack 元素弹出到-> DeleteStack
	for i := len(this.AppendStack) - 1; i >= 0; i-- {
		this.DeleteStack = append(this.DeleteStack, this.AppendStack[i])
	}
	//AppendStack 元素弹出完毕，则栈置为空
	this.AppendStack = nil
	//删除栈元素出栈
	value := this.DeleteStack[len(this.DeleteStack)-1]
	this.DeleteStack = this.DeleteStack[:len(this.DeleteStack)-1]

	//将删除栈的元素弹出到->AppendStack
	for i := len(this.DeleteStack) - 1; i >= 0; i-- {
		this.AppendStack = append(this.AppendStack, this.DeleteStack[i])
	}
	this.DeleteStack = nil
	//返回删除元素
	return value
}

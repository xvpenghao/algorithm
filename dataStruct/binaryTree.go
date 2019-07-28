package dataStruct

import (
	"container/list"
	"fmt"
)

type BinaryTree struct {
	Data interface{}
	Left *BinaryTree
	Right *BinaryTree
}

/**
*@Author:徐鹏豪
*@Time:2019/7/24 0024
*@Description:创建二叉树
*/
func CreateBinaryTree()*BinaryTree{
    //根据先序创建二叉树
	//AB##C##
	var Tree *BinaryTree
	var data string
	fmt.Scanln(&data)
	if data == "#"{
		return nil
	}

	//开始创建二叉树
	Tree = new(BinaryTree)
	Tree.Data = data
	Tree.Left = CreateBinaryTree()
	Tree.Right = CreateBinaryTree()
	return Tree
}

/**
*@Author:徐鹏豪
*@Time:2019/7/24 0024
*@Description:前序遍历
*/
func PreTraverse(Tree *BinaryTree){
    if Tree == nil{
    	return
	}
    fmt.Println(Tree.Data)
    PreTraverse(Tree.Left)
    PreTraverse(Tree.Right)
}

/**
*@Author:徐鹏豪
*@Time:2019/7/24 0024
*@Description:中序遍历
*/
func MidTraverse(Tree *BinaryTree){
	if Tree ==nil{
		return
	}
	MidTraverse(Tree.Left)
	fmt.Println(Tree.Data)
	MidTraverse(Tree.Right)
}

/**
*@Author:徐鹏豪
*@Time:2019/7/24 0024
*@Description:后序遍历
*/
func PostOrderTraverse(Tree *BinaryTree){
    if Tree == nil{
    	return
	}
    PostOrderTraverse(Tree.Left)
    PostOrderTraverse(Tree.Right)
    fmt.Println(Tree.Data)
}

/**
*@Author:徐鹏豪
*@Time:2019/7/24 0024
*@Description:求二叉树的高度
*/
func GetBinaryTreeHeight(Tree *BinaryTree)int{
    if Tree ==nil{
    	return 0
	}
    leftHeight := GetBinaryTreeHeight(Tree.Left)
    rightHeight := GetBinaryTreeHeight(Tree.Right)

    return getMax(leftHeight,rightHeight)+1
}
/**
*@Author:徐鹏豪
*@Time:2019/7/24 0024
*@Description:求最大值
*/
func getMax(n1,n2 int)int{
    if n1 >n2{
    	return n1
	}
    return n2
}

/**
*@Author:徐鹏豪
*@Time:2019/7/24 0024
*@Description:递归反转，二叉树的反转
*/
func ReverseBinaryTree(Tree *BinaryTree){
    if Tree ==nil{
    	return
	}
    ReverseBinaryTree(Tree.Left)
    ReverseBinaryTree(Tree.Right)
    Tree.Left,Tree.Right = Tree.Right,Tree.Left
}

/**
*@Author:徐鹏豪
*@Time:2019/7/24 0024
*@Description:递推二叉树的反转-利用栈
*/
func ReverseBinaryTree2(Tree *BinaryTree){
	//利用栈完成二叉树的反转
	stack := []*BinaryTree{}
	stack = append(stack,Tree)
	for len(stack) >0 {
		//取出栈顶元素
		node := stack[len(stack)-1]
		//弹出元素
		stack = stack[0:len(stack)-1]
		if node !=nil{
			stack = append(stack,node.Left)
			stack = append(stack,node.Right)
			//交换元素
			node.Left,node.Right = node.Right,node.Left
		}
	}
}

/**
*@Author:徐鹏豪
*@Time:2019/7/28 0028
*@Description:递推二叉树的反转-利用队列
*/
func ReverseBinaryTree3(tree *BinaryTree){
	if tree == nil{
		return
	}
	//层次遍历交换节点
	queue := list.New()
	queue.PushBack(tree)

	for queue.Len()>0{
		t := queue.Remove(queue.Front()).(*BinaryTree)
		t.Left,t.Right = t.Right , t.Left
		if t.Left !=nil{
			queue.PushBack(t.Left)
		}
		if t.Right !=nil{
			queue.PushBack(t.Right)
		}
	}
}

/**
*@Author:徐鹏豪
*@Time:2019/7/24 0024
*@Description:非递归前序遍历
*/
func IterationPreTraverse(Tree *BinaryTree){
    T2 := Tree
    stack := []*BinaryTree{}
    //根左右
	for len(stack)>0 || T2 !=nil{
		for T2 !=nil{
			fmt.Println(T2.Data)
			stack = append(stack,T2.Left)
			T2 = T2.Left
		}
		//弹出元素
		T2 = stack[len(stack)-1]
		stack = stack[0:len(stack)-1]
		//这里弹出的可能有nil，T2.nil就会报错了
		T2 = T2.Right
	}
}

/**
*@Author:徐鹏豪
*@Time:2019/7/24 0024
*@Description:迭代中序遍历
*/
func IterationMidTraverse(Tree *BinaryTree){
	T2 := Tree
	stack := []*BinaryTree{}
	for len(stack) >0 || T2 !=nil{
		for T2 !=nil{
			stack = append(stack,T2)
			T2 = T2.Left
		}

		T2 = stack[len(stack)-1]
		stack = stack[0:len(stack)-1]
		fmt.Println(T2.Data)
		T2 = T2.Right
	}
}

/**
*@Author:徐鹏豪
*@Time:2019/7/24 0024
*@Description:层次遍历,利用队列实现层次遍历
*/
func LevelTraverse(Tree *BinaryTree){
    queue := list.New()
    queue.PushBack(Tree)

    var temp *BinaryTree
    for queue.Len() >0{
		temp = queue.Front().Value.(*BinaryTree)
    	fmt.Println(temp.Data)
		queue.Remove(queue.Front())

		if temp.Left !=nil{
			queue.PushBack(temp.Left)
		}
		if temp.Right !=nil{
			queue.PushBack(temp.Right)
		}
	}
}

/**
*@Author:徐鹏豪
*@Time:2019/7/24 0024
*@Description:判断一个二叉树是否一个完全二叉树
完全二叉树的在做层级遍历是，空节点nil一定在最后，如果发现在中间
则不是完全二叉树
*/
func IsCompleteBinaryTree(Tree *BinaryTree)bool{
    queue := list.New()
    queue.PushBack(Tree)

    for queue.Len() >0{
    	ele := queue.Front()
    	//(nil).(type),会报错，而这里没有报错，是因为
    	//它的类型不为nil，值为nil，所以可以转为type
		t2 := queue.Remove(ele).(*BinaryTree)
		//如果发现为空，则退出循环，
		if t2 ==nil{
			break
		}
		queue.PushBack(t2.Left)
		queue.PushBack(t2.Right)
	}

    //如果后面还有元素的，则说明不是玩去二叉树
    //例如，a b c nil d e,
    //
    for queue.Len()>0{
		ele := queue.Front()
		t2 := queue.Remove(ele).(*BinaryTree)
		if t2 !=nil{
			return false
		}
	}

    return true
}

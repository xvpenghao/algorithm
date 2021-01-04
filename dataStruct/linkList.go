package dataStruct

import "fmt"

// 双向链表

// 单向链表，头和尾插发
type LinkList struct {
	length int
	head   *Node
}

type LinkListError struct {
	format string
	msg    string
}

func (s LinkListError) Error() string {
	return fmt.Sprintf(s.format, s.msg)
}

func (L *LinkList) addToFirst(data interface{}) {
	node := new(Node)
	node.Data = data
	node.Next = L.head
	L.head = node
	L.length++
}

func (L *LinkList) addToLast(data interface{}) {
	node := new(Node)
	node.Data = data
	node.Next = nil

	L.length++

	// 找到最后一个节点
	L2 := L.head
	for L2.Next != nil {
		L2 = L2.Next
	}
	L2.Next = node

}

func (L *LinkList) Get(index int) (*Node, error) {
	if index <= 0 || index > L.length {
		return nil, LinkListError{"%s", "参数不合法"}
	}

	// 因为你第一个节点也有值
	current := L.head

	for index > 1 {
		current = current.Next
		index--
	}

	return current, nil
}

// 得到前一个，是为了在指定的位置添加
func (L *LinkList) getPerv(index int) (*Node, error) {
	if index == 0 || index > L.length {
		return nil, LinkListError{"%s", "参数不合法"}
	}

	L2 := L.head

	for index > 2 {
		L2 = L2.Next
		index--
	}

	return L2, nil
}

func (L *LinkList) InsertNode(index int, data interface{}) error {

	// 添加到第一个
	if index <= 1 {
		L.addToFirst(data)
		return nil
	}
	// 添加到最后
	if index >= L.length {
		L.addToLast(data)
		return nil
	}

	// 则插入到中间
	prevNode, err := L.getPerv(index)
	if err != nil {
		return err
	}

	// 节点的插入
	node := new(Node)
	node.Data = data
	node.Next = prevNode.Next
	prevNode.Next = node
	L.length++

	return nil
}

// 这样节点的添加的就舒服了
func (L *LinkList) Add(data interface{}) {
	if L.length == 0 {
		L.addToFirst(data)
		return
	}

	L.addToLast(data)
}

func (L *LinkList) Delete(index int) error {
	// 接受者不能为空，需要判断一下
	if L == nil {
		return LinkListError{"%s", "接受者不能为空"}
	}
	if index <= 0 || index > L.length {
		return LinkListError{"%s", "参数不合法"}
	}

	// 删除第一个节点有点诡异其他还行
	if index == 1 {
		L.head = L.head.Next
		L.length--
		return nil
	}
	prevNode, err := L.getPerv(index)
	if err != nil {
		return err
	}

	prevNode.Next = prevNode.Next.Next
	L.length--
	return nil
}

// 单链表的反转
// a->b->c->d
// b->a->c->d
// c->b->a->d
// d->c->b->a
func (L *LinkList) Reverse() (*LinkList, error) {

	if L == nil || L.length == 1 {
		return nil, LinkListError{"%s", "无需反转"}
	}

	var tmp = &Node{-1, L.head}
	var prev = tmp.Next
	var current = prev.Next

	// 从第二个开始
	for current != nil {
		// a c d
		prev.Next = current.Next
		// b a c d
		current.Next = tmp.Next
		// b a c d
		tmp.Next = current
		// c
		current = prev.Next
	}

	return &LinkList{L.length, tmp.Next}, nil
}

// 获取指定值的位置
func (L *LinkList) GetValueLocation(data interface{}) (int, error) {
	if L == nil || data == nil {
		return -1, LinkListError{"%s", "参数或者接受者错误"}
	}

	if L.length == 0 {
		return -1, LinkListError{"%s", "链表为空"}
	}

	location := L.length

	L2 := L.head

	for location > 1 && L2.Data != data {
		L2 = L2.Next
		location--
	}
	return location, nil
}

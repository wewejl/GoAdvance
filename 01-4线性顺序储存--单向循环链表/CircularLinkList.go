package main

import "fmt"

//创建循环链表结点结构体
type CircularLinkNode struct {
	Data interface{}
	Next *CircularLinkNode
}

// 创建循环链表
func (node *CircularLinkNode) Create(Data ...interface{}) {
	//容错
	if node == nil || len(Data) == 0 {
		return
	}
	//保存一下头结点地址
	start := node
	//循环获取数据 并创建新数据
	for _, v := range Data {
		newNode := new(CircularLinkNode)
		newNode.Data = v
		newNode.Next = nil
		node.Next = newNode
		node = node.Next
	}

	//最后一个数据的指定要到第一个数据结点
	node.Next = start.Next
}

// 打印循环链表
func (node *CircularLinkNode) Print() {
	//容错
	if node == nil {
		return
	}
	//保存要停止的位置
	temp := node.Next
	//node=node.Next
	//循环  结束条件是node.Next!=temp
	for {
		node = node.Next
		if node.Data != nil {
			fmt.Print(node.Data, " ")
		}
		if node.Next == temp {
			break
		}
	}
}

// 获取循环链表长度
func (node *CircularLinkNode) Length() int {
	if node == nil {
		return -1
	}
	//保存起始位置
	start := node
	//定义一个变量记录数据
	i := 0
	//循环遍历
	for {
		i++
		node = node.Next
		if start.Next == node.Next {
			break
		}
	}
	return i

}

// 插入链表结点
func (node *CircularLinkNode) Insert(index int, Data interface{}) {
	//容错
	if node == nil || index <= 0 || index > node.Length() {
		return
	}
	start := node.Next

	preNode := node
	for i := 0; i < index; i++ {
		preNode = node
		node = node.Next
	}
	NewNode := new(CircularLinkNode)
	NewNode.Data = Data
	NewNode.Next = node
	preNode.Next = NewNode
	if index == 1 {
		for {
			if start == node.Next { // 找到 尾结点.
				break
			}
			node = node.Next
		}
		node.Next = NewNode
	}

}

// 删除链表结点
func (node *CircularLinkNode) Delete(index int) {
	if node == nil {
		return
	}
	if index <= 0 || index > node.Length() {
		return
	}

	//定义标记位置
	start := node.Next
	// 定义preNode, 用来标记index 对应结点的前一个结点
	preNode := node
	for i := 0; i < index; i++ {
		preNode = node
		node = node.Next
	}

	//删除index=1时  要把尾结点指向重定向
	if index == 1 {
		temp := node
		for {
			if temp.Next == start {
				break
			}
			temp=temp.Next
		}
		temp.Next=node.Next
	}
	//开始删除
	preNode.Next = node.Next

	// 置空
	node.Data = nil
	node.Next = nil
	node = nil

}

func main() {
	steck := new(CircularLinkNode)
	steck.Create(1, 2, 3, 4, 5, 6, 7, 8, 9)
	steck.Print()
	fmt.Println()
	fmt.Println(steck.Length())
	//steck.Insert(1, 2)
	//fmt.Println()

	steck.Delete(3)
	steck.Print()
}

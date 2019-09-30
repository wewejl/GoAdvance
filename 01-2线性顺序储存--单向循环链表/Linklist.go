package main

import (
	"fmt"
	"reflect"
)

type LinkNode struct {
	Data interface{} // 数据域（值域）
	Next *LinkNode   // Next LinkNode 这种定义错误！无法开辟内存。
}

//创建单向链表Create(...数据interface{})
func (node *LinkNode) Create(Data ...int) {
	//容错
	if node == nil || len(Data) == 0 {
		return
	}
	//进行数据循环
	for _, v := range Data {
		newNode := new(LinkNode)
		newNode.Data = v
		newNode.Next = nil
		node.Next = newNode
		node = newNode
	}
}

//打印单向链表Print()
func (node *LinkNode) Print() {
	for {
		newNode := node.Next
		if newNode == nil {
			return
		}
		fmt.Print(newNode.Data, " ")
		node = newNode
	}
}

//递归打印
func (node *LinkNode) Print1() {
	//递归停止的条件
	if node == nil {
		return
	}
	//在递归调用前进行打印
	if node.Data != nil {
		fmt.Print(node.Data, " ")
	}
	//进行递归调用
	node.Next.Print1()

}

//获取单向链表的长度Length
func (node *LinkNode) Length() int {

	if node == nil {
		return -1
	}
	i := 0
	for newNode := node.Next; newNode != nil; i++ {
		newNode = node.Next
		node = newNode
	}
	return i - 1
}

//单向列表插入--头插法insertByHead(Data interface{})
func (node *LinkNode) insertByHead(Data interface{}) {
	//
	if node == nil {
		return
	}
	newNode := new(LinkNode)
	newNode.Next = node.Next
	newNode.Data = Data
	node.Next = newNode
}

//单向列表插入--尾插法insertByTail(Data interface)
func (node *LinkNode) insertByTail(Data interface{}) {
	if node == nil || Data == nil {
		return
	}
	newNode := new(LinkNode)
	newNode.Data = Data
	newNode.Next = nil
	for node.Next != nil {
		node = node.Next
	}

	node.Next = newNode
}

//单向链表插入--按位置插入insertByindex()
func (node *LinkNode) insertByindex(Data interface{}, index int) {

	if node == nil || index <= 1 || index > node.Length() || Data == nil {
		return
	}

	newNode := new(LinkNode)
	newNode.Data = Data
	newNode.Next = nil
	i := 0
	for node.Next != nil {
		i++
		node = node.Next
		if index-1 == i {
			newNode.Next = node.Next
			node.Next = newNode
			return
		}
	}
}

//删除链表结点--按位置删除 DeleteByindex()
func (node *LinkNode) DeleteByindex(index int) {
	if node == nil {
		return
	}
	index--
	if index == 0 {
		node.Next = node.Next.Next
		return
	}
	node.Next.DeleteByindex(index)

}

//删除链表结点--按数据删除 DeleteByData()
func (node *LinkNode) DeleteByData(Data interface{}) {
	if node == nil || Data == nil {
		return
	}
	var prenode *LinkNode
	for node.Next!=nil {
		prenode=node
		node=node.Next
		if reflect.DeepEqual(node.Data,Data) {
			prenode.Next=node.Next
			node.Data = nil
			node.Next = nil
			node = nil
			return
		}
	}

}

func main() {
	lis := new(LinkNode)
	lis.Create(1, 1, 2, 3, 4, 5, 6)

	lis.Print()
	lis.DeleteByData(4)
	fmt.Println()
	lis.Print()
}
package main

import (
	"fmt"
	"reflect"
	"runtime"
)

type TwoWayLinkList struct {
	Data interface{}
	Prev *TwoWayLinkList
	Next *TwoWayLinkList
}

func (node *TwoWayLinkList) Create(Data ...interface{}) {
	if node == nil {
		return
	}
	for _, v := range Data {
		NewNode := new(TwoWayLinkList)
		NewNode.Data = v
		NewNode.Prev = node
		NewNode.Next = nil

		node.Next = NewNode
		node = node.Next
	}
}

//打印双向链表
func (node *TwoWayLinkList) Print() {
	if node == nil {
		return
	}
	if node.Data != nil {
		fmt.Print(node.Data, " ")
	}
	node.Next.Print()
}

func (node *TwoWayLinkList) Print1() {
	if node == nil {
		return
	}
	for node.Next != nil {
		node = node.Next
	}
	for node.Prev != nil {
		if node.Data != nil {
			fmt.Print(node.Data, " ")
		}
		node = node.Prev
	}
	fmt.Println()
}

//获取双向链表长度
func (node *TwoWayLinkList) Length() int {
	if node == nil {
		return -1
	}
	i := 0
	for node.Next != nil {
		i++
		node = node.Next
	}
	return i
}

//插入链表结点 insert()
func (node *TwoWayLinkList) insert(Data interface{}, index int) {
	//容错
	if node == nil || Data == nil || index < 0 {
		return
	}

	//创建新结点
	newTwoWayLinkList := new(TwoWayLinkList)
	newTwoWayLinkList.Data = Data
	newTwoWayLinkList.Next = nil
	newTwoWayLinkList.Prev = nil
	//进行数据循环
	len := node.Length()
	if index >= len {
		for i := 0; i < len; i++ {
			node = node.Next
		}
		node.Next = newTwoWayLinkList
		newTwoWayLinkList.Prev = node
	} else {
		for i := 0; i < index; i++ {
			node = node.Next
		}

		newTwoWayLinkList.Next = node
		newTwoWayLinkList.Prev = node.Prev
		node.Prev.Next = newTwoWayLinkList
		node.Prev = newTwoWayLinkList
	}

}

//删除链表结点 Delete() 下标
func (node *TwoWayLinkList) Delete(index int) {
	//容错
	if node == nil || index <= 0 || index > node.Length() {
		return
	}
	node = node.Next
	index--
	node.Delete(index)
	if index == 0 {
		node.Prev.Next = node.Next
		node.Next.Prev = node.Prev
		node.Data = nil
		node.Prev = nil
		node.Next = nil
	}

}

//删除链表结点Delete() 数据
func (node *TwoWayLinkList) Delete1(Data interface{}) {
	if node == nil || Data == nil {
		return
	}
	len := node.Length()
	for i := 0; i < len; i++ {
		node = node.Next
		if reflect.DeepEqual(node.Data, Data) {
			break
		}
	}
	node.Next.Prev = node.Prev
	node.Prev.Next = node.Next
	node.Next = nil
	node.Prev = nil
	node.Data = nil
}

//销毁双向链表 Destroy()
func (node *TwoWayLinkList) Destroy() {
	//容错
	if node == nil {
		return
	}
	//循环
	for node.Next != nil {
		node = node.Next
		Destorynode:=node
		Destorynode.Next=nil
		Destorynode.Data=nil
		Destorynode.Prev=nil
	}
	runtime.GC()

}
func main() {
	twolis := new(TwoWayLinkList)
	twolis.Create(1, 2, 3, 4, 5, 6, 7, 8)

	twolis.Destroy()
	fmt.Println()
	twolis.Print()
	fmt.Println("length :=", twolis.Length())

}

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

//��ӡ˫������
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

//��ȡ˫��������
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

//���������� insert()
func (node *TwoWayLinkList) insert(Data interface{}, index int) {
	//�ݴ�
	if node == nil || Data == nil || index < 0 {
		return
	}

	//�����½��
	newTwoWayLinkList := new(TwoWayLinkList)
	newTwoWayLinkList.Data = Data
	newTwoWayLinkList.Next = nil
	newTwoWayLinkList.Prev = nil
	//��������ѭ��
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

//ɾ�������� Delete() �±�
func (node *TwoWayLinkList) Delete(index int) {
	//�ݴ�
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

//ɾ��������Delete() ����
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

//����˫������ Destroy()
func (node *TwoWayLinkList) Destroy() {
	//�ݴ�
	if node == nil {
		return
	}
	//ѭ��
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

package main

import (
	"fmt"
	"reflect"
)

type LinkNode struct {
	Data interface{} // ������ֵ��
	Next *LinkNode   // Next LinkNode ���ֶ�������޷������ڴ档
}

//������������Create(...����interface{})
func (node *LinkNode) Create(Data ...int) {
	//�ݴ�
	if node == nil || len(Data) == 0 {
		return
	}
	//��������ѭ��
	for _, v := range Data {
		newNode := new(LinkNode)
		newNode.Data = v
		newNode.Next = nil
		node.Next = newNode
		node = newNode
	}
}

//��ӡ��������Print()
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

//�ݹ��ӡ
func (node *LinkNode) Print1() {
	//�ݹ�ֹͣ������
	if node == nil {
		return
	}
	//�ڵݹ����ǰ���д�ӡ
	if node.Data != nil {
		fmt.Print(node.Data, " ")
	}
	//���еݹ����
	node.Next.Print1()

}

//��ȡ��������ĳ���Length
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

//�����б����--ͷ�巨insertByHead(Data interface{})
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

//�����б����--β�巨insertByTail(Data interface)
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

//�����������--��λ�ò���insertByindex()
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

//ɾ��������--��λ��ɾ�� DeleteByindex()
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

//ɾ��������--������ɾ�� DeleteByData()
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
package main

import (
	"fmt"
)

type QueueNode struct {
	Data interface{}
	Next *QueueNode
}

// 创建链式队列
func (queue *QueueNode) Create(Data ...interface{}) {
	if queue == nil || Data == nil {
		return
	}
	if len(Data) == 0 {
		return
	}
	head := queue
	// 创建 链式队列
	for _, v := range Data {
		newNode := new(QueueNode)
		newNode.Data = v
		newNode.Next = nil

		queue.Next = newNode
		queue = queue.Next
	}
	queue = head
}
// 打印链式队列
func (queue *QueueNode) Print() {
	if queue == nil {
		return
	}
	for queue.Next != nil {
		queue = queue.Next
		if queue.Data != nil {
			fmt.Print(queue.Data, " ")
		}
	}
	fmt.Println()
}

// 获取队列长度
func (queue *QueueNode) Length() int {
	if queue == nil {
		return -1
	}
	i := 0
	for queue.Next != nil {
		i++
		queue = queue.Next
	}
	return i
}

// 入队 —— 尾插
func (queue *QueueNode) Push(Data interface{}) {
	if queue == nil || Data == nil {
		return
	}
	for queue.Next != nil {
		queue = queue.Next
	}
	// 创建新结点，尾插
	newNode := new(QueueNode)
	newNode.Data = Data
	newNode.Next = nil

	queue.Next = newNode
}

// 出队 —— 删第一个结点
func (queue *QueueNode) Pop() {
	if queue == nil {
		return
	}
	queue.Next = queue.Next.Next
}

func main() {
	queue := new(QueueNode)
	queue.Create(1,2,3,4,5,6)

	//fmt.Println(queue)
	queue.Print()

	// 获取长度
	count := queue.Length()
	fmt.Println("长度：", count)

	// 入队
	queue.Push(7)
	queue.Print()

	// 出队
	queue.Pop()
	queue.Print()

}

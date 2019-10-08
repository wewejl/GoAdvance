package main

import (
	"fmt"
	"math/rand"
	"time"
)

//初始化堆
func HeapSort(arr []int) {
	// 将二叉树调整成最大堆存储，保证各个子树，（3个数中的）最大的值在“根”位置。
	length := len(arr)							// 9
	for i := length/2 - 1; i >= 0; i-- {		// 3 2 1 0
		CreateMaxHeap(arr, i, length-1)			// (3,8) (2,8) (1,8) (0,8)
	}
	// 上for循环，不从树上摘下结点

	// 此循环开始，根结点保存整个数组的最大值
	for i := length - 1; i > 0; i-- {			// 8 7 6 5 4 3 2 1
		// 如果只剩根结点和根的左子结点，排序结束。
		if i == 1 && arr[0] <= arr[i] {
			break
		}
		// 将根结点 和 最后一个叶子结点交换
		arr[0], arr[i] = arr[i], arr[0]

		// 将 交换后的 最后一个结点摘下（已经确认是最大，无需再比较）
		CreateMaxHeap(arr, 0, i-1)
	}
}

//获取堆中最大值  放在根结点  (3, 8) (2, 8) ...
func CreateMaxHeap(arr []int, startNode int, maxNode int) {

	// 存放较大值的下标
	var max int

	// 定义左子结点下标 和 右子结点下标
	lChild := startNode*2 + 1
	rChild := lChild + 1

	// 左子结点下标超出最大下标 跳出递归
	if lChild >= maxNode {
		return
	}
	// 左右比较  找到最大值
	if rChild <= maxNode && arr[rChild] > arr[lChild] {
		max = rChild
	} else {
		max = lChild
	}

	// 和跟结点比较
	if arr[max] <= arr[startNode] {
		return
	}

	// 交换数据
	arr[startNode], arr[max] = arr[max], arr[startNode]

	// 递归进行下次比较
	CreateMaxHeap(arr, max, maxNode)
}

func main()  {
	arr := make([]int, 0)
	rand.Seed(time.Now().UnixNano()) // 播种子
	// 打造 十万个 在 0-999 之间的 数据集
	for i := 0; i < 1000; i++ {
		arr = append(arr, rand.Intn(50))
	}

	HeapSort(arr)
	fmt.Println(arr)
}


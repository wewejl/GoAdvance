package main

import (
	"fmt"
	"math/rand"
	"time"
)

func QuickSort(arr []int, start, end int) {
	if start < end {
		i, j, base := start, end, arr[start]
		for i < j {
			for i < j && arr[j] >= base {
				j--
			}
			arr[j],arr[i]=arr[i],arr[j]
			for i < j && arr[i] <= base {
				i++
			}
			arr[j],arr[i]=arr[i],arr[j]
			//fmt.Println("时间一",i)
		}
		//这里要进行递归调用
		QuickSort(arr,start,i-1)
		QuickSort(arr,i+1,end)
	}
}

func main() {
	//进行验证
	arr := []int{}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 100; i++ {
		il := rand.Intn(100)
		arr = append(arr, il)
	}
	QuickSort(arr,0,len(arr)-1)
	fmt.Println(arr)
}

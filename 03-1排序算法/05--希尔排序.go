package main

import (
	"fmt"
	"math/rand"
	"time"
)

func ShellSort(arr []int) {
	//这个是增量
	for inc:=len(arr)/2;inc>=1;inc=inc/2 {
	//inc:=50
		for i:=0;i<len(arr)-inc;i++ {
				//发现小于 把数据加入  把前面的数据进行排序
				for j:=i;j>=0;j-=inc {
					fmt.Println("j:=",j)
					//前面都是有序的进行排序
					if arr[j]>arr[j+inc] {
						arr[j],arr[j+inc]=arr[j+inc],arr[j]
					}else {
						break
					}
				}
		}
	}
}
func main() {
	arr := []int{}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 100; i++ {
		il := rand.Intn(100)
		arr = append(arr, il)
	}
	ShellSort(arr)
	fmt.Println(arr)
}

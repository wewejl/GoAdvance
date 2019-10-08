package main

import "fmt"

func InsertSort(arr []int) {
	//外层控制行
	for i := 1; i < len(arr); i++ {
		//把 无序组的第一个和有序组的最后一个进行比较  如果大于 就把无序组的第一个变成有序组的最后一个 然后排序
		if arr[i-1]>arr[i] {
			//就把无序组的第一个变成有序组的最后一个
			for j := i; j > 0; j-- {
				//排序
				if arr[j-1]>arr[j] {
					arr[j-1],arr[j]=arr[j],arr[j-1]
				}else {
					break
				}
			}
		}
	}
}
func main() {
	arr := []int{6, 8, 3, 7, 5, 1, 4, 2, 9, 10}
	InsertSort(arr)
	fmt.Println(arr)
}

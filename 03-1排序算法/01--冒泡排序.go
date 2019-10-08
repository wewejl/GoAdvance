package main

import "fmt"

//冒泡排序实现
func BubbleSort(arr []int)  {
	//这个是计算要执行多少次
	start:=0
	//外层控制行
	for i:=0;i<len(arr)-1;i++ {
		//内层控制列
		for j:=0;j<len(arr)-1-i;j++ {
			start++
			//相邻进行比较
			if arr[j]>arr[j+1] {
				arr[j],arr[j+1]=arr[j+1],arr[j]
			}
		}
	}
	fmt.Println(start)
}

//冒泡排序优化
func BubbleSort2(arr []int)  {
	//这个是计算要执行多少次
	start:=0
	//定义一个全局变量
	fig :=false
	//外层控制行
	for i:=0;i<len(arr)-1;i++ {
		//内层控制列
		for j:=0;j<len(arr)-1-i;j++ {
			start++
			//相邻进行比较
			if arr[j]>arr[j+1] {
				arr[j],arr[j+1]=arr[j+1],arr[j]
				fig=true
			}
		}
		if !fig {
			break
		}else {
			fig=false
		}
	}
	fmt.Println("执行次数:",start)
}
func main() {
	arr:=[]int{0,1,2,3,4,5,6,8,7,9}
	BubbleSort(arr)
	fmt.Println(arr)
}

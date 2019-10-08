package main


import "fmt"

func SelectSort(arr []int) {
	//fmt.Println("len =",len(arr))
	for i := 0; i < len(arr)-1; i++ {
		//fmt.Println(i)
		index := 0
		for j := 1; j < len(arr)-i; j++ {
			//fmt.Println("j =",j)
			if arr[j] > arr[index] {
				index = j
			}
		}
		arr[index], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[index]
	}
}

func main() {
	arr := []int{1, 2, 3, 6, 5, 4, 7, 8, 9}
	SelectSort(arr)
	fmt.Println(arr)
}

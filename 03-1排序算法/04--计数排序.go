package main

import (
	"fmt"
	"math/rand"
	"time"
)

func CountingSort(arr []int) []int {
	MapSort :=make(map[int]int)
	for _,v:=range arr{
		MapSort[v]+=1
	}
	var newarr []int
	for i:=0;i<1000;i++ {
		for j:=0;j<MapSort[i];j++ {
			newarr=append(newarr,i)
		}
	}
	return newarr

}

func main() {
	arr:=[]int{}
	rand.Seed(time.Now().UnixNano())
	for i:=0;i<100000;i++ {
		il:=rand.Intn(999)
		arr=append(arr,il)
	}
	arr=CountingSort(arr)

	fmt.Println(arr)
}
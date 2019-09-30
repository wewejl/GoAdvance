package main

/*
#include <stdlib.h>
*/
import "C"
import (
	"unsafe"
	"fmt"
)

const Tag = 8

type Slice struct {
	Data unsafe.Pointer
	Len  int
	Cap  int
}

//创建切片数据结构Create
func (s *Slice) Create(l int, cp int, data ...int) {
	if s == nil {
		return
	}
	//判断错误
	if l < 0 || cp < 0 || l > cp || len(data) > cp {
		return
	}
	//拿到数据地址
	s.Data = C.malloc(C.size_t(cp) * Tag)
	//把指针的转换成数值类型进行添加
	p := uintptr(s.Data)
	for _, v := range data {
		*(*int)(unsafe.Pointer(p)) = v
		p += Tag
	}
	s.Len = l
	s.Cap = cp

}

//打印切片
func (s *Slice) Print() {
	//容错
	if s == nil || s.Len == 0 {
		return
	}
	//把地址转换成数据  然后地址就可以进行添加数了
	SliceData := uintptr(s.Data)
	//循环把数据读出来
	for i := 0; i < s.Len; i++ {
		fmt.Print(*(* int)(unsafe.Pointer(SliceData)), " ")
		SliceData += Tag
	}
	fmt.Println()
}

//追加元素Append(...数据)
func (s *Slice) Append(Data ...int) {
	//容错
	if s == nil || len(Data) == 0 {
		return
	}
	//如果数据添加然后总长度超过容器 就要进行扩容
	for s.Len+len(Data) > s.Cap {
		//如果超过了就进行2倍扩容
		s.Cap = s.Cap * 2
	}

	//把地址转换程数据,然后地址就可以进行偏移 (数据的加和减)
	SliceData := uintptr(s.Data)
	//把地址偏移到最后面
	SliceData += uintptr(s.Len) * Tag
	//循环把数据添加 上去
	for _, v := range Data {
		*(* int)(unsafe.Pointer(SliceData)) = v
		SliceData += Tag
	}
	s.Len += len(Data)
}

//获取切片元素GetData(下标)int
func (s *Slice) GetData(Data int) int {
	//容错
	if s == nil || Data < 0 || Data >= s.Len {
		return -1
	}

	//将地址转化成数据
	SliceData := uintptr(s.Data)
	//找到下标的地址
	spSliceData := SliceData + uintptr(Data)*Tag
	//把数据话转化成地址 并返回去
	return *(* int)(unsafe.Pointer(spSliceData))

}

//查找切片元素Search(数据)下标
func (s *Slice) Search(Data int) int {
	//容错
	if s == nil {
		return -1
	}
	//把地址转化成数据
	SliceData := uintptr(s.Data)
	//循环把数据遍历
	for i := 0; i < s.Len; i++ {
		if Data == *(*int)(unsafe.Pointer(SliceData)) {
			return i
		}
		SliceData += Tag
	}
	return -1
}

//插入元素到切片中insert(数据,下标)
func (s *Slice) insert(Data, index int) {
	//容错
	if s == nil || Data == 0 || index > s.Len || index < 0 {
		return
	}
	if s.Len+1 > s.Cap {
		s.Cap = s.Cap * 2
	}
	//进行数据保存
	// 地址 -- 数值
	p := uintptr(s.Data)

	// 保存 index 对应的内存地址
	p += uintptr(index) * Tag

	// 定义变量 tmp , 用来移动 index 后序元素.
	tmp := uintptr(s.Data) + uintptr(s.Len)*Tag // 将 tmp 指向切片尾元素的后一个位置.

	// 从后向前循环,前一个元素,给后一个元素赋值.
	for i := s.Len; i > index; i-- {
		// 前一个元素,给后一个元素赋值
		*(*int)(unsafe.Pointer(tmp)) = *(*int)(unsafe.Pointer(tmp - 8))
		tmp -= Tag // 指针前移
	}

	// 在index 位置写入,新插入数据值
	*(*int)(unsafe.Pointer(p)) = Data

	// 更新有效长度
	s.Len++
}

//删除切片中的元素Delete(下标)
func (s *Slice) Delete(index int) {
	//容错
	if s == nil || index > s.Len || index < 0 {
		return
	}
	//进行地址的数据转化
	SliceData := uintptr(s.Data)
	//SliceData += Tag * uintptr(index)
	//进行
	ptrindex := s.Len - index - 1
	for i := 0; i < ptrindex; i++ {
		if i==ptrindex {
			*(*int)(unsafe.Pointer(SliceData + Tag))=0
		}
		*(*int)(unsafe.Pointer(SliceData)) = *(*int)(unsafe.Pointer(SliceData + Tag))
		SliceData+=Tag
	}
	s.Len--
}

//销毁切片
func (s *Slice)Destroy()  {
	s.Len=0
	s.Data=nil
	s.Cap=0
}

func main() {
	var s Slice
	s.Create(5, 5, 1, 2, 3, 4, 5)
	//s.Append(123, 123, 12, 2, 2, 2, 2, 2)
	//s.Print()
	//fmt.Println(s.GetData(11))

	s.Delete(0)
	s.Print()
	s.Destroy()

	s.Print()
}

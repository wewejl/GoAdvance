package main

import "fmt"

// 定义链式栈结点结构体
type StackNode struct {
	Data interface{}
	Next *StackNode
}

// 创建链式栈
func CreateStack(Data ...interface{}) *StackNode {
	//容错
	if len(Data) == 0 {
		return nil
	}
	// 创建头结点
	var stack *StackNode
	// 创建结点,保存 新结点的下一个结点
	var NextStack *StackNode
	// 循环获取元素,创建结点,组成链栈
	for _,v:=range Data{
		newNode:=new(StackNode)
		newNode.Data=v
		newNode.Next=nil
		stack=newNode
		stack.Next=NextStack
		NextStack=stack
	}
	return stack
}


//打印链栈表
func PrintStack(s *StackNode) {
	for s!=nil  {
		fmt.Print(s.Data," ")
		s=s.Next
	}
}

//获取链栈的长度
func (s *StackNode)LengthStack()int  {
	if s==nil {
		return 0
	}
	s=s.Next
	sH:=s.LengthStack()
	sH++
	return sH
}

//压栈（入栈）-- 头插
func (s *StackNode)Push(Data interface{})  {
	newStackNode:=new(StackNode)
	newStackNode.Data=s.Data
	newStackNode.Next=s.Next
	s.Data=Data
	s.Next=newStackNode
}

//弹栈（出栈）-- 头删
func (s *StackNode)Pop()  {
	s.Data=s.Next.Data
	s.Next=s.Next.Next
}
func main() {
	stack:=CreateStack(1,2,3,4,5,6,7,8,9,10)
	//fmt.Println(stack)
	//PrintStack(stack)
	//fmt.Println("长度:",stack.LengthStack())
	stack.Push(888)

	stack.Pop()
	PrintStack(stack)
}
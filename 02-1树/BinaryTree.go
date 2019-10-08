package main

import (
	"fmt"
	"reflect"
)

type BinTreeNode struct {
	Data   interface{}
	LChild *BinTreeNode
	RChild *BinTreeNode
}

//创建二叉树 --- 根据树的结构
func (node *BinTreeNode) Create() {
	node.Data = 0
	node.LChild = nil
	node.RChild = nil
	node1 := BinTreeNode{1, nil, nil}
	node2 := BinTreeNode{2, nil, nil}
	node3 := BinTreeNode{3, nil, nil}
	node4 := BinTreeNode{4, nil, nil}
	node5 := BinTreeNode{5, nil, nil}
	node6 := BinTreeNode{6, nil, nil}
	node7 := BinTreeNode{7, nil, nil}
	node.LChild = &node1
	node.RChild = &node2
	node1.LChild = &node3
	node1.RChild = &node4
	node2.LChild = &node5
	node2.RChild = &node6
	node3.LChild = &node7
}

//先序遍历 -- DLR
func (node *BinTreeNode) PreOrder() {
	if node == nil {
		return
	}
	fmt.Print(node.Data, " ")
	node.LChild.PreOrder()
	node.RChild.PreOrder()
}

//中序遍历 --LDR
func (node *BinTreeNode) MidOrder() {
	//容错,递归出口
	if node == nil {
		return
	}
	node.LChild.MidOrder()
	fmt.Print(node.Data, " ")
	node.RChild.MidOrder()
}

//后序遍历 --LRD
func (node *BinTreeNode) PostOrder() {
	//容错,递归出口
	if node == nil {
		return
	}
	node.RChild.PostOrder()
	node.LChild.PostOrder()
	fmt.Print(node.Data, " ")
}

//获取二叉数据的高度(深度)
func (node *BinTreeNode) TreeHeight() int {
	//容错,递归出口
	if node == nil {
		return 0
	}
	//递归
	lh := node.LChild.TreeHeight()
	rh := node.RChild.TreeHeight()
	if lh > rh {
		lh++
		return lh
	} else {
		rh++
		return rh
	}
}

//获取叶子结点个数
//全局变量
// 统计二叉树的叶子结点个数方法1
var num = 0 // 统计叶子结点个数
func (node *BinTreeNode)LeafNum1() {
	if node == nil {
		return
	}
	//寻找叶子结点
	if node.LChild == nil && node.RChild == nil {
	num++
	}
	//左/右子数调用递归
	node.LChild.LeafNum1()
	node.RChild.LeafNum1()
}

//获取二叉数叶子结点子数方法2
func (node *BinTreeNode)LeafNum2(n *int)  {
	//容错,递归退出条件
	if node==nil {
		return
	}
	//寻找叶子结点
	if node.LChild == nil && node.RChild == nil {
		(*n)++
	}
	//左/右子数调用递归
	node.LChild.LeafNum2(n)
	node.RChild.LeafNum2(n)

}

//获取二叉树叶子结点子数方法3
func (node *BinTreeNode) LeafNum3() int {
	if node == nil {
		return 0
	}
	if node.RChild == nil && node.LChild == nil {
		return 1
	}
	lN := node.LChild.LeafNum3()
	rN := node.RChild.LeafNum3()
	//fmt.Println(node.Data, " LN =", lN, "rN =", rN)
	return lN + rN
}

// 查找二叉树上是否含有 xxx 数据
var bl = false		// 表示是否在树上找到数据
func (node *BinTreeNode)Search(Data interface{})  {
	if node==nil {
		return
	}
	if reflect.DeepEqual(node.Data,Data) {
		bl=true
		return
	}
	node.RChild.Search(Data)
	node.LChild.Search(Data)
}

//二叉树的翻转
func (node *BinTreeNode)Reverse()  {
	if node==nil {
		return
	}
	//利用go语言多重赋值
	node.LChild,node.RChild=node.RChild,node.LChild
	node.LChild.Reverse()
	node.RChild.Reverse()
}

//二叉树的拷贝
func (node *BinTreeNode)Copy() *BinTreeNode {
	if node==nil {
		return nil
	}
	leftChild:=node.LChild.Copy()
	rightChild:=node.RChild.Copy()

	//创建新结点
	newNode :=new(BinTreeNode)
	newNode.Data=node.Data
	newNode.LChild=leftChild
	newNode.RChild=rightChild

	return newNode
}

func main() {
	tree := new(BinTreeNode)
	tree.Create()
	//tree.PreOrder()
	//fmt.Println()
	tree.LeafNum1()
	fmt.Println("叶子数:=",num)
	//tree.PostOrder()
	//fmt.Println(tree.TreeHeight())

	//查找二叉树上是否含有 xxx 数据
	tree.Search(2)
	fmt.Println(bl)
	////验证二叉树是否翻转
	//tree.MidOrder()
	//fmt.Println()
	//tree.Reverse()
	//tree.MidOrder()

	// 测试二叉树拷贝
	tree.PreOrder()
	newTree := tree.Copy()
	newTree.Data = 999
	fmt.Println()
	newTree.PreOrder()
}

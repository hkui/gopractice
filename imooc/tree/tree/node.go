package tree

import "fmt"

type Node struct {
	Value int
	Left,Right *Node
}
func CreateNode(Value int) *Node{
	return &Node{Value:Value}  //返回了局部变量的地址
}
func (node Node)Print()  {
	fmt.Print(node.Value)
}
func (node *Node)SetValue(Value int)  {
	if(node ==nil){
		fmt.Println("setting Value=",Value," to nil")
		return
	}
	node.Value=Value
}
//中序遍历(左中右)
func (node *Node)Traverse()  {
	if(node ==nil){
		return
	}
	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}


package _2_tree

import "testing"

type Node struct {
	data  int
	left  *Node
	right *Node
}

func Build_Node() *Node {
	node4 := Node{8, nil, nil}
	node3 := Node{7, nil, nil}
	node2 := Node{6, &node4, nil}
	node1 := Node{5, &node2, &node3}
	return &node1
}

func Pre_Order(node *Node){
	if node != nil {
		println(node.data)
		Pre_Order(node.left)
		Pre_Order(node.right)
	}
}

func In_Order(node *Node) {
	if node != nil{
		In_Order(node.left)
		println(node.data)
		In_Order(node.right)
	}
}


func Post_Order(node *Node) {
	if node != nil{
		Post_Order(node.left)
		Post_Order(node.right)
		println(node.data)
	}
}


func Test_Tree(t *testing.T) {
	root := Build_Node()
	// 前序遍历
	Pre_Order(root)

	// 中序遍历
	In_Order(root)

	// 后序遍历
	Post_Order(root)
}

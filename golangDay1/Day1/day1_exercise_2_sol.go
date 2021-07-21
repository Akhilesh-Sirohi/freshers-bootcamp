package main

import "fmt"

type treeNode struct {
	data byte
	left *treeNode
	right *treeNode
}
func newNode(data byte) treeNode{
	return treeNode{data,nil,nil}
}
func print(Node treeNode){
	fmt.Println(Node)
}

func constructTree() treeNode{
	root:=newNode('+')
	temp1:=newNode('a')
	root.left=&temp1
	temp2:=newNode('-')
	root.right=&temp2
	temp3:=newNode('b')
	(*root.right).left=&temp3
	temp4:=newNode('c')
	(*root.right).right=&temp4
	return root;
}
func printCh(data byte){
	//to print character
	fmt.Printf("%q\n", data)
}
func preOrder(root *treeNode){
	//root->root.left->root.right
	if root==nil {
		return
	}
		printCh(root.data)
		preOrder(root.left)
		preOrder(root.right)

}
func postOrder(root *treeNode){
	//root.left->root.right->root
	if root!=nil {
		postOrder(root.left)
		postOrder(root.right)
		printCh(root.data)
	}

}
func main() {
	//a:='a'
	//fmt.Printf("%q\n", a)
	fmt.Println("Pre Order : ")
	root:=constructTree()
	preOrder(&root)
	fmt.Println("Post Order : ")
	postOrder(&root)
	//print(root)
	//print(*root.left)
}
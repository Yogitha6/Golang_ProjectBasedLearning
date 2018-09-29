package main

import (
	"fmt"
)

type bstnode struct {
	data  int
	left  *bstnode
	right *bstnode
}

func (n *bstnode) addNode(data int) {
	if n.data == 0 {
		n.data = data
		return
	}

	for {
		if data < n.data {
			if n.left == nil {
				n.left = &bstnode{data: data}
				//fmt.Println("L assigned data ", data)
				break
			} else {
				n = n.left
			}
		} else if data > n.data {
			if n.right == nil {
				n.right = &bstnode{data: data}
				//fmt.Println("R assigned data ", data)
				break
			} else {
				n = n.right
			}
		} else {
			break
		}
	}

}

//left, root, right
func (n *bstnode) printInOrder(root *bstnode) {
	if root != nil {
		n.printInOrder(root.left)
		fmt.Println(root.data)
		n.printInOrder(root.right)
	}
}

func (n *bstnode) deleteNode(data int) {
	var prev *bstnode
	for {
		if data < n.data {
			if n.left == nil {
				break
			} else {
				prev = n
				n = n.left
			}
		} else if data > n.data {
			if n.right == nil {
				break
			} else {
				prev = n
				n = n.right
			}
		} else if data == n.data {
			if n.left == nil && n.right != nil {
				if prev.left.data == n.data {
					prev.left = n.right
				} else {
					prev.right = n.right
				}
				break
			} else if n.left != nil && n.right == nil {
				if prev.left.data == n.data {
					prev.left = n.left
				} else {
					prev.right = n.left
				}
				break
			} else if n.left == nil && n.right == nil {
				if prev.left.data == n.data {
					prev.left = nil
				} else {
					prev.right = nil
				}
				break
			} else {
				//find largest in left most tree
				//under construction
				break
			}
		}
	}

}

func (n *bstnode) searchNode(data int) bool {
	if n.data == data {
		return true
	}

	for {
		if data < n.data {
			if n.left == nil {
				return false
			} else {
				n = n.left
			}
		} else if data > n.data {
			if n.right == nil {
				return false
			} else {
				n = n.right
			}
		} else if data == n.data {
			return true
		}
	}
}

func main() {
	fmt.Println("Hello, playground")
	var node bstnode
	root := &node
	root.addNode(6)
	root.addNode(4)
	root.addNode(3)
	root.addNode(8)
	root.addNode(7)
	root.addNode(10)
	root.printInOrder(root)
	fmt.Println(root.searchNode(9))
	root.deleteNode(4)
	root.printInOrder(root)
}

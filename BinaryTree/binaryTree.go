package BinaryTree

import "fmt"

type Node struct {
	key   int
	Left  *Node
	Right *Node
}

// Insert
func (n *Node) Insert(value int) {
	if n.key == 0 { //adding root
		n.key = value
		return
	}
	if n.key > value { //smaller
		if n.Left == nil {
			n.Left = &Node{key: value}
		} else {
			n.Left.Insert(value)
		}
	} else if n.key < value { //larger
		if n.Right == nil {
			n.Right = &Node{key: value}
		} else {
			n.Right.Insert(value)
		}
	} else if n.key == value {
		panic("Dublicate : The value already exists in the tree") // It depends on you, if you want to add then write  condition for that in Left or Right
	}
}

// Search
func (n *Node) Search(value int) bool {
	if n == nil {
		return false
	}
	if n.key > value {
		return n.Left.Search(value)
	} else if n.key < value {
		return n.Right.Search(value)
	}
	return true
}

func StartBinaryTree() {
	tree := &Node{}
	var value, sVal int
	ans := "y"
	sans := "y"
	for ans == "y" || ans == "Y" {
		fmt.Println("Enter a Value to Add to the Tree:")
		fmt.Scanf("%v", &value)
		tree.Insert(value)
		fmt.Println(tree)
		fmt.Println("Would you like to add another value to the Tree?(y/n)")
		fmt.Scanf("%v", &ans)
	}

	for sans == "y" || sans == "Y" {
		fmt.Println("Enter a Value to Search in the Tree:")
		fmt.Scanf("%v", &sVal)
		fmt.Println(tree.Search(sVal))
		fmt.Println("Would you like to search another value in the Tree?(y/n)")
		fmt.Scanf("%v", &sans)
	}
}

// Single Linked List
package LinkedList

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	Head   *Node
	length int
}

// Insert
func (l *LinkedList) Insert(node *Node) {
	// dummy := l.Head
	// l.Head = node
	// l.Head.next = dummy
	if l.Head == nil {
		l.Head = node
	} else {
		n := l.Head
		for n.next != nil {
			n = n.next
		}
		n.next = node
	}
	l.length++
}

func (l LinkedList) PrintLinkedList() {
	node := l.Head
	for l.length != 0 {
		fmt.Printf("-> %v ", node.data)
		node = node.next
		l.length--
	}
	fmt.Println()
}

func (l *LinkedList) DleteValue(value int) {
	if l.length == 0 {
		return
	}
	if l.Head.data == value {
		l.Head = l.Head.next
		l.length--
		return
	}
	prvToDelete := l.Head
	for prvToDelete.next.data != value {
		if prvToDelete.next.next == nil {
			return
		}
		prvToDelete = prvToDelete.next
	}
	prvToDelete.next = prvToDelete.next.next
	l.length--
}

// Convert single-linked-list to circular-single-linked-list
func (l *LinkedList) ConvertSingleToCircularLinkedList() {
	node := l.Head
	for node.next != nil {
		node = node.next
	}
	node.next = l.Head
}

func (l LinkedList) PrintCircularSingleLinkedList() {
	node := l.Head
	for {
		if l.Head == node.next {
			fmt.Printf("-> %v ", node.data)
			break
		}
		fmt.Printf("-> %v ", node.data)
		node = node.next
	}
	fmt.Println()
}

func (l *LinkedList) DleteValueFromCircularSingleLinkedList(value int) {
	if l.length == 0 {
		return
	}
	if l.Head.data == value {
		oldhead := l.Head
		l.Head = l.Head.next
		node := l.Head
		for node.next != oldhead {
			node = node.next
		}
		node.next = l.Head
		l.length--
		return
	}
	prvToDelete := l.Head
	for prvToDelete.next.data != value {
		if prvToDelete.next.next == nil {
			return
		}
		prvToDelete = prvToDelete.next
	}
	prvToDelete.next = prvToDelete.next.next
	if prvToDelete.next.next == nil {
		prvToDelete.next = l.Head
	}
	l.length--

}

func StartLinkedList() {
	linkedList := &LinkedList{}
	var value int
	ans := "y"
	for ans == "y" || ans == "Y" {
		fmt.Println("Enter a Value to Add to the LinkedList:")
		fmt.Scanf("%v", &value)
		node := &Node{data: value}
		linkedList.Insert(node)
		fmt.Println("Would you like to add another value to the LinkedList?(y/n)")
		fmt.Scanf("%v", &ans)
	}
	linkedList.PrintLinkedList()
	linkedList.DleteValue(10) //paas the value tht you wnt remove from the list
	linkedList.PrintLinkedList()
	fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@")
	linkedList.ConvertSingleToCircularLinkedList()
	linkedList.PrintCircularSingleLinkedList()
	linkedList.DleteValueFromCircularSingleLinkedList(20)
	linkedList.PrintCircularSingleLinkedList()
}

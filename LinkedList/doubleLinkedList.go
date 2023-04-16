// Double Linked List
package LinkedList

import "fmt"

type ListNode struct {
	data int
	next *ListNode
	prev *ListNode
}

type DoubleLinkedList struct {
	Head   *ListNode
	Tail   *ListNode
	length int
}

// Insert
func (l *DoubleLinkedList) InsertToDoubleLinkedList(node *ListNode) {
	if l.Head == nil {
		l.Head = node
		l.Tail = node
	} else {
		n := l.Head
		for n.next != nil {
			n = n.next
		}
		node.prev = n
		n.next = node
		l.Tail = node
	}
	l.length++
}

func (l DoubleLinkedList) PrintDoubleLinkedList() {
	node := l.Head
	for node != nil {
		fmt.Printf("-> %v ", node.data)
		node = node.next
	}
	fmt.Println()
}

func (l DoubleLinkedList) PrintDoubleLinkedListInReverse() {
	node := l.Tail
	for node != nil {
		fmt.Printf("-> %v ", node.data)
		node = node.prev
	}
	fmt.Println()
}

// Convert double-linked-list to circular-single-linked-list
func (l *DoubleLinkedList) ConvertDoubleToCircularLinkedList() {
	l.Head.prev = l.Tail
	l.Tail.next = l.Head
}

func (l DoubleLinkedList) PrintCircularDoubleLinkedList() {
	node := l.Head
	for {
		if node.next == l.Head {
			fmt.Printf("-> %v ", node.data)
			break
		}
		fmt.Printf("-> %v ", node.data)
		node = node.next
	}
	fmt.Println()
}

func (l DoubleLinkedList) PrintCircularDoubleLinkedListInReverse() {
	node := l.Tail
	for {
		if node.prev == l.Tail {
			fmt.Printf("-> %v ", node.data)
			break
		}
		fmt.Printf("-> %v ", node.data)
		node = node.prev
	}
	fmt.Println()
}

func (l *DoubleLinkedList) DleteValueFromDoubleLinkedList(value int) {
	if l.length == 0 {
		return
	}
	if l.Head.data == value {
		l.Head = l.Head.next
		l.Head.prev = nil
		l.length--
		return
	}
	if l.Tail.data == value {
		l.Tail = l.Tail.prev
		l.Tail.next = nil
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
	prvToDelete.next.next.prev = prvToDelete //setting previous data first
	prvToDelete.next = prvToDelete.next.next //setting next data later
	l.length--
}

func (l *DoubleLinkedList) DleteValueFromCircularDoubleLinkedList(value int) {
	if l.length == 0 {
		return
	}
	if l.Head.data == value {
		l.Head = l.Head.next
		l.Head.prev = l.Tail
		l.Tail.next = l.Head
		l.length--
		return
	}
	if l.Tail.data == value {
		l.Tail = l.Tail.prev
		l.Tail.next = l.Head
		l.Head.prev = l.Tail
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
	prvToDelete.next.next.prev = prvToDelete //setting previous data first
	prvToDelete.next = prvToDelete.next.next //setting next data later
	l.length--
}

func StartDoubleLinkedList() {
	linkedList := &DoubleLinkedList{}
	var value int
	ans := "y"
	for ans == "y" || ans == "Y" {
		fmt.Println("Enter a Value to Add to the LinkedList:")
		fmt.Scanf("%v", &value)
		node := &ListNode{data: value}
		linkedList.InsertToDoubleLinkedList(node)
		fmt.Println("Would you like to add another value to the LinkedList?(y/n)")
		fmt.Scanf("%v", &ans)
	}
	linkedList.PrintDoubleLinkedList()
	linkedList.PrintDoubleLinkedListInReverse()
	linkedList.DleteValueFromDoubleLinkedList(10)
	linkedList.PrintDoubleLinkedList()
	linkedList.PrintDoubleLinkedListInReverse()
	fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@")
	linkedList.ConvertDoubleToCircularLinkedList()
	linkedList.PrintCircularDoubleLinkedList()
	linkedList.PrintCircularDoubleLinkedListInReverse()
	linkedList.DleteValueFromCircularDoubleLinkedList(20)
	linkedList.PrintCircularDoubleLinkedList()
	linkedList.PrintCircularDoubleLinkedListInReverse()
}

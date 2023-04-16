package HashTable

import "fmt"

type HashTable struct {
	table  []*BucketList
	hIndex int
}

type BucketList struct {
	Head *BucketNode
}

type BucketNode struct {
	key  string
	next *BucketNode
}

func Hash(w string) int {
	sum := 0
	for _, ch := range w {
		sum += int(ch)
	}
	return sum % len(w)
}

func (h *HashTable) Insert(k string) {
	index := Hash(k)
	if len(h.table) == 0 {
		h.table = append(h.table, &BucketList{})
		h.hIndex = 0
	}
	if h.hIndex < index {
		for i := h.hIndex + 1; i <= index; i++ {
			h.table = append(h.table, &BucketList{})
		}
		h.hIndex = index
	}
	h.table[index].Insert(k)
}

func (h *HashTable) Search(k string) bool {
	index := Hash(k)
	return h.table[index].Search(k)
}

func (h *HashTable) Delete(k string) {
	index := Hash(k)
	h.table[index].Delete(k)
}

func (b *BucketList) Insert(k string) {
	if !b.Search(k) {
		node := &BucketNode{key: k}
		node.next = b.Head
		b.Head = node
	} else {
		panic("Already exists in the Bucket")
	}
}

func (b *BucketList) Search(k string) bool {
	node := b.Head
	for node != nil {
		if node.key == k {
			return true
		}
		node = node.next
	}
	return false
}

func (b *BucketList) Delete(k string) {
	if b.Head.key == k {
		b.Head = b.Head.next
		return
	}
	prvToDelete := b.Head
	if prvToDelete.next != nil {
		if prvToDelete.next.key == k {
			prvToDelete.next = prvToDelete.next.next
		}
		prvToDelete = prvToDelete.next
	}
}

func Init() *HashTable {
	htable := &HashTable{}
	for i := 0; i < len(htable.table); i++ {
		htable.table[i] = &BucketList{}
	}
	return htable
}

func StartHashTable() {
	hashTable := Init()
	var key, sKey, dKey string
	ans := "y"
	for ans == "y" || ans == "Y" {
		fmt.Println("Enter a Value to Add to the HashTable:")
		fmt.Scanf("%v", &key)
		hashTable.Insert(key)
		fmt.Println("Would you like to add another value to the HashTable?(y/n)")
		fmt.Scanf("%v", &ans)
	}
	fmt.Println("Enter the value to search in the Key:")
	fmt.Scanf("%v", &sKey)
	fmt.Println(hashTable.Search(sKey))
	fmt.Println("Enter the value to delete in the Key:")
	fmt.Scanf("%v", &dKey)
	hashTable.Delete(dKey)
	fmt.Println(hashTable.Search(dKey))
}

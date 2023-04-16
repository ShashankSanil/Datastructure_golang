package Heap

import "fmt"

type Heap struct {
	HeapTree []int
	HeapType string // MaxHeap or MinHeap
}

// In Heap left Child will have odd index and right Child will have even index
func (H *Heap) Parent(childIndex int) int {
	return (childIndex - 1) / 2
}

func (H *Heap) Right(parentIndex int) int {
	return (parentIndex * 2) + 2
}

func (H *Heap) Left(parentIndex int) int {
	return (parentIndex * 2) + 1
}

// Swapping Elements
func (H *Heap) Swap(index1, index2 int) {
	H.HeapTree[index1], H.HeapTree[index2] = H.HeapTree[index2], H.HeapTree[index1]
}

func (H *Heap) maxHeapifyUp() {
	index := len(H.HeapTree) - 1
	for H.HeapTree[H.Parent(index)] < H.HeapTree[index] {
		H.Swap(H.Parent(index), index)
		index = H.Parent(index)
	}
}

func (H *Heap) minHeapifyUp() {
	index := len(H.HeapTree) - 1
	for H.HeapTree[H.Parent(index)] > H.HeapTree[index] {
		H.Swap(H.Parent(index), index)
		index = H.Parent(index)
	}
}

func (H *Heap) maxHeapifyDown() {
	index := 0
	lastIndex := len(H.HeapTree) - 1
	leftIndex, rightIndex := H.Left(index), H.Right(index)
	cmpIndex := 0
	for leftIndex <= lastIndex {
		if leftIndex == lastIndex {
			cmpIndex = leftIndex
		} else if H.HeapTree[leftIndex] > H.HeapTree[rightIndex] {
			cmpIndex = leftIndex
		} else {
			cmpIndex = rightIndex
		}

		if H.HeapTree[index] < H.HeapTree[cmpIndex] {
			H.Swap(index, cmpIndex)
			index = cmpIndex
			leftIndex, rightIndex = H.Left(index), H.Right(index)
		} else {
			return
		}
	}
}

func (H *Heap) minHeapifyDown() {
	index := 0
	lastIndex := len(H.HeapTree) - 1
	leftIndex, rightIndex := H.Left(index), H.Right(index)
	cmpIndex := 0
	for leftIndex <= lastIndex {
		if leftIndex == lastIndex {
			cmpIndex = leftIndex
		} else if H.HeapTree[leftIndex] < H.HeapTree[rightIndex] {
			cmpIndex = leftIndex
		} else {
			cmpIndex = rightIndex
		}

		if H.HeapTree[index] > H.HeapTree[cmpIndex] {
			H.Swap(index, cmpIndex)
			index = cmpIndex
			leftIndex, rightIndex = H.Left(index), H.Right(index)
		} else {
			return
		}
	}
}

// Insert --> O(log n)
func (H *Heap) Insert(key int) {
	fmt.Println(H.HeapType)
	H.HeapTree = append(H.HeapTree, key)
	if H.HeapType == "MinHeap" {
		H.minHeapifyUp()
	} else {
		H.maxHeapifyUp()
	}
}

// Deletion of priority --> O(log n)
func (H *Heap) Extract() {
	if len(H.HeapTree) == 0 {
		panic("The heap is empty!!!")
	}
	lastIndex := len(H.HeapTree) - 1
	H.HeapTree[0] = H.HeapTree[lastIndex]
	H.HeapTree = H.HeapTree[:lastIndex]

	if H.HeapType == "MinHeap" {
		H.minHeapifyDown()
	} else {
		H.maxHeapifyDown()
	}
}

// Getting Peek (depends on HeapType so i wrote 2 function that will return min or max if case will anything)
// GetMin
func (H *Heap) getMin() int {
	H.minHeapifyUp()
	return H.HeapTree[0]
}

// GetMax
func (H *Heap) getMax() int {
	H.maxHeapifyUp()
	return H.HeapTree[0]
}

func StratHeap() {
	heap := &Heap{}
	var value int
	fmt.Println("Enter a HeapType:\n1.MaxHeap\n2.MinHeap")
	fmt.Scanf("%v", &heap.HeapType)
	ans := "y"
	for ans == "y" || ans == "Y" {
		fmt.Println("Enter a Value to Add to the Heap:")
		fmt.Scanf("%v", &value)
		heap.Insert(value)
		fmt.Println(heap.HeapTree)
		fmt.Println("Would you like to add another value to the heap?(y/n)")
		fmt.Scanf("%v", &ans)
	}
	fmt.Println("GetMax : ", heap.getMax())
	fmt.Println("GetMin : ", heap.getMin())
	fmt.Println("**********************************************************************")
	for i := 0; i < 3; i++ {
		heap.Extract()
		fmt.Println(heap.HeapTree)
	}

}

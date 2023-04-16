package Tries

import "fmt"

type Node struct {
	children [26]*Node
	isEnd    bool
}

type Tries struct {
	root *Node
}

func (t *Tries) Insert(w string) {
	wordLength := len(w)
	currentIndex := t.root
	for i := 0; i < wordLength; i++ {
		index := w[i] - 'a'
		if currentIndex.children[index] == nil {
			currentIndex.children[index] = &Node{}
		}
		currentIndex = currentIndex.children[index]
	}
	currentIndex.isEnd = true
}

func (t *Tries) Search(w string) bool {
	wordLength := len(w)
	currentIndex := t.root
	for i := 0; i < wordLength; i++ {
		index := w[i] - 'a'
		if currentIndex.children[index] == nil {
			return false
		}
		currentIndex = currentIndex.children[index]
	}
	return currentIndex.isEnd
}

func StartTries() {
	tries := &Tries{root: &Node{}}
	var value string
	ans := "y"
	for ans == "y" || ans == "Y" {
		fmt.Println("Enter a word to Add to the Tries:")
		fmt.Scanf("%v", &value)
		tries.Insert(value)
		fmt.Println("Would you like to add another word to the Tries?(y/n)")
		fmt.Scanf("%v", &ans)
	}
	var sWord string
	fmt.Println("Enter a word to Search in the Tries:")
	fmt.Scanf("%v", &sWord)
	fmt.Println(tries.Search(sWord))
}

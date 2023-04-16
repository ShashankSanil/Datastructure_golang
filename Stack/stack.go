//							||          ||  |=========    =======
//							||			||  ||          ||       ||
//							||			||  ||          ||       ||
//							||       	||  |========   ||       ||
//							||			||  ||          ||       ||
//							||			||  ||          ||       ||
//							|=========	||  ||            =======
//
//
//
//									Push     -------     Pop
//								*  *  *	 * >|		| *  *  *  *             
//					 ------- *				|		|             * > -------
//					|		|				 -------	|			 |		 |
//			    	|		|				|		|   |    		 |		 |
//					 -------				|		|   |			  -------
//											 -------    |
//											|		|   |
//											|		|  \/
//											 -------
//											|		|
//											|		|
//											 -------




package Stack

import "fmt"

type Stack struct {
	stackList []int
}

func (s *Stack) Push(key int) {
	s.stackList = append(s.stackList, key)
}

func (s *Stack) Pop() int {
	lastIndex := len(s.stackList) - 1
	extractVal := s.stackList[lastIndex]
	s.stackList = s.stackList[:lastIndex]
	return extractVal
}

func StartStack() {
	stack := &Stack{}
	var value int
	ans := "y"
	for ans == "y" || ans == "Y" {
		fmt.Println("Enter a Value to Add to the Stack:")
		fmt.Scanf("%v", &value)
		stack.Push(value)
		fmt.Println(stack.stackList)
		fmt.Println("Would you like to add another value to the Stack?(y/n)")
		fmt.Scanf("%v", &ans)
	}

	for i := 0; i < 3; i++ {
		stack.Pop()
		fmt.Println(stack.stackList)
	}
}

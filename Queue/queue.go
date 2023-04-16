//							|=========  ||  |=========    =======
//							||			||  ||          ||       ||
//							||			||  ||          ||       ||
//							|========	||  |========   ||       ||
//							||			||  ||          ||       ||
//							||			||  ||          ||       ||
//							||			||  ||            =======
//
//
//								 								  	 ____
//																  	|    |  Dequeue
//																  	|____|
//																    /\
//								 								   *
//								  ------------------->            *                         
//								 ________________________________*
//								|	 |	  |	   |	 |	   |	 |                  
//								|____|____|____|_____|_____|_____|
//								/\		 								
//								*
//							  *
//		   					*
//						 ____
//		 	 		    |    |  Enqueue
//						|____|


package Queue

import "fmt"

type Queue struct {
	queueList []int
}

func (q *Queue) Enqueue(key int) {
	q.queueList = append(q.queueList, key)
}

func (q *Queue) Dequeue() int {
	extractVal := q.queueList[0]
	q.queueList = q.queueList[1:]
	return extractVal
}

func StartQueue() {
	queue := &Queue{}
	var value int
	ans := "y"
	for ans == "y" || ans == "Y" {
		fmt.Println("Enter a Value to Add to the Queue:")
		fmt.Scanf("%v", &value)
		queue.Enqueue(value)
		fmt.Println(queue.queueList)
		fmt.Println("Would you like to add another value to the Queue?(y/n)")
		fmt.Scanf("%v", &ans)
	}

	for i := 0; i < 3; i++ {
		queue.Dequeue()
		fmt.Println(queue.queueList)
	}
}

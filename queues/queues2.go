package queues

import (
	"fmt"
	"time"
	// "errors"
)

func LaunchQueue(queue *Node){
	current := queue
	for i:=0;current != nil;i++{
		fmt.Printf("job running: %s | Node: %d\n", current.run.name, current.nb)
		time.Sleep(1 * time.Second)

		ShowNodes(queue)
		err := DelNode(0,&queue)
		if err != nil {
			fmt.Printf("Error: %s", err)
			return 
		}
		if current.next != nil {
			fmt.Printf("Done, next one : %s\n", current.next.run.name)
		} else {
			fmt.Println("Done")
		}
		// fmt.Println(current.nb)
		current = current.next
	}
	ShowNodes(queue)
}

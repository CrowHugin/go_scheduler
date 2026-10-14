package main

import (
	"fmt"
	// "errors"
	"go_scheduler/queues"
	"go_scheduler/simu"
)

func main() {
	testNb := 5
	jobs, err := simu.LaunchSimu(testNb)
	if err != nil {
		fmt.Printf("Error:\n%v", err)
		return 
	}
	shortQueue, err := queues.CreateNode(testNb, jobs)
	fmt.Println("----------\nShowing nodes for short queue:")
	queues.ShowNodes(shortQueue)
	queues.LaunchQueue(shortQueue)
	// queues.ShowNodes(shortQueue)

	// nb := 0
	// fmt.Printf("----------\nDeleting node nb %d:\n", nb)
	// err = queues.DelNode(nb,&shortQueue)
	// if err != nil {
	// 	fmt.Printf("Error: %s", err)
	// 	return 
	// }
	// fmt.Println("----------\nShowing nodes for short queue:")
	// queues.ShowNodes(shortQueue)
	// fmt.Println()
	// queues.ShowNode(shortQueue, 1)
	// fmt.Println("----------")
}

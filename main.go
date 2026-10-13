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
	}
	shortQueue, err := queues.CreateNode(testNb, jobs)
	fmt.Println("----------\nShowing nodes for short queue:")
	queues.ShowNodes(shortQueue)
}

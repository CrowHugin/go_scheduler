package virt

import (
	"fmt"
	"errors"
	// "os"
)

type state int 
const (
	STOP = iota
	RUNNING
	KILLED
	WAITING
)

type job struct {
	id int
	name string
	user string
	size uint32
}

type node struct {
	nb int
	next *node
	run	job
}

func setupNode(nbNode int) (*node, error){
	NodeError := errors.New("while creating nodes / queue")
	fmt.Println("----------\nCreating nodes\n----------")
	if nbNode <= 0 {
		return nil, NodeError
	}
	head := &node{nb: 0}
	current := head

	for i:= 1; i<nbNode+1; i++ {
		newNode := &node{nb: i}
		current.next = newNode
		current.run.id = 1
		current.run.name = "test jobs"
		current.run.user = "bku"
		current.run.size = 13
		// parsing de file
		// current.run.id = fonction random
		// current.run.name = job fourni
		// current.run.user = name user
		// current.run.size = get_size
		current = newNode
	}
	return head, nil
}

func showNode(queue *node, nbNode int) (error){
	current := queue
	for i:=0; i < nbNode && current != nil ;i++{
		fmt.Printf("Node index: %d: %v\n", i, current.run)
		current = current.next
	}
	return nil
}

// shows the job with the given id
func ShowJob(id int, queue *node, nbNode int) (error) {
	IdNotFoundError := errors.New("Cannot find id")
	current := queue
	for i:=0;i<nbNode && current != nil;i++{
		if id == current.run.id{
			fmt.Printf("%v\n", current.run)
			fmt.Printf("%d\n", current.nb)
			return nil
		}
		current = current.next
	}
	return IdNotFoundError
}

func LaunchSimu() (error) {
	maxNode := 10
	fmt.Println("launching virtu")
	queue, err := setupNode(maxNode)
	if err != nil{
		return err
	}
	showNode(queue, maxNode)
	ShowJob(1, queue, maxNode)

	return nil
}

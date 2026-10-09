package virt

import (
	"fmt"
	"errors"
)

type state int 
const (
	STOP = iota
	RUNNING
	KILLED
	WAITING
)

type job struct {
	id uint16
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
	fmt.Println("----------\nCreating nodes")
	if nbNode <= 0 {
		return nil, NodeError
	}
	head := &node{nb: 0}
	current := head

	for i:= 1; i<nbNode; i++ {
		newNode := &node{nb: i}
		current.next = newNode
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
		fmt.Printf("Node index: %d\n", i)
		current = current.next
	}
	return nil
}

func LaunchSimu() (error) {
	maxNode := 10
	fmt.Println("launching virtu")
	queue, err := setupNode(maxNode)
	if err != nil{
		return err
	}
	showNode(queue, maxNode)

	return nil
}

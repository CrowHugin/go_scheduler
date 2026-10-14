package queues

import (
	"errors"
	"fmt"
	"strings"
	"strconv"
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

type Node struct {
	nb int
	next *Node
	run	job
}

func getJobs(jobs []string) ([][]string, error){
	var jobSlice [][] string
	if len(jobs) <= 0 {
		return nil, errors.New("no jobs")
	}
	for _, jobStr := range jobs {
		parts := strings.Split(jobStr, ";")
		jobSlice = append(jobSlice, parts)
	}
	return jobSlice, nil
}

func CreateNode(nbNodes int, jobs []string) (*Node, error){
	NodeError := errors.New("Error while creating nodes")
	if nbNodes <= 0 {
		return nil, NodeError
	}
	jobList, err := getJobs(jobs)
	if err != nil {
		return nil, errors.New("Error while getting jobs")
	}
	id, _ := strconv.Atoi(jobList[0][0])
	size, _ := strconv.ParseUint(jobList[0][3], 10, 32)
	j := job{
		id: id,
		name: jobList[0][1],
		user: jobList[0][2],
		size: uint32(size),
	}
	head := &Node{
		nb: 0,
		run: j,
}
	current := head
	for i := 1; i<nbNodes; i++ {
		id, _ := strconv.Atoi(jobList[i][0])
		size, _ := strconv.ParseUint(jobList[i][3], 10, 32)

		j := job{
			id: id,
			name: jobList[i][1],
			user: jobList[i][2],
			size: uint32(size),
		}
		newNode := &Node{
			nb: i,
			run: j,
		}

		current.next = newNode
		current = newNode
	}
	return head, nil
}

func ShowNodes(queue *Node){
	current := queue
	for i:=0;current != nil;i++{
		fmt.Printf("Node index: %d: %v\n", i, current.run)
		current = current.next
	}
}

func ShowNode(queue *Node, nbNode int){
	if queue == nil {
		fmt.Println("Queue is empty")
		return 
	}
	current := queue
	for i:=0; i < nbNode && current != nil ;i++{
		current = current.next
	}
	if current == nil {
		fmt.Printf("Node index %d not found", nbNode)
		return
	}
	fmt.Printf("Node index: %d: %v\n", current.nb, current.run)
}

func reindexQueue(queue *Node) {
	current := queue
	index := 0
	for current != nil {
		current.nb = index
		index++
		current = current.next
	}
}

func DelNode(id int, queue **Node) (error){
	if queue == nil || *queue == nil {
		return errors.New("Queue is empty")
	}
	
	if (*queue).nb == id {
		*queue = (*queue).next
		reindexQueue(*queue)
		return nil
	}
	current := *queue
	for current.next != nil {
		if current.next.nb == id {
			current.next = current.next.next
			reindexQueue(*queue)
			return nil
		}
		current = current.next
	}
	return errors.New("node not in queue")
}

// shows the job with the given id
// any instead of id later on - so job can be found by user size etc
func ShowJob(id int, queue *Node, nbNode int) (error) {
	IdNotFoundError := errors.New("Cannot find id")
	current := queue
	for i:=0;i<nbNode && current != nil;i++{
		if id == current.run.id{
			fmt.Printf("id: %d | name: %s | user: %s | size: %d\n", current.run.id, current.run.name, current.run.user, current.run.size)
			fmt.Printf("%d\n", current.nb)
			return nil
		}
		current = current.next
	}
	return IdNotFoundError
}

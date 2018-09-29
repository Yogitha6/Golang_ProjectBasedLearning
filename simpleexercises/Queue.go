package main

import (
	"fmt"
)

type node struct {
	value int
	next *node
}

type queue struct {
	front *node
	last *node
}

func (q *queue) enQueue(input int) {
	newnode := node {input, nil}
	if q.isEmpty() {
		q.last = &newnode
		q.front = &newnode
	} else {
		q.last.next = &newnode
		q.last = q.last.next
	}
}

func (q *queue) deQueue() *node{
	if q.front == q.last {
		q.last = q.front.next	
	}
	temp := q.front
	q.front = q.front.next
	return temp
}

func (q *queue) isEmpty() bool {
	if q.last == nil && q.front == nil {
		return true
	}
	return false
}

func (q *queue) size() int {
	temp := q.front
	size := 0
	for temp!=nil {
		size = size + 1
		temp = temp.next
	}
	return size
}

func main() {
	fmt.Println("Queue program for BFS")
	queueList := queue{}
	queueList.enQueue(10)
	queueList.enQueue(12)
	queueList.enQueue(11)	
	fmt.Println(queueList.size())
	fmt.Println(queueList.deQueue().value)
	fmt.Println(queueList.deQueue().value)
	fmt.Println(queueList.deQueue().value)
	fmt.Println(queueList.size())
	fmt.Println(queueList.isEmpty())
}

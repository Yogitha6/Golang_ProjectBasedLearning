package main

import (
	"fmt"
)

//Queue part of the program

type node struct {
	value int
	next  *node
}

type queue struct {
	front *node
	last  *node
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
	temp := q.front
	if q.front == q.last {
		q.last = q.front.next
	}
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
	for temp != nil {
		size = size + 1
		temp = temp.next
	}
	return size
}

//Graph part of the program

type graph struct {
	totalVertices int
	//slice of slices
	edges [][]int
}

func (g *graph) addEdge(src int, dest int) {
	//undirected graph
	g.edges[src] = append(g.edges[src], dest)
	g.edges[dest] = append(g.edges[dest], src)
}

func (g *graph) printEdges() {
	for index, adjacentslc := range g.edges {
		for _, value := range adjacentslc {
			fmt.Println(index, "->", value)
		}
	}
}

func (g *graph) getAdjacentVertices(src int) []int {
	return g.edges[src]
}

func (g *graph) BFS(current int) {

	//use queue to maintain the order
	queueList := queue{}
	visited := make([]bool, g.totalVertices)

	fmt.Println(current)
	queueList.enQueue(current)
	visited[current]=true

	for queueList.isEmpty() == false {

		//deQueue and start exploring
		adj := g.getAdjacentVertices(queueList.deQueue().value)
		for _, adjvertx := range adj {
			if visited[adjvertx] != true {
			fmt.Println(adjvertx) 			
			queueList.enQueue(adjvertx)
			visited[adjvertx] = true
			}
		}

	}

}


func main() {
	fmt.Println("Graph Implementation")
	newGraph := &graph{totalVertices: 5, edges: make([][]int, 5)}

	newGraph.addEdge(4, 2)
	newGraph.addEdge(3, 2)
	newGraph.addEdge(3, 1)
	newGraph.addEdge(0, 1)
	newGraph.printEdges()
	fmt.Println(newGraph.getAdjacentVertices(2))
	fmt.Println("BFS:")
	newGraph.BFS(1)
}

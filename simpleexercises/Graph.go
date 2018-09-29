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
	queueList := queue{}

	fmt.Println(current)
	queueList.enQueue(current)

	for queueList.isEmpty() == false {

		//deQueue and start exploring
		fmt.Println(queueList.size())
		adj := g.getAdjacentVertices(queueList.deQueue().value)
		for _, adjvertx := range adj {
			fmt.Println(adjvertx) //visited
			queueList.enQueue(adjvertx)
			fmt.Println(adjvertx, "added to the queue") // add to queue
		}

	}

}

/*func (g *graph) DFS() {



}*/

func main() {
	fmt.Println("Hello, playground")
	newGraph := &graph{totalVertices: 5, edges: make([][]int, 5)}

	newGraph.addEdge(1, 2)
	newGraph.addEdge(4, 2)
	newGraph.addEdge(3, 2)
	newGraph.addEdge(3, 1)
	//newGraph.printEdges()
	//fmt.Println(newGraph.getAdjacentVertices(2))
	newGraph.BFS(1)
}

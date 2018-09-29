package main

import (
	"fmt"
)

//Stack part of the program
type node struct {
	value int
	bottom *node
}

type stack struct {
	top *node
}

func (s *stack) push(value int) {
	if s.top==nil {
		n := node {value: value, bottom: nil}
		s.top = &n
	} else {
		n := node {value: value, bottom: s.top}
		s.top = &n
	}
}

func (s *stack) pop() int {
	if s.top == nil {
		fmt.Println("No element to Pop")
		return -1
	} else {
		temp := s.top.value
		s.top = s.top.bottom
		return temp
	}
}

func (s *stack) peek() int {
	if s.top == nil {
		fmt.Println("No element to Peek")
		return -1
	} else {
		return s.top.value
	}
}

func (s *stack) isEmpty() bool {
	if s.top == nil {
		return true
	}
	return false
}

func (s *stack) size() int {
	size := 0
	temp:= s.top
	for temp!=nil {
		temp = temp.bottom
		size = size + 1
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

func (g *graph) DFS(current int) {
	st := stack{}
	visited := make([]bool, g.totalVertices)
	
	//push current item and visit it
	fmt.Println(current)
	st.push(current)
	visited[current] = true
	
	for st.isEmpty()==false {
	     adjtoexplore := g.findAdjacentNotVisited(current, visited)
		if adjtoexplore != -1 {
			fmt.Println(adjtoexplore)
			st.push(current)
			visited[adjtoexplore] = true
			current = adjtoexplore
		} else {
			current = st.pop()
		}
	
	}
	
}

func (g *graph) findAdjacentNotVisited(current int, visited []bool) int {
	adjvercs:= g.edges[current]
	for _, value := range adjvercs {
		if visited[value] == false {
			return value
		}
	}
	return -1
}

func main() {
	fmt.Println("Graph Implementation")
	newGraph := &graph{totalVertices: 5, edges: make([][]int, 5)}

	newGraph.addEdge(1, 4)
	newGraph.addEdge(1, 2)
	newGraph.addEdge(4, 3)
	newGraph.addEdge(4, 2)
	newGraph.addEdge(0, 3)
	newGraph.printEdges()
	fmt.Println(newGraph.getAdjacentVertices(4))
	fmt.Println("DFS:")
	newGraph.DFS(1)
}

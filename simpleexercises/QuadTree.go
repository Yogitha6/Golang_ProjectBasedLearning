package main

import (
	"fmt"
	"math/rand"
)

type coordinate struct {
	x int
	y int
}

type node struct {
	value    coordinate
	x1y1     coordinate
	x2y1     coordinate
	x1y2     coordinate
	x2y2     coordinate
	quadrant string
	children []*node
	parent   *node
	divided  bool
}

type quadtree struct {
	root  *node
	count int
}

func (qt *quadtree) initialize(boundary int) {
	point := coordinate{0, 0}
	x1y1 := coordinate{0, 0}
	x2y1 := coordinate{boundary, 0}
	x1y2 := coordinate{0, boundary}
	x2y2 := coordinate{boundary, boundary}

	n := node{value: point, x1y1: x1y1, x2y1: x2y1, x1y2: x1y2, x2y2: x2y2, divided: false, quadrant: "0"}

	qt.root = &n
	qt.count = qt.count + 1

}

func (qtnode *node) addCoordinate(c coordinate) bool {

	//check if the point is in the boundaries if so go forward to insert else return
	if c.x > qtnode.x2y2.x || c.y > qtnode.x2y2.y || c.x < qtnode.x1y1.x || c.y < qtnode.x1y1.y {
		return false
	}

	if len(qtnode.children) < 4 {
		//if there are less than 4 children then just add this as a child
		n := node{value: c, parent: qtnode, divided: false, quadrant: "0"}
		qtnode.children = append(qtnode.children, &n)
	} else {
		// if there are already 4 children then we need to divide and add
		if qtnode.divided == false {
			qtnode.divide(c)
		}
		for _, child := range qtnode.children {
			//fmt.Println("adding ", c)
			if child.addCoordinate(c) {
				//fmt.Println("added to ", child.quadrant, child.x1y1, child.x2y1, child.x1y2, child.x2y2)
			}
		}

	}
	return true
}

func (qtnode *node) divide(c coordinate) {
	q1 := node{x1y1: qtnode.x1y1, x2y1: coordinate{qtnode.x2y1.x / 2, qtnode.x2y1.y}, x1y2: coordinate{qtnode.x1y2.x, qtnode.x1y2.y / 2}, x2y2: coordinate{qtnode.x2y2.x / 2, qtnode.x2y2.y / 2}, divided: false, quadrant: "Q1"}
	q2 := node{x1y1: coordinate{qtnode.x2y1.x / 2, qtnode.x2y1.y}, x2y1: qtnode.x2y1, x1y2: coordinate{qtnode.x2y2.x / 2, qtnode.x2y2.y / 2}, x2y2: coordinate{qtnode.x2y2.x, qtnode.x2y2.y / 2}, divided: false, quadrant: "Q2"}
	q3 := node{x1y1: coordinate{qtnode.x2y2.x / 2, qtnode.x2y2.y / 2}, x2y1: coordinate{qtnode.x2y2.x, qtnode.x2y2.y / 2}, x1y2: coordinate{qtnode.x2y2.x / 2, qtnode.x2y2.y}, x2y2: qtnode.x2y2, divided: false, quadrant: "Q3"}
	q4 := node{x1y1: coordinate{qtnode.x1y2.x, qtnode.x1y2.y / 2}, x2y1: coordinate{qtnode.x2y2.x / 2, qtnode.x2y2.y / 2}, x1y2: qtnode.x1y2, x2y2: coordinate{qtnode.x2y2.x / 2, qtnode.x2y2.y}, divided: false, quadrant: "Q4"}
	existingchildren := qtnode.children
	qtnode.children = []*node{}
	qtnode.children = append(qtnode.children, &q1, &q2, &q3, &q4)
	qtnode.divided = true

	for _, child := range existingchildren {
		qtnode.addCoordinate(child.value)
	}
}

func (qtnode node) printQuadTree() {
	empty := coordinate{0, 0}
	if qtnode.quadrant == "0" && qtnode.value == empty {
		fmt.Println("root")
	} else if qtnode.quadrant == "0" && qtnode.value != empty {
		fmt.Println("child", qtnode.value)
	} else {
		fmt.Println("Quadrant: ", qtnode.quadrant)
		fmt.Println("boundaries : ", qtnode.x1y1, qtnode.x2y1, qtnode.x1y2, qtnode.x2y2)
	}
	for _, child := range qtnode.children {
		child.printQuadTree()
	}

}

/*func (qt *quadtree) qtnearestNeighbours(input coordinate, length) []coordinate{
	//given point from root - pick a child and traverse
	
	
	//if found, print all children under that parent whose distance is less than given length


}*/

func main() {
	fmt.Println("Quad Tree Implementation")
	qt := quadtree{nil, 0}
	qt.initialize(1000)

	for i := 0; i < 30; i++ {
		qt.root.addCoordinate(coordinate{rand.Intn(1000), rand.Intn(1000)})
	}
	qt.root.printQuadTree()
	
	//input := coordinate{300, 200}
	//qt.nearestNeighbours(input, 50)
}

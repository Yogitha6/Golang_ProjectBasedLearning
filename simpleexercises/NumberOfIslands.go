package main

import (
	"fmt"
)

func findIslands(input [][]int) int {
  count := 0
  x_indices := make([]int, 0)
  y_indices := make([]int, 0)

   for x, oned := range input {
     for y, value := range oned {
        if value == 1 {
	   //find if adjacent positions have 1, if not then increase counter
	    if findAdjacency(x_indices, y_indices,x, y) == false {
	       count = count + 1
	    }
	    x_indices = append(x_indices, x)
	    y_indices = append(y_indices, y)
	}
     }
   }
   return count
}

func findAdjacency(xcords []int, ycords []int, x int, y int) bool{
    adjacent := false
    for i:=0; i<len(xcords); i++ {
       if Abs(xcords[i] - x) <= 1 && Abs(ycords[i] - y) <= 1 {
  	adjacent = true
       }
     }
    return adjacent
}

func Abs(x int) int {
   if x < 0 {
       return -x
     }
   return x
}

func main() {
	fmt.Println("Hello, playground")
	input := [][]int {{1,1,0}, {0,0,1}, {1,0,1}}
	fmt.Println(findIslands(input))
}

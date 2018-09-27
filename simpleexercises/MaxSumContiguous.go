package main

import (
	"fmt"
)


func findMaxSumCon(input []int) int {
	var sum int = 0
	var maxSum int = input[0]
	for i:=0; i<len(input); i++ {
	
	if input[i] > sum + input[i] {
	  sum = input[i]
	  } else {
	  sum = sum + input[i]
	  }
	
	  if sum > maxSum {
	   maxSum = sum
	   }	
	}
    return maxSum
}

func main() {
	fmt.Println("Hello, playground")
	input:= []int{-2,3,3,-1,-2,1,5,-3}
	fmt.Println(findMaxSumCon(input))
	input= []int{10,-8,-3,2,3,6,9,-10,-20,-4,16}
	fmt.Println(findMaxSumCon(input))
	input= []int{-5, 6, 7, 1, 4, -8, 16}
        fmt.Println(findMaxSumCon(input))
        input= []int{-1,-2,1,-2,-2,0,1}
        fmt.Println(findMaxSumCon(input))
}
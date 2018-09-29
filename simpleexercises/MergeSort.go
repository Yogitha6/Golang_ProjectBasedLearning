package main

import (
	"fmt"
)

func mergesort(slc []int) []int{
	if len(slc) <= 1 {
	   return slc
	}
	
	leftslc, rightslc := split(slc)
	
	leftslc = mergesort(leftslc)
	rightslc = mergesort(rightslc)
	
	mergedslc := mergeslcs(leftslc, rightslc)
	return mergedslc
}

func split(slc []int) ([]int, []int) {
	return slc[0:len(slc)/2], slc[len(slc)/2:]
}

func mergeslcs(leftslc []int, rightslc []int)[]int {
        size := len(leftslc) + len(rightslc)
	output := make([]int, size)
	for l, r,o := 0, 0, 0; o < size; o++{
	     
	     if l >= len(leftslc) {
	        output[o] = rightslc[r]
	        r = r + 1
		continue
		} else if r >= len(rightslc) {
		output[o] = leftslc[l]
		l = l + 1
		continue
		}
		
	     if leftslc[l] < rightslc[r] {
		output[o] = leftslc[l]
	        l = l + 1
		} else {
		output[o] = rightslc[r]
	        r = r + 1
		}
	}
	return output
}

func main() {
	fmt.Println("Hello, playground")
	input:= []int{8,3,2,1,6, 5}
	
	fmt.Println("Sorted : ", mergesort(input))
	
}

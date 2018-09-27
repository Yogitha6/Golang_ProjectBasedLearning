package main

import (
	"fmt"
)


/*func findMinMax(input []int) (int, int) {
   min = math.MaxInt32
   max = math.MinInt32
   for i:=0; i<len(input); i++ {
   if input[i] < min {
    min = input[i]
   }
   if input[i] > max {
    max = input[i]
    }
   }
   return min,max
}*/

func addallPosNeg(input []int) (int, int) {
       var posSum, negSum int
	for i:=0; i<len(input); i++ {
	   if input[i] > 0 {
	      posSum = posSum + input[i]
	    } else {
	      negSum = negSum + input[i]
		}
	  }
	return posSum, negSum
}


func findMaxSumCon(input []int, expectedsum int) int{
	var localsum int = 0
	totalposSum, totalnegSum := addallPosNeg(input)
	
	for i:=0; i<len(input); i++ {
	
	if input[i] < 0 {
	totalnegSum = totalnegSum - input[i]
	
	} else {
	totalposSum= totalposSum - input[i]
	}
	
	if localsum + input[i] == expectedsum {
	    return 1
	   } else if input[i] == expectedsum {
	    return 1
	   } else {
	   localsum = localsum + input[i]
	   }

	if localsum + input[i] + totalposSum < expectedsum {
		  localsum = input[i]
	} else if localsum + input[i] - totalnegSum > expectedsum {
	   localsum = input[i]
	} 

}	
	return -1

}

func main() {
	fmt.Println("Hello, playground")
	input:= []int{10, 2, -2, -20, 10}
	fmt.Println(findMaxSumCon(input, -10))
	input= []int{1, 4, 20, 3, 10, 5}
	fmt.Println(findMaxSumCon(input, 33))
}

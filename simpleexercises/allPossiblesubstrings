package main

import (
	"fmt"
)

func createsubstrings(input string) {
   for i:=0; i<len(input); i++ {
       var substring string = string(input[i])
       fmt.Println(substring)
         for j:=i+1; j<len(input); j++ {
           substring = substring + string(input[j])
           fmt.Println(substring)
        }
    }
}

func main() {
	fmt.Println("Hello, playground")
	input := "abcd"
	createsubstrings(input)	
}

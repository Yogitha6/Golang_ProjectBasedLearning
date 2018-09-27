package main

import (
	"fmt"
)

func longestPalindrome(s string) string {

    var startindex int = 0
    var maxLength int
    var state [1000][1000]bool
    
    for i:=0; i<len(s); i++ {
        state[i][i] = true
        maxLength = 1
    }
    
    for j:=0; j<len(s)-1; j++ {
        if s[j] == s[j+1] {
            state[j][j+1] = true
            maxLength = 2
        }
    }

    
    for k:=3; k<=len(s); k++ {
        for l:=0; l< len(s)-k+1; l++ {
            endindex:= l+k-1
            if state[l+1][endindex-1] == true && s[l] == s[endindex] {
                state[l][endindex]=true
            if maxLength < k {
                 maxLength = k
                 startindex = l
            }
            }         
        }
    }
return s[startindex:startindex+maxLength]

}

func main() {
	fmt.Println("Hello, playground")
	fmt.Println(longestPalindrome("abba"))
}
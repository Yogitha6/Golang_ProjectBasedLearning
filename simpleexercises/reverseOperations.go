package main

import (
	"fmt"
	"strings"
)

func reverseSlice(inputSlice []int) []int {
	sliceLength := len(inputSlice)
	for i:=0; i<sliceLength/2; i++ {
		temp := inputSlice[i]
		inputSlice[i] = inputSlice[sliceLength-i-1]
		inputSlice[sliceLength-i-1] = temp
	}

	//testing built in methods from sort pkg
	//sort.Reverse(sort.IntSlice(inputSlice))

	return inputSlice

}

func reverseStringSlice(inputSlice []string) []string {
	sliceLength := len(inputSlice)
	for i:=0; i<sliceLength/2; i++ {
		temp := inputSlice[i]
		inputSlice[i] = inputSlice[sliceLength-i-1]
		inputSlice[sliceLength-i-1] = temp
	}

	//testing built in methods from sort pkg
	//sort.Reverse(sort.StringSlice(inputSlice))

	return inputSlice
}

func reverseString(input string) string {
	characters := []rune(input)
	sliceLength := len(characters)
	for i:=0; i<sliceLength/2; i++ {
		temp := characters[i]
		characters[i] = characters[sliceLength-i-1]
		characters[sliceLength-i-1] = temp
	}
	return string(characters)
}

func reverseWordsInString(input string) string{
	words := strings.Fields(input)
	for i, j:= 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
}

func main() {
  var givenSlice []int = []int {9,3,2,6,4}
  fmt.Println(reverseSlice(givenSlice))
  givenSlice = []int {0,1,2,3,4,5}
  fmt.Println(reverseSlice(givenSlice))
  givenSlice = []int {0}
  fmt.Println(reverseSlice(givenSlice))

  var inputSlice []string = []string {"hello", "world", "sample"}
  fmt.Println(reverseStringSlice(inputSlice))

  stringInput := "hello"
  fmt.Println(reverseString(stringInput))

  wordsInput := "this is a test"
  fmt.Println(reverseWordsInString(wordsInput))

  fmt.Println(reverseWordsInString("test"))
  fmt.Println(reverseWordsInString("odd number words"))

  //To Sort you can use:
  // sort.Ints(givenSlice)
  // sort.Strings(inputSlice)
}
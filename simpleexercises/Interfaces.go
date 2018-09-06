package main

import ( 
	"fmt"
	"reflect"
)

func main() {
    // This program helps to understand the behavior of Kind() method of reflect package
    var x float64= 3.14
	var y interface{} = x
	fmt.Println("value:", reflect.ValueOf(x).String())
	fmt.Println("value:", reflect.ValueOf(x).Float())
	fmt.Println("value:", reflect.TypeOf(y))
	fmt.Println("value:", reflect.TypeOf(y).Kind())

	type MyInt int
	var a MyInt = 7
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.ValueOf(a).Type())
	fmt.Println(reflect.ValueOf(a).Kind())
	fmt.Println(reflect.ValueOf(a))
}
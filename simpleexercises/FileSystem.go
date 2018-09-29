package main

import (
	"fmt"
)

type node struct {
	name     string
	dtype    string
	children []*node
	parent   *node
}

type basedirectory struct {
	count int
	root  *node
}

func initializeDirectory() basedirectory {
	rootnode := node{name: "/", dtype: "dir", children: []*node{}, parent: nil}
	dir := basedirectory{1, &rootnode}
	return dir
}


func (currdir *node) changeDirectory(name string) *node{
	//list - sequential search, map would have given better search performance
	for _, childdir := range currdir.children {
		if childdir.name == name {
			return childdir
		}
	}
	return nil
}

func (currdir *node) showPWD() {
	fmt.Println(currdir.name)
	//fmt.Println("Following are the objects in this directory: ")
	for _,childpointer := range currdir.children {
		fmt.Println(childpointer.name)
	}
}

func (currdir *node) add(name string, dtype string) {
	nodeobj := node{name: name, dtype: dtype, children: []*node{}, parent: currdir}
	currdir.children = append(currdir.children, &nodeobj)
}

func (currdir *node) movetoParentDir() *node{
	return currdir.parent
}

func main() {
	fmt.Println("Weclome to a simple file system using tree structure")
	baseroot := initializeDirectory()
	currdir := baseroot.root
	fmt.Println("Created a root directory")
	currdir.showPWD()
	
	fmt.Println("Adding directories and files to root directory")
	currdir.add("home", "dir")
	currdir.add("etc", "dir")
	currdir.add("var", "dir")
	currdir.add("opt", "dir")
	currdir.add("test.txt", "file")
	currdir.showPWD()
	
	fmt.Println("Changing to home directory")
	currdir = currdir.changeDirectory("home")
	currdir.showPWD()
	
	fmt.Println("Adding directories to current directory")
	currdir.add("yogitha", "dir")
	currdir.add("user2", "dir")
	currdir.add("user3", "dir")
	currdir.showPWD()
	
	fmt.Println("Moving back to parent directory")
	currdir = currdir.movetoParentDir()
	currdir.showPWD()
}

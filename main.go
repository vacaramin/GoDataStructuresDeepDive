package main

import (
	"fmt"

	linkList "github.com/vacaramin/GoDataStructuresDeepDive/LinkList"
)

func main() {
	fmt.Println("Welcome to the world of Data Structures in GO")
	fmt.Println("\n\t**********Creating a LinkList**********")
	var linklist linkList.LinkList

	fmt.Println("\nInserting 1 and 2 in the LinkList")

	linklist.Add(1)
	linklist.Add(2)
	linklist.Print()

	fmt.Println("\n\nReversing the LinkList")
	linklist.Reverse()

	linklist.Print()

	fmt.Println("\n\nGet index 0")
	val, err := linklist.Get(0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(val)
	}

	fmt.Println("\nGet index 1")
	val, err = linklist.Get(1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(val)
	}
	fmt.Println("\nGet index 3")
	val, err = linklist.Get(3)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(val)
	}

}

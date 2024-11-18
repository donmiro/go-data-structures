package main

import (
	"fmt"
	"go-data-structures/linkedlists"
)

func main() {
	linkedList := linkedlists.LinkedList{}
	linkedList.AddToHead(100)
	linkedList.AddToHead(3)
	linkedList.AddToHead(30)
	// fmt.Println(linkedList.HeadNode.Property)
	// linkedList.IterateList()
	fmt.Println(linkedList.LastNode())
}

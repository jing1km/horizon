package main

import "fmt"

type LinkedNode struct {
	Val  int32
	Next *LinkedNode
}

func main() {
	var l1 []*LinkedNode
	fmt.Println(l1)
}

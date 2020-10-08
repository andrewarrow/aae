package main

import "fmt"

func main() {
	fmt.Println("hello1")
	list := []int{0, 542, 5424, 54245, 94245, 99999, 94245, 54245, 5424, 542, 0, -542, -5424, -54245, -94245, -99999, -94245, -54245, -5424, -542, 0}
	fmt.Println(list)
	prev := 0
	for i, item := range list {
		delta := item - prev
		fmt.Println(i, item, prev, delta)
		prev = item
	}
}

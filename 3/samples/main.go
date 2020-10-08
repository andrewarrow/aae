package main

import "fmt"

func main() {
	fmt.Println("hello1")
	list := []int{0, 542, 5424, 54245, 94245, 99999, 94245, 54245, 5424, 542, 0, -542, -5424, -54245, -94245, -99999, -94245, -54245, -5424, -542, 0}
	fmt.Println(list)
	prev := 0
	prevDelta := 0
	for i, item := range list {
		delta := item - prev
		compute(delta, prevDelta)
		fmt.Println(i, item, prev, delta)
		prev = item
		prevDelta = delta
	}
}

func compute(delta, prevDelta int) {
	d := delta
	p := prevDelta

	if d < 0 {
		d = d * -1
	}

	if p < 0 {
		p = p * -1
	}

	if p < d {
		val := float64(d-p) / float64(d)
		fmt.Printf("------- %d %d %.2f\n", p, d, val)
	} else if p > d {
		val := float64(p-d) / float64(p)
		fmt.Printf("!------ %d %d %.2f\n", p, d, val)
	}
}

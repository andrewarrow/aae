package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	data, _ := ioutil.ReadFile("hello.aiff") // read in 31,440 numbers
	count := 1
	for i, b := range data[4095:] {
		fmt.Printf("%d (%d) ", b, i)
		if count == 2 {
			fmt.Println("")
			count = 0
		}
		count++
		//converted := float64(b) / math.Pow(2, 16)
		// math.Float64frombits(binary.LittleEndian.Uint64(sample)))

	}
	fmt.Println("")
	fmt.Println("")

}

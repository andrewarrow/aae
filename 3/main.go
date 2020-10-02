package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	data, _ := ioutil.ReadFile("hello.aiff") // read in 31,440 numbers
	fmt.Println(data)
}

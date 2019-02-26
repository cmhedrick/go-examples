package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	data, err := ioutil.ReadFile("text.txt")
	if err != nil {
		return
	}
	fmt.Println(string(data))
}

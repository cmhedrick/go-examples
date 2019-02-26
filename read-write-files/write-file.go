package main

import "os"

func main() {
	file, err := os.Create("text.txt")
	if err != nil {
		return
	}
	defer file.Close()

	file.WriteString("test\nhello")
}

package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func readFile(inFile string) string {
	data, err := ioutil.ReadFile(inFile)
	if err != nil {
		fmt.Println("[X] Failed to read")
	}

	return string(data)
}

func writeFile(newFile string, fileData string) bool {
	file, err := os.Create(newFile)
	if err != nil {
		return false
	}
	defer file.Close()

	file.WriteString(fileData)
	return true
}

func main() {
	fmt.Println("[-] Reading file")
	fileData := readFile("test.txt")
	fmt.Println("[+] File read!")
	fmt.Println("[-] Writing file")
	writeFile("copied-file.txt", fileData)
	fmt.Println("[?] File written")
}

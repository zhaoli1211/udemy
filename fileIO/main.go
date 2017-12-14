package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	b, err := ioutil.ReadFile("/Users/lizhao/Documents/Go/src/github.com/zhaoli1211/udemy/fileIO/output")

	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))

	//write
	stringContent := string(b) + "\nmore new stuff"
	fmt.Println("write to file", stringContent)
	err = ioutil.WriteFile("/Users/lizhao/Documents/Go/src/github.com/zhaoli1211/udemy/fileIO/output", []byte(stringContent), 0644)
	if err != nil {
		panic(err)
	}
}

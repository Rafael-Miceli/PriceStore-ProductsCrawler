package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	fmt.Println("Initializing Products Crawler")

	b, err := ioutil.ReadFile("test-page.html") // just pass the file name

	if err != nil {
		fmt.Print(err)
	}

	str := string(b) // convert content to a 'string'

	if strings.Contains(str, "R$") {
		fmt.Printf("Found product price ")
	} else {
		fmt.Printf("NOT Found product price ")
	}
}

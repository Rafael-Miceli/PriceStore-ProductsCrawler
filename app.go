package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

func main() {
	fmt.Println("Initializing Products Crawler")

	str := getMarketHTML()

	if regexp.MustCompile(`R\$`).MatchString(str) {
		fmt.Printf("Found product price ")
	} else {
		fmt.Printf("NOT Found product price ")
	}
}

func getMarketHTML() string {
	b, err := ioutil.ReadFile("test-page.html") // just pass the file name

	if err != nil {
		fmt.Print(err)
	}

	return string(b)
}

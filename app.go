package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

func main() {
	fmt.Println("Initializing Products Crawler")

	str := getMarketHTML()

	r, err := regexp.Compile(`R\$`)

	if err != nil {
		fmt.Printf("There is a problem with your regexp.\n")
		return
	}

	if r.MatchString(str) == true {
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

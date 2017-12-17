package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

func main() {
	fmt.Println("Initializing Products Crawler")

	str := getMarketHTML()

	r, err := regexp.Compile(`R\$&nbsp;[0-9]+,[0-9]+`)

	if err != nil {
		fmt.Print(err)
	}

	res := r.FindAllString(str, -1)

	fmt.Printf("%v", res)

}

func getMarketHTML() string {
	b, err := ioutil.ReadFile("test-page.html") // just pass the file name

	if err != nil {
		fmt.Print(err)
	}

	return string(b)
}

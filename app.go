package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"golang.org/x/net/html"
)

var (
	page1 = "http://www.redeultramercado.com.br/bovinos-aves-e-suinos.html"
)

func main() {
	fmt.Println("Initializing Products Crawler")

	//str := getMarketHTMLLocal()

	resp, err := http.Get(page1)

	if err != nil {
		fmt.Print(err)
	}

	//body, err := ioutil.ReadAll(resp.Body)
	z := html.NewTokenizer(resp.Body)

	//fmt.Printf("%v", string(body))

	//resp.Body.Close()

	fmt.Println("Preparando para iterar sobre html")
	fmt.Println("Z ", z)

	for {
		tt := z.Next()

		switch {
		case tt == html.ErrorToken:
			fmt.Println("Fim do html")
			// End of the document, we're done
			return
		case tt == html.StartTagToken:
			t := z.Token()

			fmt.Println("t ", t)
			//fmt.Println("t.Data ", t.Data)

			isAnchor := t.Data == "a"
			if isAnchor {
				fmt.Println("We found a link!")
			}
		}
	}

	resp.Body.Close()

	//Não usar Regex!!!

	// r, err := regexp.Compile(`R\$&nbsp;[0-9]+,[0-9]+`)

	// if err != nil {
	// 	fmt.Print(err)
	// }

	// res := r.FindAllString(str, -1)

	// //prices :=

	// fmt.Printf("%v", res)

}

func getMarketHTMLLocal() string {
	b, err := ioutil.ReadFile("test-page.html") // just pass the file name

	if err != nil {
		fmt.Print(err)
	}

	return string(b)
}

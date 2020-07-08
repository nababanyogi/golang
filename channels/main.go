package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://amazon.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://google.com",
	}

	c := make(chan string)

	for _, link := range links {
		// fmt.Println(link)
		go checkLink(link, c)
	}

	// //only checked once
	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-c)
	// }

	// // ping all the time
	// for {
	// 	go checkLink(<-c, c)
	// }

	// // alternative loop sintax
	// for l := range c {
	// 	go checkLink(l, c)
	// }

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			go checkLink(link, c)
		}(l)
	}

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "Migh be down!")
		c <- link
	}

	fmt.Println(link, "is up!")
	c <- link
}

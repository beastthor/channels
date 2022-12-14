package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {

		go checkLink(link, c)

	}

	//for i := 0; i < len(links); i++ {
	//	fmt.Println(<-c)
	//}

	// to repeat the routines
	//for {
	//go checkLink(<-c, c)
	//}

	//for developers to understand the code and have a clear understanding
	for l := range c {
		//not a good way to put time.sleep here because it is stopping mainmethod
		//time.Sleep(5 * time.Second)

		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(l, c)
		}(l)

		//go checkLink(l, c)
	}

}

func checkLink(link string, c chan string) {

	//time.Sleep(5 * time.Second)
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		//c <- "Might be down I think"
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	//c <- "Yep its up"
	c <- link
}

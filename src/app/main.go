package main

import (
	"log"
	"net/http"
	"sync"
)

var Revision string

var Urls = []string{
	"http://freecoder.dev",
	"http://google.com",
	"http://apple.com",
	"http://microsoft.com",
	"http://github.com",
	"http://amazon.com",
}

var wg sync.WaitGroup

func healthCheck(url string) {
	_, err := http.Get(url)
	if err != nil {
		log.Printf("The website %s is OFF", url)
	} else {
		log.Printf("The website %s is ON", url)
	}
	wg.Done()
}

func main() {
	wg.Add(len(Urls))

	for _, url := range Urls {
		go healthCheck(url)
	}

	wg.Wait()
}

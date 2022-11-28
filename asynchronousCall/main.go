package main

import (
	"fmt"
	"net/http"
	"time"
)

// this example i'll try to check the status of 4 different website and return their status
// i'll use go routines and channels to accomplish this task
type siteResponse struct {
	site       string
	statusCode int
	time       int
}

func main() {
	sites := []string{
		"http://www.google.com",
		"http://www.facebook.com",
		"http://www.amazon.com",
		"http://www.stackoverflow.com",
	}

	c := make(chan siteResponse)

	// call sites
	for _, s := range sites {
		go checkStatus(s, c)
	}

	// check for result from channels
	for i := 0; i < len(sites); i++ {
		result := <-c
		fmt.Printf("calling %s status %d time %dms\n", result.site, result.statusCode, result.time)
	}
}

func checkStatus(url string, c chan siteResponse) {
	t := time.Now()

	resp, _ := http.Get(url)

	// set status
	c <- siteResponse{
		site:       url,
		statusCode: resp.StatusCode,
		time:       int(time.Since(t).Milliseconds()),
	}
}

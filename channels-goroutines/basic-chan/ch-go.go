package main

import (
	"log"
	"net/http"
	"time"
)

type result struct{
	url string
	err error
	latency time.Duration
}

func get(url string, ch chan <- result){
	start := time.Now()
	if resp,err := http.Get(url); err!= nil {
		ch <- result{url, err, 0}
	
	}else{
		t := time.Since(start).Round(time.Millisecond)
		ch <- result{url, nil, t}
		resp.Body.Close()
	}
}

func main(){
	results := make(chan result)

	lists := []string{
		"https://amazon.com",
		"https://google.com",
		"https://nytimes.com",
		"https://wsj.com",

	}

	for _, url := range lists {
		go get(url, results)
	}

	for range lists {
		r := <-results

		if r.err != nil {
			log.Printf("%-20s %s\n", r.url, r.err)
		}else{
			log.Printf("%-20s %s\n", r.url, r.latency)
		}
	}
}
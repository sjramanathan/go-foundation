package main

import (
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

func siteTime(url string) {
	start := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		log.Printf("ERROR: %s -> %s", url, err)
	}
	defer resp.Body.Close()

	if _, err := io.Copy(io.Discard, resp.Body); err != nil {
		log.Printf("ERROR: %s -> %s", url, err)
	}

	duration := time.Since(start)
	log.Printf("INFO: %s -> %v", url, duration)
}

func siteTimeWorker(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	siteTime(url)
}

func main() {
	urls := []string{
		"https://google.com",
		"https://apple.com",
		"https://no-such-site.biz",
	}

	var wg sync.WaitGroup
	wg.Add(len(urls))

	for _, url := range urls {
		go siteTimeWorker(url, &wg)
	}
	wg.Wait()
}

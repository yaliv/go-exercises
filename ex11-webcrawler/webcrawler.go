package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type Response struct {
	url, body string
	err       error
}

type Crawler struct {
	cwls int
	ch   chan Response
	cu   map[string]bool
	mux  sync.Mutex
}

func (c *Crawler) ReadCU(url string) bool {
	c.mux.Lock()
	defer c.mux.Unlock()
	if _, ok := c.cu[url]; ok {
		return true
	}
	return false
}

func (c *Crawler) WriteCU(url string) []string {
	c.mux.Lock()
	defer c.mux.Unlock()
	if _, ok := c.cu[url]; ok {
		return nil
	}
	// Write URL to cu.
	c.cu[url] = true
	// Fetch the URL.
	body, urls, err := fetcher.Fetch(url)
	// Send the response to the channel.
	c.ch <- Response{url: url, body: body, err: err}
	// Return the slice of URLs.
	return urls
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func (c *Crawler) Crawl(url string, depth int, fetcher Fetcher) {
	// Never continue to crawl zero depth.
	if depth <= 0 {
		close(c.ch)
		return
	}
	// When we just started to crawl,
	// our cached URLs (cu) are still empty.
	if len(c.cu) == 0 {
		// We start with one crawler.
		c.cwls = 1
	}

	// Write the current URL to cu
	// and get the slice of URLs in return.
	urls := c.WriteCU(url)

	for _, u := range urls {
		// Do not crawl an existing cached URL.
		if c.ReadCU(u) {
			continue
		}
		if depth > 1 {
			go c.Crawl(u, depth-1, fetcher)
			// Add new crawler.
			c.cwls++
		}
	}

	// This crawler has done crawling.
	c.cwls--

	// All crawlers has done crawling.
	if c.cwls <= 0 {
		close(c.ch)
	}
}

func main() {
	c := Crawler{
		ch: make(chan Response),
		cu: make(map[string]bool),
	}

	// Of course we won't crawl that deep in reality.
	go c.Crawl("http://golang.org/pkg/os/", 1234, fetcher)

	for res := range c.ch {
		if res.err == nil {
			fmt.Printf("found: %s %q\n", res.url, res.body)
			continue
		}
		fmt.Println(res.err)
	}
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}

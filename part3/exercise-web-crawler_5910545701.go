package main

/**
*  @author Theeruth Borisuth
 */
import (
	"fmt"
	"log"
	"sync"
)

var wg sync.WaitGroup

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

//Constructor for work & Crawl function
type Job struct {
	url   string
	depth int
}

// work function kicks off a goroutine for each unique URL it found, and the goroutine then adds URLs from that page back to the channel.
func work(fetcher Fetcher, ch chan Job, wg *sync.WaitGroup) {
	seen := make(map[string]bool)
	for job := range ch {
		if seen[job.url] || job.depth <= 0 {
			wg.Done()
			continue
		}
		seen[job.url] = true
		go Crawl(job, fetcher, ch, wg)
	}
}

func Crawl(job Job, fetcher Fetcher, q chan Job, wg *sync.WaitGroup) {
	defer wg.Done()
	// if depth <= 0 {
	// 	return
	// }
	body, urls, err := fetcher.Fetch(job.url)
	if err != nil {
		log.Println("Crawl failure:", err)
		return
	}
	log.Printf("Crawl: found %s\t%q\n", job.url, body)

	if job.depth <= 1 {
		return
	}

	wg.Add(len(urls))
	for _, u := range urls {
		q <- Job{u, job.depth - 1}
	}
}

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(1)

	q := make(chan Job)
	go work(fetcher, q, wg)
	q <- Job{"http://golang.org/", 4}
	wg.Wait()
	// close q  so that the work goroutine
	// will finish.
	close(q)
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

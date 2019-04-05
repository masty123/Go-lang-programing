// // The server program issues Google search requests. It serves on port 8080.
// //
// // The /search endpoint accepts these query params:
// //   q=the Google search query
// //
// // For example, http://localhost:8080/search?q=golang serves the first
// // few Google search results for "golang".
package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"time"

	"golang.org/x/talks/2016/applicative/google"
)

type Result struct{ Title, URL string }

type Writer interface {
	Write(p []byte) (n int, err error)
}

// handleSearch handles URLs like "/search?q=golang" by running a
// Google search for "golang" and writing the results as HTML to w.
// The query parameter "output" selects alternate output formats:
// "json" for JSON, "prettyjson" for human-readable JSON.
func handleSearch(w http.ResponseWriter, req *http.Request) { // HL
	log.Println("serving", req.URL)

	// Check the search query.
	query := req.FormValue("q") // HL
	if query == "" {
		http.Error(w, `missing "q" URL parameter`, http.StatusBadRequest)
		return
	}
	// ENDQUERY OMIT

	// Run the Google search.
	start := time.Now()
	results, err := google.Search(query) // HL
	elapsed := time.Since(start)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// ENDSEARCH OMIT

	// Create the structured response.
	type response struct {
		Results []google.Result
		Elapsed time.Duration
	}
	resp := response{results, elapsed} // HL
	// ENDRESPONSE OMIT

	// Render the response.
	switch req.FormValue("output") {
	case "json":
		err = json.NewEncoder(w).Encode(resp) // HL
	case "prettyjson":
		var b []byte
		b, err = json.MarshalIndent(resp, "", "  ") // HL
		if err == nil {
			_, err = w.Write(b)
		}
	default: // HTML
		err = responseTemplate.Execute(w, resp) // HL
	}
	// ENDRENDER OMIT
	if err != nil {
		log.Print(err)
		return
	}
}

//Search function

func Search(query string) ([]Result, error) {
	results := []Result{
		Web(query),
		Image(query),
		Video(query),
	}
	return results, nil
}

//Fake search for testing

func FakeSearch(kind, title, url string) SearchFunc {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result{
			Title: fmt.Sprintf("%s(%q): %s", kind, query, title),
			URL:   url,
		}
	}
}

//Parallel Search function

func SearchParallel(query string) ([]Result, error) {
	c := make(chan Result)
	go func() { c <- Web(query) }()
	go func() { c <- Image(query) }()
	go func() { c <- Video(query) }()

	return []Result{<-c, <-c, <-c}, nil
}

func SearchTimeout(query string, timeout time.Duration) ([]Result, error) {
	timer := time.After(timeout)
	c := make(chan Result, 3)
	go func() { c <- Web(query) }()
	go func() { c <- Image(query) }()
	go func() { c <- Video(query) }()

	var results []Result
	for i := 0; i < 3; i++ {
		select {
		case result := <-c:
			results = append(results, result)
		case <-timer:
			return results, errors.New("timed out")
		}
	}
	return results, nil
}

func First(replicas ...SearchFunc) SearchFunc {
	return func(query string) Result {
		c := make(chan Result, len(replicas))
		searchReplica := func(i int) {
			c <- replicas[i](query)
		}
		for i := range replicas {
			go searchReplica(i)
		}
		return <-c
	}
}

type SearchFunc func(query string) Result

// var (
// 	replicatedWeb   = First(Web1, Web2)
// 	replicatedImage = First(Image1, Image2)
// 	replicatedVideo = First(Video1, Video2)
// )

var (
	Web   = FakeSearch("web", "The Go Programming Language", "http://golang.org")
	Image = FakeSearch("image", "The Go gopher", "https://blog.golang.org/gopher/gopher.png")
	Video = FakeSearch("video", "Concurrency is not Parallelism", "https://www.youtube.com/watch?v=cN_DpYBzKso")
)

var responseTemplate = template.Must(template.New("results").Parse(`
<html>
<head/>
<body>
  <ol>
  {{range .Results}}
    <li>{{.Title}} - <a href="{{.URL}}">{{.URL}}</a></li>
  {{end}}
  </ol>
  <p>{{len .Results}} results in {{.Elapsed}}</p>
</body>
</html>
`))

// FakeSearch Function

//Run the program

func main() {
	//http://localhost:8080/search?q=golang
	start := time.Now()

	//All types of search function

	//results, err := google.Search("golang")
	//results, err := google.SearchParallel("golang")
	//results, err := google.SearchTimeout("golang", 80*time.Millisecond)
	results, err := google.SearchReplicated("golang", 80*time.Millisecond)

	http.HandleFunc("/search", handleSearch) // HL
	fmt.Println("serving on http://localhost:8080/search")
	elapsed := time.Since(start)
	fmt.Println(results)
	fmt.Println(elapsed, err)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))

	// start := time.Now()
	// results, err := google.Search("golang")
	// elapsed := time.Since(start)
	// fmt.Println(results)
	// fmt.Println(elapsed, err)

	//results, err = google.Search("golang")
	// results, err := google.SearchParallel("golang")
	// results, err := google.SearchTimeout("golang", 80*time.Millisecond)
	// results, err := google.SearchReplicated("golang", 80*time.Millisecond)

	//First Function

	// start := time.Now()
	// search := google.First(
	// 	google.FakeSearch("replica 1", "I'm #1!", "golang"),
	// 	google.FakeSearch("replica 2", "#2 wins!", "golang"))
	// result := search("golang")
	// elapsed := time.Since(start)
	// fmt.Println(result)
	// fmt.Println(elapsed)
}

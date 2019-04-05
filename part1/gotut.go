package main

import (
	"encoding/xml"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

// Part 16&17 HTML templates, Using template

type SitemapIndex struct {
	//Locations []string `xml:"sitemap"`
	Locations []string `xml:"sitemap>loc"`
}

type News struct {
	// Titles    []string `xml:"url>news>title"`
	// Keywords  []string `xml:"url>news>keywords"`
	// Locations []string `xml:"url>loc"`
	Titles    []string `xml:"url>image>title"`
	Keywords  []string `xml:"url>image>caption"`
	Locations []string `xml:"url>loc"`
}

type NewsMap struct {
	Keyword  string
	Location string
}

type NewsAggPage struct {
	Title string
	//News  string
	News map[string]NewsMap
}

func newsAggHandler(w http.ResponseWriter, r *http.Request) {
	var s SitemapIndex
	var n News
	resp, _ := http.Get("https://www.thairath.co.th/sitemap.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	xml.Unmarshal(bytes, &s)
	news_map := make(map[string]NewsMap)

	//for _, Location := range s.Locations {
	for i := 0; i < 3; i++ {
		resp, _ := http.Get(s.Locations[i])
		bytes, _ := ioutil.ReadAll(resp.Body)
		xml.Unmarshal(bytes, &n)

		for idx, _ := range n.Titles {
			news_map[n.Titles[idx]] = NewsMap{n.Keywords[idx], n.Locations[idx]}
		}

		// for idx, data := range news_map {
		// 	fmt.Println("\n\n\n", idx)
		// 	fmt.Println("\n", data.Keyword)
		// }
	}

	p := NewsAggPage{Title: "Amazing News Aggregator", News: news_map}
	t, _ := template.ParseFiles("newsaggtemplate.html")
	//fmt.Println(err)
	t.Execute(w, p)
}

func indexHanlder(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Whoa, Go is neat!</h1>")
}

func main() {
	http.HandleFunc("/", indexHanlder)
	http.HandleFunc("/agg/", newsAggHandler)
	http.ListenAndServe(":8000", nil)
}

// //Part 15 Mapping new data

// type SitemapIndex struct {
// 	//Locations []string `xml:"sitemap"`
// 	Locations []string `xml:"sitemap>loc"`
// }

// type News struct {
// 	Titles    []string `xml:"url>news>title"`
// 	Keywords  []string `xml:"url>news>keywords"`
// 	Locations []string `xml:"url>loc"`
// }

// type NewsMap struct {
// 	Keyword  string
// 	Location string
// }

// func main() {
// 	var s SitemapIndex
// 	var n News
// 	news_map := make(map[string]NewsMap)
// 	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
// 	bytes, _ := ioutil.ReadAll(resp.Body)
// 	xml.Unmarshal(bytes, &s)

// 	for _, Location := range s.Locations {
// 		resp, _ := http.Get(Location)
// 		bytes, _ := ioutil.ReadAll(resp.Body)
// 		xml.Unmarshal(bytes, &n)

// 		for idx, _ := range n.Titles {
// 			news_map[n.Titles[idx]] = NewsMap{n.Keywords[idx], n.Locations[idx]}
// 		}

// 		for idx, data := range news_map {
// 			fmt.Println("\n\n\n", idx)
// 			fmt.Println("\n", data.Keyword)
// 			fmt.Println("\n", data.Keyword)
// 		}
// 	}
// }

//Part 14 Map

// func main() {
// 	grades := make(map[string]float32)
// 	grades["Timmy"] = 42
// 	grades["Jess"] = 92
// 	grades["Sam"] = 67

// 	fmt.Println(grades)

// 	TimsGrade := grades["Timmy"]
// 	fmt.Println(TimsGrade)

// 	delete(grades, "Timmy")
// 	fmt.Println(grades)

// 	for k, v := range grades {
// 		fmt.Println(k, ":", v)
// 	}
// }

//Part 12&13 Looping & Web Application

// func main() {
// 	//i := 0

// 	// for i < 10 {
// 	// 	fmt.Println(i)
// 	// 	//i++
// 	// 	i += 5
// 	// }

// 	// for i := 0; i < 10; i++ {
// 	// 	fmt.Println(i)
// 	// }

// 	// for {
// 	// 	fmt.Println("Do stuff")
// 	// }

// 	// x := 5
// 	// for {
// 	// 	fmt.Println("Do stuff", x)
// 	// 	x += 3
// 	// 	if x > 25 {
// 	// 		break
// 	// 	}
// 	// }

// 	// a := 3
// 	// for x := 5; a < 25; x += 3 {
// 	// 	fmt.Println("do stuff", x)
// 	// 	a += 4
// 	// }
// 	// var s SitemapIndex
// 	// var n News
// 	// resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
// 	// //bytes, _ := washPostXML
// 	// bytes, _ := ioutil.ReadAll(resp.Body)
// 	// //string_body := string(bytes)
// 	// //fmt.Println(string_body)
// 	// //resp.Body.Close()
// 	// xml.Unmarshal(bytes, &s)
// 	// //fmt.Println(s.Locations)
// 	// for _, Location := range s.Locations {
// 	// 	//fmt.Printf("\n%s", Location)
// 	// 	resp, _ := http.Get(Location)
// 	// 	bytes, _ := ioutil.ReadAll(resp.Body)
// 	// 	xml.Unmarshal(bytes, &n)
// 	// }

// }

//Part 10&11 - XML also use in Part 12,13

// var washPostXML = []byte(`
// <sitemapindex>
//    <sitemap>
//       <loc>http://www.washingtonpost.com/news-politics-sitemap.xml</loc>
//    </sitemap>
//    <sitemap>
//       <loc>http://www.washingtonpost.com/news-blogs-politics-sitemap.xml</loc>
//    </sitemap>
//    <sitemap>
//       <loc>http://www.washingtonpost.com/news-opinions-sitemap.xml</loc>
//    </sitemap>
// </sitemapindex>`)

// type SitemapIndex struct {
// 	//Locations []string `xml:"sitemap"`
// 	Locations []string `xml:"sitemap>loc"`

// }

// type News struct {
// 	Titles []string `xml:"url>news>title"`
// 	Keywords []string `xml:"url>news>Keywords"`
// 	Locations []string `xml:"url>loc"`
// }

// type Location struct {
// 	Loc string `xml:"loc"`
// }

// func (l Location) String() string {
// 	return fmt.Sprintf(l.Loc)
// }

// [5 5]type == array
// []type == slice

// func main() {
// 	//resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
// 	bytes := washPostXML
// 	//string_body := string(bytes)
// 	//fmt.Println(string_body)
// 	//resp.Body.Close()
// 	var s SitemapIndex
// 	xml.Unmarshal(bytes, &s)
// 	//fmt.Println(s.Locations)
// 	for _, Location := range s.Locations {
// 		fmt.Printf("\n%s", Location)
// 	}
// }

//Part 1-9: Intro, Syntax, Types, Pointers, Simple Web App, Structs, Methods, Pointer Receievers, More Web Dev and Accessing the Internet in Go

// "net/http"
//math"
//"math/rand"

// func add(x, y float32) float32 {
// 	return x + y
// }

// func multiple(a, b string) (string, string) {
// 	return a, b
// }

// func index_handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Whoa, Go is neat!")
// }

// func about_handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Expert web design by Sentdex")
// }
// func (c car) kmh() float64 {
// 	//c.top_speed_kmh = 500
// 	return float64(c.gas_pedal) * (c.top_speed_kmh / usixteenbitmax)
// }

// func (c car) mph() float64 {
// 	c.top_speed_kmh = 500
// 	return float64(c.gas_pedal) * (c.top_speed_kmh / usixteenbitmax / kmh_multiple)
// }

// func (c *car) new_top_speed(newspeed float64) {
// 	c.top_speed_kmh = newspeed
// }

// func newer_top_speed(c car, speed float64) car {
// 	c.top_speed_kmh = speed
// 	return c
// }

// const usixteenbitmax float64 = 65535
// const kmh_multiple float64 = 1.60934

// type car struct {
// 	gas_pedal      uint16 //min 0 max 65535
// 	brake_pedal    uint16
// 	steering_wheel int16 // -32 - +32k
// 	top_speed_kmh  float64
// }

// func main() {
// 	// num1, num2 := 5.6, 9.5
// 	//fmt.Println(add(num1, num2))
// 	// w1, w2 := "Hey", "there"

// 	// fmt.Println(multiple(w1, w2))

// 	// var a int = 62
// 	// var b float64 = float64(a)
// 	// x := a // x will be type int
// 	// x := 15
// 	// a := &x //memory address

// 	// fmt.Println(a)
// 	// fmt.Println(*a)
// 	// *a = 5
// 	// fmt.Println(x)
// 	// *a = *a * *a
// 	// fmt.Println(x)
// 	// fmt.Println(*a)

// 	// http.HandleFunc("/", index_handler)
// 	// http.HandleFunc("/about/", about_handler)
// 	// http.ListenAndServe(":8000", nil)
// 	// a_car := car{
// 	// 	// gas_pedal:      22341,
// 	// 	gas_pedal:      65000,
// 	// 	brake_pedal:    0,
// 	// 	steering_wheel: 12561,
// 	// 	top_speed_kmh:  225.0}

// 	// fmt.Println(a_car.gas_pedal)
// 	// fmt.Println(a_car.kmh())
// 	// fmt.Println(a_car.mph())
// 	// //a_car.new_top_speed(500)
// 	// a_car = newer_top_speed(a_car, 500)
// 	// fmt.Println(a_car.kmh())
// 	// fmt.Println(a_car.mph())

// }

// func index_hanlder(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, `<h1>Hey there</h1>
// 		<p>Go is fast!</p>
// 		<p>...and simple!</p>`)

// 	// fmt.Fprintf(w, "<h1>Hey there</h1>")
// 	// fmt.Fprintf(w, "<p>Go is fast!</p>")
// 	// fmt.Fprintf(w, "<p>...and simple!</p>")
// 	// fmt.Fprintf(w, "<p>You %s even add %s</p>", "can", "<strong>variables</strong>")

// }

// func main() {
// 	http.HandleFunc("/", index_hanlder)
// 	http.ListenAndServe(":8000", nil)
// }

// func main() {
// 	fmt.Println("A number from 1-100: ", rand.Intn(100))
// }

// func foo() {
// 	fmt.Println("The square root of 4 is", math.Sqrt(4))
// }

// func main() {
// 	foo()
// }

// func main(){
// 	fmt.Println("Welcome to Go!")
// }

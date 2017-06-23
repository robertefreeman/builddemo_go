package main

import (
	"html/template"
	"log"
	"net/http"
	"sync"
	"github.com/garyburd/redigo/redis"
	"fmt"
)

type Page struct {
	Body  string
	Title string
	Name  string
}

var mu sync.Mutex

// Global pool that handlers can grab a connection from
var pool = newPool()

// Pool configuration
func newPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
	}
}
func ignoreIcon(res http.ResponseWriter, req *http.Request) {}


func serveTemplate(w http.ResponseWriter, r *http.Request) {

	tmpl := template.New("index.html") //create a new template with some name
	tmpl, _ = tmpl.ParseFiles("index.html")

	count := counter()
	bodString := fmt.Sprintf("Hello World! This page has been viewed %v times", count)
	p := Page{Body: bodString, Title: `Docker EE CI/CD demo`, Name: "Robert"}

	if err := tmpl.Execute(w, p); err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(500), 500)
	}
}

// increment counter by one and return total counter value
func counter() int {
	// Grab a connection and make sure to close it with defer
	conn := pool.Get()
	defer conn.Close()
	mu.Lock()
	conn.Do("INCR", "viewCount")
	mu.Unlock()
	count, _ := redis.Int(conn.Do("GET", "viewCount"))
	return count
}

/*
func HttpIndexFileHandler(response http.ResponseWriter, request *http.Request) {
	//fmt.Fprintf(w, "Hi from e %s!", r.URL.Path[1:])
	http.ServeFile(response, request, "example.html")
}
*/

func main() {
	http.HandleFunc("/favicon.ico", ignoreIcon)
	http.HandleFunc("/", serveTemplate)
	log.Println("Server Listening...")
	http.ListenAndServe(":8000", nil)
}
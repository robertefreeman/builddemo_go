/*
package main


import (
	"sync"
	"github.com/garyburd/redigo/redis"
	"net/http"
	"fmt"
)

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

// increment counter by one and return total counter value
func counter(res http.ResponseWriter, req *http.Request) {
	// Grab a connection and make sure to close it with defer
	conn := pool.Get()
	defer conn.Close()
	mu.Lock()
	conn.Do("INCR", "viewCount")
	mu.Unlock()
	count, _ := redis.Bytes(conn.Do("GET", "viewCount"))
	fmt.Fprintf(res, "Hello World, this page has been seen %s times \n", count )
	//res.Write(count)
}

// ignore Chrome looking for favicon.ico and incrementing counter
func ignoreIcon(res http.ResponseWriter, req *http.Request) {}


func main() {
	http.HandleFunc("/favicon.ico", ignoreIcon)
	http.HandleFunc("/", counter)
	http.ListenAndServe(":8000", nil)
}
*/

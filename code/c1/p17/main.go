package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"

	lissajous "goProgrammingLanguage/code/c1/p14"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	http.HandleFunc("/gif", liss)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "Remote Addr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Counter: %d\n", count)
	mu.Unlock()
}

func liss(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	cycles := values.Get("cycles")
	iCycles, err := strconv.Atoi(cycles)
	if err != nil {
		iCycles = 5
	}
	lissajous.Lissajous(w, iCycles)
}

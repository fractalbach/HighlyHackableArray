package main

import (
	"flag"
	"fmt"
	"github.com/fractalbach/HighlyHackableArray/hha"
	"log"
	"net/http"
)

var (
	home_page = flag.String("index", "index.html", "Home page accessible at root.")
	addr      = flag.String("a", "localhost:8080", "http service address")
	h         = hha.Create(1000)
)

func main() {
	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("/", serveIndex)
	mux.HandleFunc("/string", serveString)
	mux.HandleFunc("/ints", serveInts)
	mux.HandleFunc("/base64", servebase64)

	s := &http.Server{
		Addr:    (*addr),
		Handler: mux,
	}
	h.CopyWrite(10, []byte("This is the Highly Hackable Array!"))
	log.Printf("Listening and Serving on %v", (*addr))
	log.Fatal(s.ListenAndServe())
}

// serveIndex will serve the main page, which contains a template.
// This is the only page that has a template, and all of the others return
// unspecified and arbitrary data.
func serveIndex(w http.ResponseWriter, r *http.Request) {
	logRequest(r)
	if r.URL.Path == "/" {
		fmt.Fprint(w, "Welcome to the Highly Hackable Array! Try going to /string, /base64, or /ints.")
		return
	}
	fmt.Fprint(w, "You're traveling down a highly hackable path.")
}

// serveString will return an HTML-ish page with a pre-formatted tag.
// Without the tags, the browser will attempt to download the
// string as if it were binary data... and that probably won't end well :)
func serveString(w http.ResponseWriter, r *http.Request) {
	logRequest(r)
	fmt.Fprint(w, "<html><pre style=\"white-space: pre-wrap;\">"+h.String()+"</pre></html>")
}

func serveInts(w http.ResponseWriter, r *http.Request) {
	logRequest(r)
	fmt.Fprint(w, h.Ints())
}

func servebase64(w http.ResponseWriter, r *http.Request) {
	logRequest(r)
	fmt.Fprint(w, h.Base64())
}

// logRequest prints out a useful message to the command line log,
// displaying information about the request that was just made to the server.
func logRequest(r *http.Request) {
	log.Printf("(%v) %v %v %v", r.RemoteAddr, r.Proto, r.Method, r.URL)
}

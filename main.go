package main

import (
	"net/http"
	"flag"
	"fmt"
	"log"
	"os"
)

var port int
var host string
var path string

func main() {
	flag.IntVar(&port, "port", 3003, "Port to listen on")
	flag.StringVar(&host, "host", "", "Host.")
	flag.StringVar(&path, "path", "", "Path to serve files from. Default: current directory.")
	flag.Parse()

	http.Handle("/", http.FileServer(http.Dir(path)))

	if path == "" {
		cwd, err := os.Getwd()

		if err != nil {
			panic(err)
		}

		path = cwd
	}

	log.Printf("Listening on %s:%d and serving %s\n", host, port, path)

	http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), nil)
}

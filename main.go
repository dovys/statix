package main

import (
	"net/http"
	"flag"
	"fmt"
	"log"
	"os"
)

var port = *flag.Int("port", 3003, "Port to listen on.")
var host = *flag.String("host", "", "Host.")
var dir = *flag.String("path", "", "Path to serve files from. Default: current directory.")

func main() {
	flag.Parse()
	http.Handle("/", http.FileServer(http.Dir(dir)))

	if dir == "" {
		cwd, err := os.Getwd()

		if err != nil {
			panic(err)
		}

		dir = cwd
	}

	log.Printf("Listening on %s:%d and serving %s\n", host, port, dir)

	http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), nil)
}

package main

import (
	"net/http"
	"flag"
	"fmt"
	"log"
	"os"
	"github.com/urfave/negroni"
)

var port int
var host string
var path string

func main() {
	flag.IntVar(&port, "port", 3003, "Port to listen on")
	flag.StringVar(&host, "host", "", "Host.")
	flag.StringVar(&path, "path", "", "Path to serve files from. Default: current directory.")
	flag.Parse()

	if path == "" {
		cwd, err := os.Getwd()

		if err != nil {
			panic(err)
		}

		path = cwd
	}

	mux := http.NewServeMux()
	n := negroni.New()
	n.Use(negroni.NewStatic(http.Dir(path)))
	n.UseHandler(mux)

	log.Printf("Listening on %s:%d and serving %s\n", host, port, path)

	http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), n)
}

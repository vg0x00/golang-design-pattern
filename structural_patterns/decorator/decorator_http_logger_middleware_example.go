package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type MyServer struct {
}

func (s *MyServer) ServeHTTP(resp http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(resp, "hello decorator")
}

// extend MyServer, add logger to io.Writer
type LoggerServer struct {
	Writer  io.Writer
	Handler http.Handler
}

func (ls *LoggerServer) ServeHTTP(resp http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(ls.Writer, "(LOG) before running core http service\n")
	fmt.Fprintf(ls.Writer, "Method: %s\n", r.Method)
	fmt.Fprintf(ls.Writer, "URI: %s\n", r.URL.RequestURI())
	fmt.Fprintf(ls.Writer, "Content Length: %d\n", r.ContentLength)

	ls.Handler.ServeHTTP(resp, r)
	fmt.Fprintf(ls.Writer, "(LOG) after running core http service\n")
	fmt.Fprintf(ls.Writer, "------------------------------------\n")
}

func main() {
	http.Handle("/", &LoggerServer{Writer: os.Stdout, Handler: &MyServer{}})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

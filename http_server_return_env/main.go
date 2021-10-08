package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/emicklei/go-restful"
)

// This example shows the minimal code needed to get a restful.WebService working.
//
// GET http://localhost:8080/hello

func main() {
	ws := new(restful.WebService)
	ws.Route(ws.GET("/hello").To(hello))
	ws.Route(ws.GET("/health").To(health))
	restful.Add(ws)
	log.Fatal(http.ListenAndServe(":80", nil))
}

func hello(_ *restful.Request, resp *restful.Response) {
	response := os.Getenv("printwhat")
	_, _ = io.WriteString(resp, response)
}

func health(_ *restful.Request, resp *restful.Response) {
	_, _ = io.WriteString(resp, "healthy")
}

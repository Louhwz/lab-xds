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

//func init()  {
//	err := os.Setenv("lightest", "yes")
//	if err != nil {
//		panic("here")
//	}
//}

func main() {
	ws := new(restful.WebService)
	ws.Route(ws.GET("/hello").To(hello))
	restful.Add(ws)

	weightLightest := os.Getenv("lightest")
	port := ""
	if weightLightest != "" {
		port = ":18081"
	} else {
		port = ":18080"
	}
	log.Printf("server start in %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func hello(req *restful.Request, resp *restful.Response) {
	weightLightest := os.Getenv("lightest")
	response := ""
	if weightLightest != "" {
		response = "命中小权重容器"
		log.Println(response)
	} else {
		response = "命中大权重容器"
		log.Println(response)
	}
	_, _ = io.WriteString(resp, response)
}

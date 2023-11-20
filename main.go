package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(w, "Hello, Kubernetes!")
		log.Println("connection from", req.RemoteAddr)

		names := make([]string, 0, len(req.Header))
		for k := range req.Header {
			names = append(names, k)
		}
		for _, name := range names {
			log.Printf("%s: %q\n", name, req.Header[name])
		}
		fmt.Fprintln(w, req.Header)
		log.Println()
		log.Println()
	})

	http.HandleFunc("/lang/zh-cn", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(w, "你好 kubernetes http/1.1")
		log.Println("connection from", req.RemoteAddr)

		names := make([]string, 0, len(req.Header))
		for k := range req.Header {
			names = append(names, k)
		}
		for _, name := range names {
			log.Printf("%s: %q\n", name, req.Header[name])
		}
		fmt.Fprintln(w, req.Header)
		log.Println()
		log.Println()
	})

	http.ListenAndServe(":8080", nil)
}

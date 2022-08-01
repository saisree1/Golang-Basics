package main

import (
	"fmt"
	"log"
	"net/http"
)

// func main() {
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		w.Write([]byte("Hello world"))
// 	})
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }

type Handler struct {
}

func main() {
	h := Handler{}
	http.Handle("/", h)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Context())
	fmt.Println(r.Body)
	fmt.Println(r.URL.Query())
	w.Write([]byte("Hello World"))
}

// func main() {
// 	h := Handler{}
// 	http.HandleFunc("/basic", h.HandleGet())
// 	http.ListenAndServe(":8080", nil)
// }

// func (h Handler) HandleGet() func(w http.ResponseWriter, r *http.Request) {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		w.Write([]byte("hello world"))
// 	}
// }

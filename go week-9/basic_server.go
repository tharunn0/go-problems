package main

import (
	"fmt"
	"log"
	"net/http"
)

func hellohandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "From hello handler")
}
func newpagehandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "From New Page handler")
}
func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hellohandler)
	mux.HandleFunc("/newpage", newpagehandler)

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintln(w, "Basic Go Server")
	// 	log.Println(r.URL.String())
	// })

	serverAdd := "127.0.0.1:8080"
	err := http.ListenAndServe(serverAdd, mux)
	if err != nil {
		log.Fatalln("Error running server", err)
	}

}

package main

import (
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./public"))
	mux := http.NewServeMux()
	mux.Handle("/", fileServer)
	mux.HandleFunc("/blog", helloFunc)
	log.Fatal(http.ListenAndServe(":3333", mux))
}

func helloFunc(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Hello From Blog"))
}

package main

import (
	"fmt"
	"net/http"

	"github.com/CourtCourt521/urlshort/handler"
)

func main() {
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/CourtCourt521/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	fmt.Println("Call mapHandler")
	mapHandler := handler.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	//yaml := `
	//- path: /urlshort
	//  url: https://github.com/gophercises/urlshort
	//- path: /urlshort-final
	//  url: https://github.com/gophercises/urlshort/tree/solution
	// `
	//yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler)
	//if err != nil {
	//	panic(err)
	//}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", mapHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

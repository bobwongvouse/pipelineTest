package main

import("fmt")
import("net/http")

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	fmt.Println("Running.")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

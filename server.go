package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./public")))

	http.HandleFunc("/comments.json", func(w http.ResponseWriter, r *http.Request) {
		file, err := os.Open("_comments.json")
		if err != nil {
			panic(err)
		}
		w.Header().Set("Content-Type", "application/json")
		io.Copy(w, file)
	})

	http.ListenAndServe(":3000", nil)
}

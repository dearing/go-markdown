package main

import (
	"net/http"
	"os"

	"github.com/dearing/go-markdown"
)

func main() {
	// This is the main function

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		content, err := os.ReadFile("test.md")
		if err != nil {
			http.Error(w, "Error reading file", http.StatusInternalServerError)
			return
		}

		md := markdown.NewMarkdown(string(content))
		html := md.ToHTML()

		w.Write([]byte(html))

	})

	http.ListenAndServe(":8080", nil)

}

package main

import (
	"fmt"
	"net/http"

	"github.com/tzaffi/urlshort"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, world!")
	})

	// Hard coded:
	pathsToUrls := map[string]string{
		"/zblog":    "http://grunschblog.blogspot.com",
		"/zwitter":  "https://twitter.com/tzaffi",
		"/zinkedin": "https://www.linkedin.com/in/zephgrunschlag",
		"/zithub":   "https://github.com/tzaffi",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	yaml := `
- path: /vcgotags
  url: https://www.reddit.com/r/golang/comments/hpaw4b/how_to_quickly_create_struct_fields_and_their
- path: /zblog
  url: http://grunschblog.blogspot.com
- path: /zwitter
  url: https://twitter.com/tzaffi
- path: /zinkedin
  url: https://www.linkedin.com/in/zephgrunschlag
- path: /zithub
  url: https://github.com/tzaffi
`
	yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}

	port := ":8080"
	fmt.Printf("Starting the server on :%s\n", port)
	http.ListenAndServe(port, yamlHandler)
}

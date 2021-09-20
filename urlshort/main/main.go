package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/tzaffi/gophercises/urlshort"
)

func main() {
	os.Setenv("DATABASE_URL", "postgres://root:root@localhost:5432/test_db?sslmode=disable")

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
- path: /zblog2
  url: http://grunschblog.blogspot.com
- path: /zwitter2
  url: https://twitter.com/tzaffi
- path: /zinkedin2
  url: https://www.linkedin.com/in/zephgrunschlag
- path: /zithub2
  url: https://github.com/tzaffi
`

	yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}

	dbHandler, err := urlshort.DBHandler(yamlHandler)
	if err != nil {
		panic(err)
	}

	port := ":8080"
	fmt.Printf("Starting the server on :%s\n", port)
	http.ListenAndServe(port, dbHandler)
}

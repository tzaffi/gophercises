package urlshort

import (
	"net/http"

	yaml "gopkg.in/yaml.v2"
)

// import "gopkg.in/yaml.v2"

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if dest, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}

func YAMLHandler(yBytes []byte, fallback http.HandlerFunc) (http.HandlerFunc, error) {
	// 1. parse the yaml
	var pathUrls []pathUrl
	err := yaml.Unmarshal(yBytes, pathUrls)
	if err != nil {
		return nil, err
	}
	// 2. convert YAML array into map
	pathsToUrls := map[string]string{}
	for _, purl := range pathUrls {
		pathsToUrls[purl.Path] = purl.URL
	}

	// 3. return a map handler using the mapping
	return MapHandler(pathsToUrls, fallback), nil
}

type pathUrl struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

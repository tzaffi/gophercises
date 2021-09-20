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
	pathUrls, err := parseUrls(yBytes)
	if err != nil {
		return nil, err
	}
	// 2. convert YAML array into map
	p2url := buildUrlMap(pathUrls)

	// 3. return a map handler using the mapping
	return MapHandler(p2url, fallback), nil
}

type pathUrl struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

func parseUrls(yBytes []byte) ([]pathUrl, error) {
	var pathUrls []pathUrl
	err := yaml.Unmarshal(yBytes, &pathUrls)
	if err != nil {
		return nil, err
	}
	return pathUrls, err
}

func buildUrlMap(purls []pathUrl) map[string]string {
	pathsToUrls := map[string]string{}
	for _, purl := range purls {
		pathsToUrls[purl.Path] = purl.URL
	}
	return pathsToUrls
}

package util

import (
	u "net/url"
	p "path"
	"strings"
)

func JoinPaths(url string, paths ...string) string {
	uri, err := u.Parse(url)
	if err != nil {
		scheme := "http"
		if strings.Contains(url, "https") {
			scheme = "https"
		}

		uri = &u.URL{
			Scheme: scheme,
			Host:   strings.TrimLeft(url, scheme+"://"),
		}
	}

	uri.Path = p.Join(insertFirstPos(uri.Path, paths)...)
	return uri.String()
}

func insertFirstPos(elem string, stringSlice []string) []string {
	return append([]string{elem}, stringSlice...)
}

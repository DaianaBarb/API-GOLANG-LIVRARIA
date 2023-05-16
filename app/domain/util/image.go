package util

func NormalizeURL(url, path string) string {
	if path != "" {
		return JoinPaths(url, path)
	}

	return ""
}

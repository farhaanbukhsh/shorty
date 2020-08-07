package helper

import (
	"strings"
)

// URLSanitizer is use to sanitize a given URL with trailing slash
// and protocol attached this is very much needed for proper redirection
func URLSanitizer(url string) string {
	cleanURL := url
	hasHTTP := strings.HasPrefix(url, "http")
	hasTraillingSlash := strings.HasSuffix(url, "/")

	if !hasHTTP {
		cleanURL = "http://" + cleanURL
	}
	if !hasTraillingSlash {
		cleanURL = cleanURL + "/"
	}
	return cleanURL
}

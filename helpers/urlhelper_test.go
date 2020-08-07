package helper

import (
	"strings"
	"testing"
)

func TestURLSantizerWithoutTrailingSlash(t *testing.T) {
	url := "http://google.com"
	sanitizeURL := URLSanitizer(url)
	if url == sanitizeURL {
		t.Errorf("Same url returned without trailling slash")
	}
	hasSlash := strings.HasSuffix(sanitizeURL, "/")
	if !hasSlash {
		t.Errorf("Doesn't have trailling slash")
	}
}

func TestURLSantizerWithoutProtocol(t *testing.T) {
	url := "google.com"
	sanitizeURL := URLSanitizer(url)
	if url == sanitizeURL {
		t.Errorf("Same url returned without trailling slash")
	}
	hasHTTP := strings.HasPrefix(sanitizeURL, "http")
	if !hasHTTP {
		t.Errorf("Didn't attach protocol to it")
	}
}

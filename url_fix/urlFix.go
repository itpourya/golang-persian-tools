package urlfix

import (
	"net/url"
	"strings"
)

func UrlFix(link string, separator string) string {
	// Decode the URL
	decodedURL, err := url.QueryUnescape(string(link))
	if err != nil {
		return ""
	}

	// If a separator is provided, replace spaces with the separator
	if separator != "" {
		return strings.ReplaceAll(decodedURL, " ", string(separator))
	}

	return decodedURL
}

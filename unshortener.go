package unshortener

import (
	"net/http"
)

// UnshortURL resolves shortened url and gives you the original one
func UnshortURL(url string) (unshortened string, err error) {

	resp, err := http.Head(url)

	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	unshortened = resp.Request.URL.String()
	return
}

package unshortener_test

import (
	"testing"

	"gitlab.com/emaele/go-unshortener"
)

func TestUnshortURL(t *testing.T) {
	var URL = "https://bit.ly/2JxxSnm"

	unshortened, err := unshortener.UnshortURL(URL)
	if unshortened == URL || err != nil {
		t.Error("Expected a different url, got same URL or err:", err)
	}
}

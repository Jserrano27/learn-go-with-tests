package concurrency

import (
	"reflect"
	"testing"
)

func mockWebsiteChecker(url string) bool {
	if url == "happy://nota.validurl" {
		return false
	}
	return true
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"https://google.com",
		"https://udemy.com",
		"happy://nota.validurl",
	}

	expected := map[string]bool{
		"https://google.com":    true,
		"https://udemy.com":     true,
		"happy://nota.validurl": false,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected: %v, but got %v instead", expected, got)
	}
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 1000)

	for i := 0; i < 100; i++ {
		urls[i] = "a url"
	}

	for i := 0; i < b.N; i++ {
		CheckWebsites(mockWebsiteChecker, urls)
	}
}

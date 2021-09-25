package elementbyid

import (
	"net/http"
	"testing"

	"golang.org/x/net/html"
)

func TestElementById(t *testing.T) {
	var tests = []struct {
		id   string
		want bool
	}{
		{"page", true},
		{"balogna", false},
		{"aaabbbcc", false},
		{"nav", true},
	}

	resp, err := http.Get("https://golang.org")
	if err != nil {
		t.Errorf("error fetching document: %s", err)
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		t.Errorf("error parsing document: %s", err)
	}

	for _, test := range tests {
		if _, found := ElementById(doc, test.id); found != test.want {
			t.Errorf("ElementById(%q) = %v", test.id, found)
		}
	}
}

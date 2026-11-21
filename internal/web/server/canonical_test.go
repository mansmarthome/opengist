package server

import (
	"net/url"
	"testing"
)

func TestCanonicalURL(t *testing.T) {
	base := "https://snippets.example"
	mustURL := func(raw string) *url.URL {
		t.Helper()
		u, err := url.Parse(raw)
		if err != nil {
			t.Fatal(err)
		}
		return u
	}

	tests := []struct {
		name string
		raw  string
		want string
	}{
		{"path only", "/roland/abc", base + "/roland/abc"},
		{"drops page", "/-/all?page=2", base + "/-/all"},
		{"drops default sort and order", "/-/all?sort=created&order=desc", base + "/-/all"},
		{"keeps non-default sort", "/-/all?sort=updated&order=desc", base + "/-/all?sort=updated"},
		{"keeps search query", "/-/search?q=homelab&page=3", base + "/-/search?q=homelab"},
		{"drops unknown params", "/roland/abc?utm_source=x&fbclid=1", base + "/roland/abc"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := canonicalURL(base, mustURL(tt.raw))
			if got != tt.want {
				t.Fatalf("canonicalURL(%q) = %q, want %q", tt.raw, got, tt.want)
			}
		})
	}

	if got := canonicalURL(base, nil); got != "" {
		t.Fatalf("nil url = %q, want empty", got)
	}
}

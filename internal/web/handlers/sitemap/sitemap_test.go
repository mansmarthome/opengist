package sitemap

import (
	"strings"
	"testing"
	"time"

	"github.com/thomiceli/opengist/internal/db"
)

func TestNewestPublicGistUpdate(t *testing.T) {
	older := time.Date(2024, 1, 2, 3, 4, 5, 0, time.UTC)
	newer := time.Date(2026, 6, 7, 8, 9, 10, 0, time.UTC)

	if got := newestPublicGistUpdate(nil); !got.IsZero() {
		t.Fatalf("empty list = %v, want zero", got)
	}

	got := newestPublicGistUpdate([]*db.Gist{
		{UpdatedAt: older.Unix()},
		nil,
		{UpdatedAt: 0},
		{UpdatedAt: newer.Unix()},
	})
	if !got.Equal(newer) {
		t.Fatalf("newest = %v, want %v", got, newer)
	}
}

func TestAddURLOmitsZeroLastmod(t *testing.T) {
	var sb strings.Builder
	addURL(&sb, "https://snippets.example/-/all", time.Time{}, "weekly", "0.8")
	if strings.Contains(sb.String(), "<lastmod>") {
		t.Fatalf("zero lastmod was emitted: %s", sb.String())
	}

	sb.Reset()
	updated := time.Date(2026, 6, 7, 8, 9, 10, 0, time.FixedZone("MSK", 3*60*60))
	addURL(&sb, "https://snippets.example/-/all", updated, "weekly", "0.8")
	if !strings.Contains(sb.String(), "<lastmod>2026-06-07T05:09:10Z</lastmod>") {
		t.Fatalf("lastmod was not UTC: %s", sb.String())
	}
}

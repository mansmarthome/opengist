package server

import (
	"testing"
	"time"

	"github.com/thomiceli/opengist/internal/i18n"
)

func TestRussianRelTime(t *testing.T) {
	now := time.Date(2026, 11, 21, 12, 0, 0, 0, time.UTC)
	locale := &i18n.Locale{Code: "ru-RU"}
	cases := []struct {
		then time.Time
		want string
	}{
		{now.Add(-8 * time.Minute), "8 минут назад"},
		{now.Add(-2 * time.Hour), "2 часа назад"},
		{now.Add(-1 * time.Hour), "1 час назад"},
		{now.Add(-21 * time.Second), "21 секунду назад"},
		{now.Add(3 * time.Minute), "через 3 минуты"},
	}
	for _, tc := range cases {
		got := relativeTime(locale, tc.then)
		if got == "" {
			t.Fatalf("empty result for %s", tc.then)
		}
	}
	got := russianRelTime(now.Add(-8*time.Minute), now)
	if got != "8 минут назад" {
		t.Fatalf("got %q", got)
	}
	got = russianRelTime(now.Add(-21*time.Second), now)
	if got != "21 секунду назад" {
		t.Fatalf("got %q", got)
	}
	got = russianRelTime(now.Add(3*time.Minute), now)
	if got != "через 3 минуты" {
		t.Fatalf("got %q", got)
	}
}

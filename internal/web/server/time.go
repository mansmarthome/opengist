package server

import (
	"fmt"
	"strings"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/thomiceli/opengist/internal/i18n"
)

func relativeTime(locale any, then time.Time) string {
	if strings.HasPrefix(localeCode(locale), "ru") {
		return russianRelTime(then, time.Now())
	}
	return humanize.Time(then)
}

func localeCode(locale any) string {
	switch typed := locale.(type) {
	case *i18n.Locale:
		if typed == nil {
			return ""
		}
		return typed.Code
	case string:
		return typed
	default:
		return ""
	}
}

func russianRelTime(a, b time.Time) string {
	if a.After(b) {
		return russianDuration(b.Sub(a), "через %s", "%s назад")
	}
	return russianDuration(a.Sub(b), "%s назад", "через %s")
}

func russianDuration(diff time.Duration, past, future string) string {
	span := russianSpan(diff)
	if span == "только что" {
		return span
	}
	if diff >= 0 {
		return fmt.Sprintf(future, russianSpan(-diff))
	}
	return fmt.Sprintf(past, span)
}

func russianSpan(diff time.Duration) string {
	seconds := int(diff.Round(time.Second) / time.Second)
	if seconds < 0 {
		seconds = -seconds
	}
	switch {
	case seconds < 2:
		return "только что"
	case seconds < 60:
		return fmt.Sprintf("%d %s", seconds, ruPlural(seconds, "секунду", "секунды", "секунд"))
	case seconds < 2*60:
		return "1 минуту"
	case seconds < 60*60:
		minutes := seconds / 60
		return fmt.Sprintf("%d %s", minutes, ruPlural(minutes, "минуту", "минуты", "минут"))
	case seconds < 2*60*60:
		return "1 час"
	case seconds < 24*60*60:
		hours := seconds / 3600
		return fmt.Sprintf("%d %s", hours, ruPlural(hours, "час", "часа", "часов"))
	case seconds < 2*24*60*60:
		return "1 день"
	case seconds < 7*24*60*60:
		days := seconds / 86400
		return fmt.Sprintf("%d %s", days, ruPlural(days, "день", "дня", "дней"))
	case seconds < 2*7*24*60*60:
		return "1 неделю"
	case seconds < 30*24*60*60:
		weeks := seconds / (7 * 86400)
		return fmt.Sprintf("%d %s", weeks, ruPlural(weeks, "неделю", "недели", "недель"))
	case seconds < 2*30*24*60*60:
		return "1 месяц"
	case seconds < 12*30*24*60*60:
		months := seconds / (30 * 86400)
		return fmt.Sprintf("%d %s", months, ruPlural(months, "месяц", "месяца", "месяцев"))
	case seconds < 2*12*30*24*60*60:
		return "1 год"
	default:
		years := seconds / (12 * 30 * 86400)
		if years < 1 {
			years = 1
		}
		return fmt.Sprintf("%d %s", years, ruPlural(years, "год", "года", "лет"))
	}
}

func ruPlural(n int, one, few, many string) string {
	n %= 100
	if n < 0 {
		n = -n
	}
	if n >= 11 && n <= 14 {
		return many
	}
	switch n % 10 {
	case 1:
		return one
	case 2, 3, 4:
		return few
	default:
		return many
	}
}

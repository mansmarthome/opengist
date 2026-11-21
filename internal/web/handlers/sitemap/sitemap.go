package sitemap

import (
	"html/template"
	"strings"
	"time"

	"github.com/thomiceli/opengist/internal/db"
	"github.com/thomiceli/opengist/internal/web/context"
)

func Sitemap(ctx *context.Context) error {
	gists, err := db.GetAllPublicGists()
	if err != nil {
		return ctx.ErrorRes(500, "Cannot fetch public gists for sitemap", err)
	}

	var sb strings.Builder
	sb.WriteString(`<?xml version="1.0" encoding="UTF-8"?>`)
	sb.WriteString(`<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">`)

	base := ctx.GetData("baseHttpUrl").(string)

	// Listing lastmod is the newest public gist, not the time the sitemap was fetched.
	addURL(&sb, base+"/-/all", newestPublicGistUpdate(gists), "weekly", "0.8")

	for _, gist := range gists {
		url := base + "/" + gist.User.Username + "/" + gist.Identifier()
		addURL(&sb, url, time.Unix(gist.UpdatedAt, 0), "monthly", "1.0")
	}

	sb.WriteString(`</urlset>`)

	ctx.Response().Header().Set("Content-Type", "application/xml")
	return ctx.PlainText(200, sb.String())
}

func newestPublicGistUpdate(gists []*db.Gist) time.Time {
	var latest time.Time
	for _, gist := range gists {
		if gist == nil || gist.UpdatedAt == 0 {
			continue
		}
		updated := time.Unix(gist.UpdatedAt, 0).UTC()
		if updated.After(latest) {
			latest = updated
		}
	}
	return latest
}

func addURL(sb *strings.Builder, loc string, lastmod time.Time, changefreq string, priority string) {
	sb.WriteString("<url>")
	sb.WriteString("<loc>" + template.HTMLEscapeString(loc) + "</loc>")
	if !lastmod.IsZero() {
		sb.WriteString("<lastmod>" + lastmod.UTC().Format(time.RFC3339) + "</lastmod>")
	}
	sb.WriteString("<changefreq>" + changefreq + "</changefreq>")
	sb.WriteString("<priority>" + priority + "</priority>")
	sb.WriteString("</url>")
}

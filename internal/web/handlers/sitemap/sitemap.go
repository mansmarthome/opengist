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

	// Add static pages
	addURL(&sb, ctx.GetData("baseHttpUrl").(string)+"/all", time.Now(), "weekly", "0.8")

	// Add public gists
	for _, gist := range gists {
		url := ctx.GetData("baseHttpUrl").(string) + "/" + gist.User.Username + "/" + gist.Identifier()
		addURL(&sb, url, time.Unix(gist.UpdatedAt, 0), "monthly", "1.0")
	}

	sb.WriteString(`</urlset>`)

	ctx.Response().Header().Set("Content-Type", "application/xml")
	return ctx.PlainText(200, sb.String())
}

func addURL(sb *strings.Builder, loc string, lastmod time.Time, changefreq string, priority string) {
	sb.WriteString("<url>")
	sb.WriteString("<loc>" + template.HTMLEscapeString(loc) + "</loc>")
	if !lastmod.IsZero() {
		sb.WriteString("<lastmod>" + lastmod.Format("2006-01-02T15:04:05Z") + "</lastmod>")
	}
	sb.WriteString("<changefreq>" + changefreq + "</changefreq>")
	sb.WriteString("<priority>" + priority + "</priority>")
	sb.WriteString("</url>")
 }

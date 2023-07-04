package middlewares

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type HSTSConfig struct {
	// MaxAge is the time, in seconds, that the browser should remember
	// that a site is only to be accessed using HTTPS.
	MaxAge int

	// If this optional parameter is specified, this rule applies to
	// all of the site's subdomains as well.
	IncludeSubdomains bool

	// Whether the domain should be listed in the preload list by Google.
	Preload bool
}

func ForceHTTPS() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		if ctx.Get("X-Forwarded-Proto") != "https" && ctx.Protocol() == "http" {
			return ctx.Redirect("https://"+ctx.Hostname()+ctx.OriginalURL(), 308)
		}
		return ctx.Next()
	}
}

func HSTS(config *HSTSConfig) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		if config.MaxAge > 0 {
			hsts := "max-age=" + strconv.Itoa(config.MaxAge)
			if config.IncludeSubdomains {
				hsts += "; includeSubDomains"
			}
			if config.Preload {
				hsts += "; preload"
			}
			ctx.Set("Strict-Transport-Security", hsts)
		}
		return ctx.Next()
	}
}

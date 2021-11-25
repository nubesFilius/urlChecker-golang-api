package app

import (
	"github.com/nubesFilius/urlChecker-golang-api/controllers/health"
	"github.com/nubesFilius/urlChecker-golang-api/controllers/url"
)

func routes() {
	router.GET("/health", health.HealthController.Check)

	router.GET("/urlinfo/1/:url", url.UrlController.Check)
}

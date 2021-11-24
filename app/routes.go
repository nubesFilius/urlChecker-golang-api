package app

import "github.com/nubesFilius/urlChecker-golang-api/controllers/health"

func routes() {
	router.GET("/health", health.Check)
	
	router.GET("/health", health.Check)
	router.GET("/health", health.Check)
	router.GET("/health", health.Check)
}

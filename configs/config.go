package configs

import (
	"os"

	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
)

var CorsConfig = cors.Config{
	AllowOriginsFunc: func(origin string) bool {
		return os.Getenv("ENV") == "development"
	},
	AllowHeaders: "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
	// AllowOrigins:     "*",
	AllowCredentials: true,
	AllowMethods:     "GET,POST,HEAD,DELETE,OPTIONS",
}

// var csrfConfig = csrf.Config{
// 	KeyLookup:      "header:X-Csrf-Token",
// 	CookieName:     "csrf_",
// 	CookieSameSite: "Lax",
// 	Expiration:     1 * time.Hour,
// 	KeyGenerator:   utils.UUID,
// }

var HelmetConfig = helmet.Config{
	// Real-world values for all headers
	XSSProtection:             "0",
	ContentTypeNosniff:        "nosniff",
	XFrameOptions:             "SAMEORIGIN",
	HSTSExcludeSubdomains:     false,
	ContentSecurityPolicy:     "default-src 'self';base-uri 'self';font-src 'self' https: data:;form-action 'self';frame-ancestors 'self';img-src 'self' data:;object-src 'none';script-src 'self';script-src-attr 'none';style-src 'self' https: 'unsafe-inline';upgrade-insecure-requests",
	CSPReportOnly:             false,
	HSTSPreloadEnabled:        true,
	ReferrerPolicy:            "no-referrer",
	PermissionPolicy:          "geolocation=(self)",
	CrossOriginEmbedderPolicy: "require-corp",
	CrossOriginOpenerPolicy:   "same-origin",
	CrossOriginResourcePolicy: "same-origin",
	OriginAgentCluster:        "?1",
	XDNSPrefetchControl:       "off",
	XDownloadOptions:          "noopen",
	XPermittedCrossDomain:     "none",
}

var CompressConfig = compress.Config{
	Level: 1,
}

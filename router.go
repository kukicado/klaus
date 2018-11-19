package klaus

import (
	"net/http"
)

// Route ...
type Route struct {
	Method     string
	Path       string
	Middleware []http.Handler
	Handler    http.HandlerFunc
}

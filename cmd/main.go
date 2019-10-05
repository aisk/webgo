package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/bnkamalesh/webgo/middleware"

	"github.com/bnkamalesh/webgo"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	// WebGo context
	wctx := webgo.Context(r)
	// URI paramaters, map[string]string
	params := wctx.Params
	// route, the webgo.Route which is executing this request
	route := wctx.Route
	webgo.R200(
		w,
		fmt.Sprintf(
			"Route name: '%s', params: '%s'",
			route.Name,
			params,
		),
	)
}

func getRoutes() []*webgo.Route {
	return []*webgo.Route{
		&webgo.Route{
			Name:          "root",                         // A label for the API/URI, this is not used anywhere.
			Method:        http.MethodGet,                 // request type
			Pattern:       "/",                            // Pattern for the route
			Handlers:      []http.HandlerFunc{helloWorld}, // route handler
			TrailingSlash: true,
		},
		&webgo.Route{
			Name:          "matchall",                     // A label for the API/URI, this is not used anywhere.
			Method:        http.MethodGet,                 // request type
			Pattern:       "/matchall/:wildcard*",         // Pattern for the route
			Handlers:      []http.HandlerFunc{helloWorld}, // route handler
			TrailingSlash: true,
		},
		&webgo.Route{
			Name:                    "api",                                             // A label for the API/URI, this is not used anywhere.
			Method:                  http.MethodGet,                                    // request type
			Pattern:                 "/api/:param",                                     // Pattern for the route
			Handlers:                []http.HandlerFunc{middleware.Cors(), helloWorld}, // route handler
			TrailingSlash:           true,
			FallThroughPostResponse: true,
		},
	}
}

func main() {
	cfg := &webgo.Config{
		Host:         "",
		Port:         "8080",
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 60 * time.Second,
	}
	router := webgo.NewRouter(cfg, getRoutes())
	router.Use(middleware.AccessLog)
	router.Start()
}
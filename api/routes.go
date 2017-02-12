package api

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Setup",
		"POST",
		"/setup",
		Setup,
	},
	Route{
		"Deploy",
		"POST",
		"/deploy",
		Deploy,
	},
	Route{
		"Undeploy",
		"POST",
		"/undeploy",
		Undeploy,
	},
}

package main

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
		"Index",
		"GET",
		"/",
		IndexV1,
	},
	Route{
		"DatabaseConnectionIndex",
		"GET",
		"/v1/databaseconnections",
		DatabaseConnectionIndexV1,
	},
	Route{
		"GetAuthenticationKey",
		"GET",
		"/v1/getauthenticationkey",
		GetAuthenticationKeyV1,
	},
	Route{
		"DatabaseConnectionCreate",
		"POST",
		"/v1/databaseconnections",
		DatabaseConnectionCreateV1,
	},
	Route{
		"DatabaseConnectionShow",
		"GET",
		"/v1/databaseconnections/{databaseConnectionId}",
		DatabaseConnectionShowV1,
	},
}

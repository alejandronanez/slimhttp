package main

import (
	"github.com/Sirupsen/logrus"

	"github.com/sdwolfe32/slimhttp"
)

func main() {
	// Create a new router
	r := slimhttp.NewRouter()

	logger := logrus.New()
	s := NewHelloService(logger)
	h := slimhttp.NewHealthcheckService(logger, "api.example.com")

	// Bind an Endpoint to the router at the specified path
	r.HandleJSONEndpoint("/healtcheck", h.Healthcheck)
	r.HandleJSONEndpoint("/hello/{name}/", s.Hello)

	// Start the service!
	r.ListenAndServe("8080")
}

package server

import (
	"log"
	"net/http"

	"github.com/erik-sostenes/receipt-processor-api/pkg/common"
	"github.com/erik-sostenes/receipt-processor-api/pkg/server/routes"
)

const defaultPort = "8080"

type (
	// Server contains all the settings for the server
	Server struct {
		*http.Server
	}
)

func New(groups ...routes.RouteGroup) *Server {
	routes := make(routes.RouteCollection, len(groups))

	for _, group := range groups {
		for key, value := range group.RouteCollection {
			routes[key] = value
		}
	}

	return &Server{
		&http.Server{
			Handler: &routes,
		},
	}
}

func (s *Server) Start() error {
	port := common.GetEnv("SERVER_PORT")
	if port == "" {
		port = defaultPort
	}

	log.Printf("server is running on port '%s'\n", port)
	return http.ListenAndServe(":"+port, s.Handler)
}

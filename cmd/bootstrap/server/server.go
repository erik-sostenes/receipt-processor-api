package server

import (
	"github.com/erik-sostenes/receipt-processor-api/pkg/server/health"
	m "github.com/erik-sostenes/receipt-processor-api/pkg/server/middlewares"
	"github.com/erik-sostenes/receipt-processor-api/pkg/server/routes"
)

func Injector() (*routes.RouteGroup, error) {
	routes := routes.NewGroup("/api/v1/receipts/", m.CORS, m.Logger, m.Recovery)
	routes.GET("/health", health.HealthCheck())

	return routes, nil
}

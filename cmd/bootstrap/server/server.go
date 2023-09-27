package server

import (
	"github.com/erik-sostenes/receipt-processor-api/internal/backoffice/receipts/business/services"
	"github.com/erik-sostenes/receipt-processor-api/internal/backoffice/receipts/infrastructure/driven/memory"
	"github.com/erik-sostenes/receipt-processor-api/internal/backoffice/receipts/infrastructure/drives/handlers"
	"github.com/erik-sostenes/receipt-processor-api/pkg/server/health"
	m "github.com/erik-sostenes/receipt-processor-api/pkg/server/middlewares"
	"github.com/erik-sostenes/receipt-processor-api/pkg/server/routes"
)

func Injector() (*routes.RouteGroup, error) {
	services := *services.NewReciptCreator(memory.NewReciptInMemory())

	routes := routes.NewGroup("/api/v1/receipts", m.CORS, m.Logger, m.Recovery)
	routes.GET("/health", health.HealthCheck())
	routes.POST("/process", handlers.HttpHandlerReceiptsCreator(services))

	return routes, nil
}

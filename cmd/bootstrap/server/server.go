package server

import (
	"github.com/erik-sostenes/receipt-processor-api/internal/backoffice/receipts/business/services"
	"github.com/erik-sostenes/receipt-processor-api/internal/backoffice/receipts/infrastructure/driven/memory"
	"github.com/erik-sostenes/receipt-processor-api/internal/backoffice/receipts/infrastructure/driven/mongo"
	"github.com/erik-sostenes/receipt-processor-api/internal/backoffice/receipts/infrastructure/drives/handlers"
	connection "github.com/erik-sostenes/receipt-processor-api/pkg/db/mongo"
	"github.com/erik-sostenes/receipt-processor-api/pkg/server/health"
	m "github.com/erik-sostenes/receipt-processor-api/pkg/server/middlewares"
	"github.com/erik-sostenes/receipt-processor-api/pkg/server/routes"
)

func Injector() (*routes.RouteGroup, error) {
	factory := connection.MongoClientFactory{}

	db, err := factory.CreateClient("mongodb://root:password@localhost:27017", "receipts_processor")
	if err != nil {
		return nil, err
	}

	collection := db.Collection("receipts")

	mongoRecepitRepository := mongo.NewMongoReceiptRepository(collection)
	receiptCreator := services.NewReciptCreator(mongoRecepitRepository)

	receiptSearcher := services.NewReceiptSearcher(memory.NewReciptInMemory())

	routes := routes.NewGroup("/api/v1/receipts/", m.CORS, m.Logger, m.Recovery)
	routes.GET("health", health.HealthCheck())
	routes.POST("process", handlers.HttpHandlerReceiptsCreator(receiptCreator))
	routes.GET("{id}/points", handlers.HttpHandlerReceiptsSearcher(receiptSearcher))

	return routes, nil
}

package server

import (
	"github.com/erik-sostenes/receipt-processor-api/internal/backoffice/receipts/business/services"
	"github.com/erik-sostenes/receipt-processor-api/internal/backoffice/receipts/infrastructure/driven/mongo"
	"github.com/erik-sostenes/receipt-processor-api/internal/backoffice/receipts/infrastructure/drives/handlers"
	"github.com/erik-sostenes/receipt-processor-api/pkg/common"
	connection "github.com/erik-sostenes/receipt-processor-api/pkg/db/mongo"
	"github.com/erik-sostenes/receipt-processor-api/pkg/server/health"
	m "github.com/erik-sostenes/receipt-processor-api/pkg/server/middlewares"
	"github.com/erik-sostenes/receipt-processor-api/pkg/server/routes"
)

func Injector() (*routes.RouteGroup, error) {
	factory := connection.MongoClientFactory{}

	db, err := factory.CreateClient(common.GetEnv("MONGO_DSN"), common.GetEnv("MONGO_DB"))
	if err != nil {
		return nil, err
	}

	collection := db.Collection("receipts")

	mongoRecepitRepository := mongo.NewReceiptSaverRepository(collection)
	receiptCreator := services.NewReciptCreator(mongoRecepitRepository)

	receiptFinderRepository := mongo.NewReceiptFinderRepository(collection)
	receiptSearcher := services.NewReceiptSearcher(receiptFinderRepository)

	routes := routes.NewGroup("/api/v1/receipts/", m.CORS, m.Logger, m.Recovery)
	routes.GET("health", health.HealthCheck())
	routes.POST("process", handlers.HttpHandlerReceiptsCreator(receiptCreator))
	routes.GET("{id}/points", handlers.HttpHandlerReceiptsSearcher(receiptSearcher))

	return routes, nil
}

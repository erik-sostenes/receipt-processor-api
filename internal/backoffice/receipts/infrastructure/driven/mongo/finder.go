package mongo

import (
	"context"
	"errors"
	"fmt"

	"github.com/erik-sostenes/receipt-processor-api/internal/backoffice/receipts/business/domain/receipt"
	"github.com/erik-sostenes/receipt-processor-api/internal/backoffice/receipts/business/ports"
	"github.com/erik-sostenes/receipt-processor-api/pkg/wrongs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ReceiptFinderRepository implements the ports.ReceiptFinder interface
var _ ports.ReceiptFinder = &ReceiptFinderRepository{}

type ReceiptFinderRepository struct {
	receiptsCollection *mongo.Collection
}

func NewReceiptFinderRepository(collection *mongo.Collection) *ReceiptFinderRepository {
	return &ReceiptFinderRepository{
		receiptsCollection: collection,
	}
}

func (r *ReceiptFinderRepository) FindReceipt(ctx context.Context, receiptId *receipt.ReceiptId) (receipt.ReceiptPoints, error) {
	filter := bson.M{"_id": receiptId.Value()}
	opts := options.FindOne().SetProjection(bson.D{{"_id", 0}, {"total_points", 1}})

	var document Receipt

	if err := r.receiptsCollection.FindOne(ctx, filter, opts).Decode(&document); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return receipt.ReceiptPoints{}, wrongs.StatusNotFound(fmt.Sprintf("receipt with id '%s' not found", receiptId.Value()))
		}
	}

	points := receipt.ReceiptPoints{}
	points.Set(document.TotalPoints)

	return points, nil
}

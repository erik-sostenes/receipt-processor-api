package mongo

import (
	"context"
	"fmt"

	"github.com/erik-sostenes/receipt-processor-api/internal/backoffice/receipts/business/domain/receipt"
	"github.com/erik-sostenes/receipt-processor-api/internal/backoffice/receipts/business/ports"
	"github.com/erik-sostenes/receipt-processor-api/pkg/wrongs"
	"go.mongodb.org/mongo-driver/mongo"
)

// MongoReceiptRepository implements the ports.Saver interface
var _ ports.ReceiptSaver = &MongoReceiptRepository{}

type MongoReceiptRepository struct {
	receiptsCollection *mongo.Collection
}

func NewMongoReceiptRepository(collection *mongo.Collection) *MongoReceiptRepository {
	return &MongoReceiptRepository{
		receiptsCollection: collection,
	}
}

func (r *MongoReceiptRepository) SaveReceipt(ctx context.Context, rcpt *receipt.Receipt) (_ receipt.ReceiptId, err error) {
	result, err := r.receiptsCollection.InsertOne(ctx, NewReceipt(rcpt))
	if err != nil {
		if ok := mongo.IsDuplicateKeyError(err); ok {
			return receipt.ReceiptId{}, wrongs.StatusBadRequest(fmt.Sprintf("receipt id '%v' already exists", rcpt.ReceiptId.Value()))
		}
		return
	}

	insertId := result.InsertedID.(string)

	return receipt.NewReceiptId(insertId)
}

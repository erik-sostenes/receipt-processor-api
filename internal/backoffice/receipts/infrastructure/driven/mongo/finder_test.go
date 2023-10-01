package mongo

import (
	"context"
	"errors"
	"testing"

	"github.com/erik-sostenes/receipt-processor-api/internal/backoffice/receipts/business/domain/receipt"
	"github.com/erik-sostenes/receipt-processor-api/internal/backoffice/receipts/infrastructure/drives/handlers/dto"
	"github.com/erik-sostenes/receipt-processor-api/pkg/common"
	connection "github.com/erik-sostenes/receipt-processor-api/pkg/db/mongo"
	"github.com/erik-sostenes/receipt-processor-api/pkg/wrongs"
	"go.mongodb.org/mongo-driver/bson"
)

func TestReceiptFinderRepository(t *testing.T) {
	type FactoryFunc func() (*ReceiptFinderRepository, error)

	tsc := map[string]struct {
		Id string
		FactoryFunc
		expectedPoints uint8
		expectedError  error
	}{
		"Given an existing valid receipt, a score of 109 is expected": {
			Id: "b4308b7b-ccfd-4faa-a16d-164d398cbcfe",
			FactoryFunc: func() (*ReceiptFinderRepository, error) {
				factory := connection.MongoClientFactory{}

				db, err := factory.CreateClient(common.GetEnv("MONGO_DSN"), common.GetEnv("MONGO_DB"))
				if err != nil {
					return nil, err
				}

				receipt, err := receipt.NewReceipt(&dto.ReceiptRequest{
					Id:           "b4308b7b-ccfd-4faa-a16d-164d398cbcfe",
					Retailer:     "M&M Corner Market",
					PurchaseDate: "2022-03-20",
					PurchaseTime: "14:33",
					ItemsRequest: []dto.ItemRequest{
						{ShortDescription: "Gatorade", Price: "2.25"},
						{ShortDescription: "Gatorade", Price: "2.25"},
						{ShortDescription: "Gatorade", Price: "2.25"},
						{ShortDescription: "Gatorade", Price: "2.25"},
					},
					Total: "9.00",
				})
				if err != nil {
					return nil, err
				}

				collection := db.Collection("receipts_test")

				_, err = collection.InsertOne(context.Background(), NewReceipt(receipt))

				return NewReceiptFinderRepository(collection), err
			},
			expectedPoints: 109,
		},
		"Given a non-existent receipt, a score of 0 is expected": {
			Id: "b4308b7b-ccfd-4faa-a16d-164d398cbcfe",
			FactoryFunc: func() (*ReceiptFinderRepository, error) {
				factory := connection.MongoClientFactory{}

				db, err := factory.CreateClient(common.GetEnv("MONGO_DSN"), common.GetEnv("MONGO_DB"))
				if err != nil {
					return nil, err
				}

				collection := db.Collection("receipts_test")

				return NewReceiptFinderRepository(collection), err
			},
			expectedPoints: 0,
			expectedError:  wrongs.StatusNotFound("receipt with id 'b4308b7b-ccfd-4faa-a16d-164d398cbcfe' not found"),
		},
	}

	for name, ts := range tsc {
		t.Run(name, func(t *testing.T) {
			db, err := ts.FactoryFunc()
			if err != nil {
				t.Fatal(err)
			}

			receiptId, err := receipt.NewReceiptId(ts.Id)
			if err != nil {
				t.Fatal(err)
			}

			t.Cleanup(func() {
				if _, err := db.receiptsCollection.DeleteOne(context.Background(), bson.M{"_id": receiptId.Value()}); err != nil {
					t.Fatal(err)
				}
			})

			receiptPoints, err := db.FindReceipt(context.Background(), &receiptId)

			if !errors.Is(ts.expectedError, err) {
				t.Fatalf("status code was expected %v, but it was obtained %v", ts.expectedError, err)
			}

			if ts.expectedPoints != receiptPoints.Value() {
				t.Errorf("points was expected %d, but it was obtained %d", ts.expectedPoints, receiptPoints.Value())
			}
		})
	}
}

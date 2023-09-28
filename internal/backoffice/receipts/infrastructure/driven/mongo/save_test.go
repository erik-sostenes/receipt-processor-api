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
)

func TestMongoReceiptRepository_Save(t *testing.T) {
	type FactoryFunc func() (*MongoReceiptRepository, error)

	tsc := map[string]struct {
		dto.ReceiptRequest
		FactoryFunc
		expectedId    string
		expectedError error
	}{
		"Given a valid non-existing receipt, it will be registered in the mongo collection": {
			ReceiptRequest: dto.ReceiptRequest{
				Id:           "128d0ca8-a70e-4fd4-a43e-6adb25360e57",
				Retailer:     "Walgreens",
				PurchaseDate: "2022-01-02",
				PurchaseTime: "08:13",
				Total:        "2.65",
			},
			FactoryFunc: func() (*MongoReceiptRepository, error) {
				factory := connection.MongoClientFactory{}

				db, err := factory.CreateClient("mongodb://root:password@localhost:27017", "receipts_processor")
				if err != nil {
					return nil, err
				}

				return NewMongoReceiptRepository(db.Collection("receipts_test")), err
			},
			expectedId: "128d0ca8-a70e-4fd4-a43e-6adb25360e57",
		},
		"Given an valid existing receipt, it will not be registered in the mongo collection": {
			ReceiptRequest: dto.ReceiptRequest{
				Id:           "128d0ca8-a70e-4fd4-a43e-6adb25360e57",
				Retailer:     "Walgreens",
				PurchaseDate: "2022-01-02",
				PurchaseTime: "08:13",
				Total:        "2.65",
			},
			FactoryFunc: func() (*MongoReceiptRepository, error) {
				factory := connection.MongoClientFactory{}

				db, err := factory.CreateClient("mongodb://root:password@localhost:27017", "receipts_processor")
				if err != nil {
					return nil, err
				}

				receipt, err := receipt.NewReceipt(&dto.ReceiptRequest{
					Id:           "128d0ca8-a70e-4fd4-a43e-6adb25360e57",
					Retailer:     "Walgreens",
					PurchaseDate: "2022-01-02",
					PurchaseTime: "08:13",
					Total:        "2.65",
				})
				if err != nil {
					return nil, err
				}

				collection := db.Collection("receipts_test")

				_, err = collection.InsertOne(context.Background(), NewReceipt(receipt))

				return NewMongoReceiptRepository(collection), err
			},
			expectedError: wrongs.StatusBadRequest("receipt id '128d0ca8-a70e-4fd4-a43e-6adb25360e57' already exists"),
		},
	}

	for name, ts := range tsc {
		t.Run(name, func(t *testing.T) {
			db, err := ts.FactoryFunc()
			if err != nil {
				t.Fatal(err)
			}

			receipt, err := receipt.NewReceipt(&ts.ReceiptRequest)
			if err != nil {
				t.Fatal(err)
			}

			t.Cleanup(func() {
				if _, err := db.receiptsCollection.DeleteOne(context.Background(), common.Map{"_id": receipt.ReceiptId.Value()}); err != nil {
					t.Fatal(err)
				}
			})

			id, err := db.SaveReceipt(context.Background(), receipt)
			if !errors.Is(err, ts.expectedError) {
				t.Errorf("'%v' error was expected, but '%s' error was obtained", ts.expectedError, err)
				t.SkipNow()
			}

			if id.Value() != ts.expectedId {
				t.Errorf("%s id was expected, but %s error was obtained", ts.expectedId, id.Value())
			}
		})
	}
}

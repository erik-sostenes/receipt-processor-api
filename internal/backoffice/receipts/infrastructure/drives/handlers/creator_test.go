package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/erik-sostenes/receipt-processor-api/internal/backoffice/receipts/business/domain/receipt"
	"github.com/erik-sostenes/receipt-processor-api/internal/backoffice/receipts/business/services"
	"github.com/erik-sostenes/receipt-processor-api/internal/backoffice/receipts/infrastructure/driven/memory"
	"github.com/erik-sostenes/receipt-processor-api/internal/backoffice/receipts/infrastructure/drives/handlers/dto"
)

func Test_HttpHandlerReceiptsCreator(t *testing.T) {
	type HandlerFunc func() (http.HandlerFunc, error)

	tsc := map[string]struct {
		*http.Request
		HandlerFunc
		expectedStatusCode int
	}{
		"Given a valid non-existing receipt a status code 201 is expected": {
			Request: httptest.NewRequest(http.MethodPut, "/api/v1/receipts/process", strings.NewReader(
				`{
					"retailer": "Walgreens",
					"purchaseDate": "2022-01-02",
					"purchaseTime": "08:13",
					"total": "2.65",
					"items": [
						{"shortDescription": "Pepsi - 12-oz", "price": "1.25"},
						{"shortDescription": "Dasani", "price": "1.40"}
					]
				}`,
			)),
			HandlerFunc: func() (http.HandlerFunc, error) {
				return HttpHandlerReceiptsCreator(services.NewReciptCreator(memory.NewReciptInMemory())), nil
			},
			expectedStatusCode: http.StatusCreated,
		},
		"Given an invalid non-existing receipt a status code 422 is expected": {
			Request: httptest.NewRequest(http.MethodPut, "/api/v1/receipts/process", strings.NewReader(
				`{
					"retailer": "Walgreens",
					"purchaseDate": "2022-01-02",
					"purchaseTime": "08:13",
					"total": 2.65,
					"items": [
						{"shortDescription": "Pepsi - 12-oz", "price": "1.25"},
						{"shortDescription": "Dasani", "price": 1.40}
					]
				}`,
			)),
			HandlerFunc: func() (http.HandlerFunc, error) {
				return HttpHandlerReceiptsCreator(services.NewReciptCreator(memory.NewReciptInMemory())), nil
			},
			expectedStatusCode: http.StatusUnprocessableEntity,
		},
		"Given an existing valid receipt, a 400 status code is expected": {
			Request: httptest.NewRequest(http.MethodPut, "/api/v1/receipts/process", strings.NewReader(
				`{
					"id": "73ccbff8-6401-4899-bf51-1d0c4e8740d7",
					"retailer": "Walgreens",
					"purchaseDate": "2022-01-02",
					"purchaseTime": "08:13",
					"total": "2.65",
					"items": [
						{"shortDescription": "Pepsi - 12-oz", "price": "1.25"},
						{"shortDescription": "Dasani", "price": "1.40"}
					]
				}`,
			)),
			HandlerFunc: func() (http.HandlerFunc, error) {
				receiptRequest := dto.ReceiptRequest{
					Id:           "73ccbff8-6401-4899-bf51-1d0c4e8740d7",
					Retailer:     "Walgreens",
					PurchaseDate: "2022-01-02",
					PurchaseTime: "08:13",
					Total:        "2.65",
				}

				v, err := receipt.NewReceipt(&receiptRequest)
				if err != nil {
					return nil, err
				}

				m := memory.NewReciptInMemory()
				m.Add(v.ReceiptId, *v)
				return HttpHandlerReceiptsCreator(services.NewReciptCreator(m)), nil
			},
			expectedStatusCode: http.StatusBadRequest,
		},
		"Given the invalid numeric fields of a non-existent receipt, a status code 400 is expected": {
			Request: httptest.NewRequest(http.MethodPut, "/api/v1/receipts/process", strings.NewReader(
				`{
					"retailer": "Walgreens",
					"purchaseDate": "2022-01-02",
					"purchaseTime": "08:13",
					"total": "2.65ss",
					"items": [
						{"shortDescription": "Pepsi - 12-oz", "price": "1.25a"},
						{"shortDescription": "Dasani", "price": "1.40t"}
					]
				}`,
			)),
			HandlerFunc: func() (http.HandlerFunc, error) {
				return HttpHandlerReceiptsCreator(services.NewReciptCreator(memory.NewReciptInMemory())), nil
			},
			expectedStatusCode: http.StatusBadRequest,
		},
		"Given the invalid time fields of a non-existent receipt, a status code 400 is expected": {
			Request: httptest.NewRequest(http.MethodPut, "/api/v1/receipts/process", strings.NewReader(
				`{
					"retailer": "Walgreens",
					"purchaseDate": "yy-mm-dd",
					"purchaseTime": "08:13Hg",
					"total": "2.65",
					"items": [
						{"shortDescription": "Pepsi - 12-oz", "price": "1.25"},
						{"shortDescription": "Dasani", "price": "1.40"}
					]
				}`,
			)),
			HandlerFunc: func() (http.HandlerFunc, error) {
				return HttpHandlerReceiptsCreator(services.NewReciptCreator(memory.NewReciptInMemory())), nil
			},
			expectedStatusCode: http.StatusBadRequest,
		},
		"Given fields are missing from a non-existent invalid receipt, a status code 400 is expected": {
			Request: httptest.NewRequest(http.MethodPut, "/api/v1/receipts/process", strings.NewReader(
				`{
					"purchaseDate": "2022-01-02",
					"purchaseTime": "08:13",
					"total": "2.65",
					"items": [
						{"price": "1.25"},
						{"shortDescription": "Dasani", "price": "1.40"}
					]
				}`,
			)),
			HandlerFunc: func() (http.HandlerFunc, error) {
				return HttpHandlerReceiptsCreator(services.NewReciptCreator(memory.NewReciptInMemory())), nil
			},
			expectedStatusCode: http.StatusBadRequest,
		},
	}

	for name, ts := range tsc {
		t.Run(name, func(t *testing.T) {
			req := ts.Request
			req.Header.Set("Content-type", "application/json; charset=utf-8")

			resp := httptest.NewRecorder()

			h, err := ts.HandlerFunc()
			if err != nil {
				t.Fatal(err)
			}

			h.ServeHTTP(resp, req)

			if ts.expectedStatusCode != resp.Code {
				t.Log(resp.Body.String())
				t.Errorf("status code was expected %d, but it was obtained %d", ts.expectedStatusCode, resp.Code)
			}
		})
	}
}

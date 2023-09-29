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

func Test_HttpHandlerReceiptsSearcher(t *testing.T) {
	type HandlerFunc func() (http.HandlerFunc, error)

	tsc := map[string]struct {
		*http.Request
		HandlerFunc
		expectedResponse   string
		expectedStatusCode int
	}{
		"Given a valid existent receipt, a status code 200 with a score of 15 is expected": {
			Request: httptest.NewRequest(http.MethodGet, "/api/v1/receipts/b4308b7b-ccfd-4faa-a16d-164d398cbcfe/points", http.NoBody),
			HandlerFunc: func() (http.HandlerFunc, error) {
				receiptRequest := dto.ReceiptRequest{
					Id:           "b4308b7b-ccfd-4faa-a16d-164d398cbcfe",
					Retailer:     "Walgreens",
					PurchaseDate: "2022-01-02",
					PurchaseTime: "08:13",
					Total:        "2.65",
					ItemsRequest: []dto.ItemRequest{
						{ShortDescription: "Pepsi - 12-oz", Price: "1.25"},
						{ShortDescription: "Dasani", Price: "1.40"},
					},
				}

				v, err := receipt.NewReceipt(&receiptRequest)
				if err != nil {
					return nil, err
				}

				m := memory.NewReciptInMemory()
				m.Add(v.ReceiptId, *v)

				searcher := services.NewReceiptSearcher(m)

				return HttpHandlerReceiptsSearcher(searcher), nil
			},

			expectedResponse:   `{"points":15}`,
			expectedStatusCode: http.StatusOK,
		},
		"Given a valid existent receipt, a status code 200 with a score of 109 is expected": {
			Request: httptest.NewRequest(http.MethodGet, "/api/v1/receipts/b4308b7b-ccfd-4faa-a16d-164d398cbcfe/points", http.NoBody),
			HandlerFunc: func() (http.HandlerFunc, error) {
				receiptRequest := dto.ReceiptRequest{
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
				}

				v, err := receipt.NewReceipt(&receiptRequest)
				if err != nil {
					return nil, err
				}

				m := memory.NewReciptInMemory()
				m.Add(v.ReceiptId, *v)

				searcher := services.NewReceiptSearcher(m)

				return HttpHandlerReceiptsSearcher(searcher), nil
			},

			expectedResponse:   `{"points":109}`,
			expectedStatusCode: http.StatusOK,
		},
		"Given a valid existent receipt, a status code 200 with a score of 28 is expected": {
			Request: httptest.NewRequest(http.MethodGet, "/api/v1/receipts/b4308b7b-ccfd-4faa-a16d-164d398cbcfe/points", http.NoBody),
			HandlerFunc: func() (http.HandlerFunc, error) {
				receiptRequest := dto.ReceiptRequest{
					Id:           "b4308b7b-ccfd-4faa-a16d-164d398cbcfe",
					Retailer:     "Target",
					PurchaseDate: "2022-01-01",
					PurchaseTime: "13:01",
					ItemsRequest: []dto.ItemRequest{
						{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
						{ShortDescription: "Emils Cheese Pizza", Price: "12.25"},
						{ShortDescription: "Knorr Creamy Chicken", Price: "1.26"},
						{ShortDescription: "Doritos Nacho Cheese", Price: "3.35"},
						{ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ", Price: "12.00"},
					},
					Total: "35.35",
				}

				v, err := receipt.NewReceipt(&receiptRequest)
				if err != nil {
					return nil, err
				}

				m := memory.NewReciptInMemory()
				m.Add(v.ReceiptId, *v)

				searcher := services.NewReceiptSearcher(m)

				return HttpHandlerReceiptsSearcher(searcher), nil
			},

			expectedResponse:   `{"points":28}`,
			expectedStatusCode: http.StatusOK,
		},
		"Given a non-existent receipt, a 404 status code is expected": {
			Request: httptest.NewRequest(http.MethodGet, "/api/v1/receipts/b4308b7b-ccfd-4faa-a16d-164d398cbcfe/points", http.NoBody),
			HandlerFunc: func() (http.HandlerFunc, error) {
				searcher := services.NewReceiptSearcher(memory.NewReciptInMemory())

				return HttpHandlerReceiptsSearcher(searcher), nil
			},

			expectedResponse:   `{"message":"receipt with id 'b4308b7b-ccfd-4faa-a16d-164d398cbcfe' not found"}`,
			expectedStatusCode: http.StatusNotFound,
		},
		"Given an erroneous identifier as a query parameter, a status code 400 is expected": {
			Request: httptest.NewRequest(http.MethodGet, "/api/v1/receipts/b4-cfd-4faa-a16d-1d398/points", http.NoBody),
			HandlerFunc: func() (http.HandlerFunc, error) {
				return HttpHandlerReceiptsSearcher(services.NewReceiptSearcher(memory.NewReciptInMemory())), nil
			},

			expectedResponse:   `{"message":"incorrect b4-cfd-4faa-a16d-1d398 uuid unique identifier, must be a uuid value"}`,
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
				t.Fatalf("status code was expected %d, but it was obtained %d", ts.expectedStatusCode, resp.Code)
			}

			if strings.TrimSpace(ts.expectedResponse) != strings.TrimSpace(resp.Body.String()) {
				t.Log(resp.Body.String())
				t.Errorf("points was expected %s, but it was obtained %s", ts.expectedResponse, resp.Body.String())
			}
		})
	}
}

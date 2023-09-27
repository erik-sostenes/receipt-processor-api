package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/erik-sostenes/receipt-processor-api/internal/backoffice/receipts/business/domain"
	"github.com/erik-sostenes/receipt-processor-api/internal/backoffice/receipts/business/services"
	"github.com/erik-sostenes/receipt-processor-api/internal/backoffice/receipts/infrastructure/driven/memory"
)

func Test_HttpHandlerReceiptsCreator(t *testing.T) {
	type HandlerFunc func() (http.HandlerFunc, error)

	tsc := map[string]struct {
		*http.Request
		HandlerFunc
		expectedStatusCode int
	}{
		"Given a valid non-existing receipt a status code 201 is expected": {
			Request: httptest.NewRequest(http.MethodPut, "/api/v1/receipts/create", strings.NewReader(
				`{
					"id": "cb3774b9-5637-433a-b669-a59e8ff8eb15",
					"retailer": "Walgreens",
					"purchaseDate": "2022-01-02",
					"purchaseTime": "08:13",
					"total": 2.65,
					"items": [
						{"shortDescription": "Pepsi - 12-oz", "price": 1.25},
						{"shortDescription": "Dasani", "price": 1.40}
					]
				}`,
			)),
			HandlerFunc: func() (http.HandlerFunc, error) {
				return HttpHandlerReceiptsCreator(*services.NewReciptCreator(memory.NewReciptInMemory())), nil
			},
			expectedStatusCode: http.StatusCreated,
		},
		"Given an invalid non-existing receipt a status code 422 is expected": {
			Request: httptest.NewRequest(http.MethodPut, "/api/v1/receipts/create", strings.NewReader(
				`{
					"id": "757ce493-020f-4e9f-8851-0deebbf39637",
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
				return HttpHandlerReceiptsCreator(*services.NewReciptCreator(memory.NewReciptInMemory())), nil
			},
			expectedStatusCode: http.StatusUnprocessableEntity,
		},
		"Given an existing valid receipt, a 400 status code is expected": {
			Request: httptest.NewRequest(http.MethodPut, "/api/v1/receipts/create", strings.NewReader(
				`{
					"id": "06058b7c-b39d-4187-9837-5c2fe21c535e",					
					"retailer": "Walgreens",
					"purchaseDate": "2022-01-02",
					"purchaseTime": "08:13",
					"total": 2.65,
					"items": [
						{"shortDescription": "Pepsi - 12-oz", "price": 1.25},
						{"shortDescription": "Dasani", "price": 1.40}
					]
				}`,
			)),
			HandlerFunc: func() (http.HandlerFunc, error) {
				receipt := domain.Receipt{
					Id: "06058b7c-b39d-4187-9837-5c2fe21c535e",
				}

				m := memory.NewReciptInMemory()
				m.Add(receipt.Id, receipt)

				return HttpHandlerReceiptsCreator(*services.NewReciptCreator(m)), nil
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

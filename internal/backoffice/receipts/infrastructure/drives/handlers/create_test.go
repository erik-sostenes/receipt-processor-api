package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_HttpHandlerReceiptsCreator(t *testing.T) {
	type HandlerFunc func() (http.HandlerFunc, error)

	tsc := map[string]struct {
		*http.Request
		HandlerFunc
		expectedStatusCode int
	}{
		"Given a valid non-existing receipts a status code 201 is expected": {
			Request: httptest.NewRequest(http.MethodPut, "/api/v1/receipts/create", strings.NewReader(
				`{
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
				return HttpHandlerReceiptsCreator(), nil
			},
			expectedStatusCode: http.StatusCreated,
		},
		"Given a invalid non-existing receipts a status code 422 is expected": {
			Request: httptest.NewRequest(http.MethodPut, "/api/v1/receipts/create", strings.NewReader(
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
				return HttpHandlerReceiptsCreator(), nil
			},
			expectedStatusCode: http.StatusUnprocessableEntity,
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

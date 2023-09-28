package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/erik-sostenes/receipt-processor-api/internal/backoffice/receipts/business/services"
)

func Test_HttpHandlerReceiptsSearcher(t *testing.T) {
	type HandlerFunc func() (http.HandlerFunc, error)

	tsc := map[string]struct {
		*http.Request
		HandlerFunc
		expectedResponse   string
		expectedStatusCode int
	}{
		"Given a valid non-existing receipt a status code 201 is expected": {
			Request: httptest.NewRequest(http.MethodGet, "/api/v1/receipts/b4308b7b-ccfd-4faa-a16d-164d398cbcfe/points", http.NoBody),
			HandlerFunc: func() (http.HandlerFunc, error) {
				searcher := services.NewReceiptSearcher()

				return HttpHandlerReceiptsSearcher(searcher), nil
			},

			expectedResponse:   `{"points":40}`,
			expectedStatusCode: http.StatusOK,
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

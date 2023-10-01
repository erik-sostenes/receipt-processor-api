package handlers

import (
	"net/http"
	"regexp"

	"github.com/erik-sostenes/receipt-processor-api/internal/backoffice/receipts/business/domain/receipt"
	"github.com/erik-sostenes/receipt-processor-api/internal/backoffice/receipts/business/ports"
	"github.com/erik-sostenes/receipt-processor-api/pkg/common"
	"github.com/erik-sostenes/receipt-processor-api/pkg/server/response"
)

func HttpHandlerReceiptsSearcher(searcher ports.ReceiptSearcher) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		patron := `\/api\/v1\/receipts\/([^/]+)\/points`

		expresionRegular := regexp.MustCompile(patron)
		id := expresionRegular.FindStringSubmatch(r.URL.Path)[1]

		receiptId, err := receipt.NewReceiptId(id)
		if err != nil {
			_ = response.ErrorHandler(w, err)
			return
		}

		receiptPoints, err := searcher.SearchReceipt(r.Context(), &receiptId)
		if err != nil {
			_ = response.ErrorHandler(w, err)
			return
		}

		_ = response.JSON(w, http.StatusOK, common.Map{"points": receiptPoints.Value()})
	}
}

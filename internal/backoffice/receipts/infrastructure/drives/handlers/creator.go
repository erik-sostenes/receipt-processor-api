package handlers

import (
	"net/http"

	"github.com/erik-sostenes/receipt-processor-api/internal/backoffice/receipts/business/domain/receipt"
	"github.com/erik-sostenes/receipt-processor-api/internal/backoffice/receipts/business/ports"
	"github.com/erik-sostenes/receipt-processor-api/internal/backoffice/receipts/infrastructure/drives/handlers/dto"
	"github.com/erik-sostenes/receipt-processor-api/pkg/server/response"
)

func HttpHandlerReceiptsCreator(creator ports.ReceiptCreator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request dto.ReceiptRequest

		if ok, err := response.Bind(w, r, &request); err != nil && !ok {
			return
		}

		receipt, err := receipt.NewReceipt(&request)
		if err != nil {
			_ = response.ErrorHandler(w, err)
			return
		}

		receiptId, err := creator.CreateReceipt(r.Context(), receipt)
		if err != nil {
			_ = response.ErrorHandler(w, err)
			return
		}

		_ = response.JSON(w, http.StatusCreated, map[string]any{"id": receiptId.Value()})
	}
}

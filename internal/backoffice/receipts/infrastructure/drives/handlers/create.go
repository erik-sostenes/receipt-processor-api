package handlers

import (
	"net/http"

	"github.com/erik-sostenes/receipt-processor-api/internal/backoffice/receipts/business/services"
	"github.com/erik-sostenes/receipt-processor-api/internal/backoffice/receipts/infrastructure/drives/handlers/dto"
	"github.com/erik-sostenes/receipt-processor-api/pkg/server/response"
)

func HttpHandlerReceiptsCreator(creator services.ReceiptCreator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request dto.ReceiptRequest

		if ok, err := response.Bind(w, r, &request); err != nil || !ok {
			return
		}

		err := creator.Create(r.Context(), &request)
		if err != nil {
			_ = response.ErrorHandler(w, err)
			return
		}

		_ = response.JSON(w, http.StatusCreated, nil)
	}
}

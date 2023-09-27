package handlers

import (
	"net/http"

	"github.com/erik-sostenes/receipt-processor-api/internal/backoffice/receipts/infrastructure/drives/handlers/dto"
	"github.com/erik-sostenes/receipt-processor-api/pkg/server/response"
)

func HttpHandlerReceiptsCreator() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request dto.ReceiptsRequest

		if ok, err := response.Bind(w, r, &request); err != nil || !ok {
			return
		}

		_ = response.JSON(w, http.StatusCreated, nil)
	}
}

package accounts

import (
	"encoding/json"
	"net/http"
	"yatter-backend-go/app/handler/auth"
	"yatter-backend-go/app/handler/httperror"
)

// Request body for `POST /v1/accounts`
type UpdateRequest struct {
	DisplayName string `json:"display_name"`
}

// Handle request for `POST /v1/accounts`
func (h *handler) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req UpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httperror.BadRequest(w, err)
		return
	}

	// 既存のaccountデータを持ってくる
	account := auth.AccountOf(r)

	// 持ってきたデータに対して更新
	if err := account.SetDisplayName(req.DisplayName); err != nil {
		httperror.BadRequest(w, err)
		return
	}

	// 更新したデータを保存
	if err := h.app.Dao.Account().Update(ctx, *account); err != nil {
		httperror.InternalServerError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(account); err != nil {
		httperror.InternalServerError(w, err)
		return
	}
}

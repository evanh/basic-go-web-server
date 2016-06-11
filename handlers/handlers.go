package handlers

import (
	"encoding/json"
	"net/http"

	"golang.org/x/net/context"

	"github.com/evanh/fundmyworld/services"
)

func Root(ctx context.Context, w http.ResponseWriter, r *http.Request, svc *services.FundService) {
	rows, err := svc.GetTestRows()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	output, err := json.Marshal(rows)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(output)
}

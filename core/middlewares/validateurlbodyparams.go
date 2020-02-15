package middlewares

import (
	"encoding/json"
	"github.com/giovannibussola/trabalho-pos-golang/core/beer"
	"github.com/giovannibussola/trabalho-pos-golang/core/errors"
	"github.com/gorilla/context"
	"net/http"
)

func ValidateBodyParams(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var b beer.Beer

		err := json.NewDecoder(r.Body).Decode(&b)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(errors.FormatJSONError(err.Error(), "Parametros do corpo da requisição inválidos!"))
			return
		}

		context.Set(r, "beer", b)
		next.ServeHTTP(w, r)
	})
}

package middlewares

import (
	"github.com/eminetto/pos-web-go/core/beer"
	"github.com/eminetto/pos-web-go/core/errors"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func ValidateUrlParam(service beer.UseCase, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.ParseInt(vars["id"], 10, 64)

		if err != nil {
			w.Write(errors.FormatJSONError(err.Error(), "Parametro id incorreto na requisição"))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		_, err = service.Get(id)
		if err != nil {
			w.Write(errors.FormatJSONError(err.Error(), "Cerveja não encontrada!"))
			w.WriteHeader(http.StatusNotFound)
			return
		}

		context.Set(r, "id", id)
		next.ServeHTTP(w, r)
	})
}

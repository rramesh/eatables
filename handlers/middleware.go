package handlers

import (
	"context"
	"net/http"

	"github.com/rramesh/eatables/data"
)

//MiddlewareValidateItem validates JSON from request body before passing back to router
func (items Items) MiddlewareValidateItem(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		it := &data.Item{}
		err := data.FromJSON(it, r.Body)
		if err != nil {
			items.l.Println("[ERROR] Deserializing JSON")
			rw.WriteHeader(http.StatusBadRequest)
			data.ToJSON(&GenericError{Message: err.Error()}, rw)
			return
		}

		errs := items.v.Validate(it)
		if len(errs) != 0 {
			items.l.Println("[ERROR] Item request validation failed")
			rw.WriteHeader(http.StatusUnprocessableEntity)
			data.ToJSON(&ValidationError{Message: errs.Errors()}, rw)
			return
		}

		ctx := context.WithValue(r.Context(), KeyItem{}, *it)

		r = r.WithContext(ctx)
		next.ServeHTTP(rw, r)
	})
}

package handlers

import (
	"context"
	"net/http"

	"github.com/rramesh/eatables/data"
)

// MiddlewareValidateItem validates JSON from request body before passing back to router
func (ih ItemHandler) MiddlewareValidateItem(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		it := &data.Item{}
		err := data.FromJSON(it, r.Body)
		if err != nil {
			ih.l.Debug("Deserializing JSON", "error", err)
			rw.WriteHeader(http.StatusBadRequest)
			data.ToJSON(&GenericMessage{Message: err.Error()}, rw)
			return
		}
		if r.Method == http.MethodPut {
			if it.SKU == "" {
				ih.l.Debug("SKU not provided for update")
				data.ToJSON(&GenericMessage{Message: "Item SKU not provided"}, rw)
				return
			}
		} else {
			errs := ih.v.Validate(it)
			if len(errs) != 0 {
				ih.l.Error("Item request validation failed", "error", errs)
				rw.WriteHeader(http.StatusUnprocessableEntity)
				data.ToJSON(&ValidationError{Message: errs.Errors()}, rw)
				return
			}
		}
		ctx := context.WithValue(r.Context(), KeyItem{}, *it)

		r = r.WithContext(ctx)
		next.ServeHTTP(rw, r)
	})
}

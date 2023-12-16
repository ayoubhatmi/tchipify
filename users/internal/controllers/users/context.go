package users

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"tchipify/users/internal/models"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

func Ctx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Parse ID from URL parameter
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			logrus.Errorf("parsing error: %s", err.Error())
			customError := &models.CustomError{
				Message: fmt.Sprintf("cannot parse id (%s) as integer", idStr),
				Code:    http.StatusUnprocessableEntity,
			}
			w.WriteHeader(customError.Code)
			body, _ := json.Marshal(customError)
			_, _ = w.Write(body)
			return
		}

		ctx := context.WithValue(r.Context(), "collectionId", id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

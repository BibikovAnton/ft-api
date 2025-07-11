package req

import (
	"net/http"

	"github.com/BibikovAnton/finance-tracker-api/internal/res"
)

func HandlerBody[T any](w *http.ResponseWriter, r *http.Request) (*T, error) {
	body, err := Decode[T](r.Body)
	if err != nil {
		res.Json(*w, err.Error(), 402)
		return nil, err
	}

	err = Validate[T](body)
	if err != nil {
		res.Json(*w, err.Error(), 402)
		return nil, err
	}
	return &body, nil
}

package req

import (
	"net/http"
	"rest_go_kv/pkg/res"
)

func HandleBody[T any](w *http.ResponseWriter, r *http.Request) (*T, error) {
	// читаем и декодируем боди
	body, err := Decode[T](r.Body)
	if err != nil {
		res.Json(*w, err.Error(), 402)
		return nil, err
	}
	err = IsValid(body)
	if err != nil {
		res.Json(*w, err.Error(), 402)
		return nil, err
	}
	// если смогли декодировать и валидировать, отправляем обратно указатель на боди и нил
	return &body, nil
}

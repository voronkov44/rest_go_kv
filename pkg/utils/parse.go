package utils

import (
	"net/http"
	"strconv"
)

// Извлекаем и преобразуем id и user_id из path value в uint

func ParseID(r *http.Request) (uint, error) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(id), nil
}

func ParseUserID(r *http.Request) (uint, error) {
	idStr := r.PathValue("user_id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(id), nil
}

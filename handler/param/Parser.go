package param

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

//Int 	:
func Int(r *http.Request, p string) int {
	idString := chi.URLParam(r, p)

	if idString != "" {
		id, err := strconv.Atoi(idString)
		if err != nil {
			return 0
		}
		return id
	}
	return 0
}

//UInt 	:
func UInt(r *http.Request, p string) uint {
	return uint(Int(r, p))
}

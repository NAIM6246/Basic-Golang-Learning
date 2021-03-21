package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func Admin(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		role, _ := strconv.Atoi(r.Header.Get(KeyUserID))
		fmt.Println("herr")
		fmt.Println(role)
		if role == 1 {
			handler.ServeHTTP(w, r)
		} else {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode("Access Denied")
		}
	})
}

package helpers

import (
	"net/http"
	"github.com/gorilla/mux"
)

func GetParam(p string, r *http.Request) string {
	vars := mux.Vars(r)
	return vars[p]
}
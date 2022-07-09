package mapper

import "net/http"

func ErrToISE(w http.ResponseWriter, err error) {

	http.Error(w, err.Error(), http.StatusInternalServerError)
}

func StrToISE(w http.ResponseWriter, str string) {

	http.Error(w, str, http.StatusInternalServerError)
}

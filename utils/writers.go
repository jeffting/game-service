package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// SendJSON writes the passed interface to the stream and sets the Content-Type to application/json
func SendJSON(w http.ResponseWriter, i interface{}) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(i)
	if err != nil {
		fmt.Println(err)
	}
}

// SendError respond to the request with a given error with it's associated code.
func SendError(w http.ResponseWriter, httpStatus int, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	json.NewEncoder(w).Encode(&err)
}

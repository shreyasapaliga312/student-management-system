package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Parses the JSON body of an HTTP request into a Go struct
func ParseBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}


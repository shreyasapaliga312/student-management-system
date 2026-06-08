package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func WriteJSONResponse(w http.ResponseWriter, status int, data interface{}) {
	res, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(res)
}

func WriteErrorResponse(w http.ResponseWriter, status int, message string) {
	e := make(map[string]string)

	e["Message"] = message
	e["Status"] = strconv.Itoa(status)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(e)

}

func GetId(r *http.Request, typeId string) int64 {
	vars := mux.Vars(r)

	id := vars[typeId]

	resultId, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		fmt.Println("Error while Parsing")
		return 0
	}
	return resultId
}

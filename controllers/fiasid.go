package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type response struct {
	FiasID string `json:"fias_id"`
}

func GetFiasID(w http.ResponseWriter, req *http.Request) {
	resp := response{FiasID: "0000-000000000000000-0000"}

	// errors here shoudn't be possible
	respJSON, _ := json.Marshal(resp)

	fmt.Fprintf(w, string(respJSON))
}

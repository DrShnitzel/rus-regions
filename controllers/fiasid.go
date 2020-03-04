package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/DrShnitzel/rus-regions/regions"
)

type response struct {
	FiasID string `json:"fias_id"`
	// for debugging
	Name   string   `json:"name"`
	Errors []string `json:"errors"`
}

// GetFiasID handles response for fias_id request.
// It expects lat and long request parapmeters.
// Request method is not checked, because there is no reason why it should
func GetFiasID(w http.ResponseWriter, r *http.Request) {
	var resp response

	w.Header().Set("Content-Type", "application/json")

	lat, latErr := strconv.ParseFloat(r.URL.Query().Get("lat"), 64)
	long, longErr := strconv.ParseFloat(r.URL.Query().Get("long"), 64)
	if latErr != nil || longErr != nil {
		w.WriteHeader(400)
		resp.Errors = append(resp.Errors, "Incorrect or missing long or lat params")
	}

	if len(resp.Errors) < 1 {
		var err error
		resp.FiasID, resp.Name, err = regions.FiasIDByLatLong(lat, long)
		if err != nil {
			w.WriteHeader(404)
			resp.Errors = append(resp.Errors, err.Error())
		}
	}

	// errors here shoudn't be possible
	respJSON, _ := json.Marshal(resp)

	log.Println(r.URL.String() + " | " + string(respJSON))
	fmt.Fprintf(w, string(respJSON))
}

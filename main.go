package main

import (
	"log"
	"net/http"

	"github.com/DrShnitzel/rus-regions/controllers"
	"github.com/DrShnitzel/rus-regions/regions"
)

func main() {
	log.Println("application started")

	regions.PrepareRegions()

	http.HandleFunc("/api/v1/fias_id", controllers.GetFiasID)
	http.ListenAndServe(":1337", nil)
}

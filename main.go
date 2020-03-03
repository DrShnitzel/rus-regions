package main

import (
	"log"
	"net/http"
	"os"

	"github.com/DrShnitzel/rus-regions/controllers"
	"github.com/DrShnitzel/rus-regions/regions"
)

func main() {
	log.Println("application started")

	regions.PrepareRegions()

	http.HandleFunc("/api/v1/fias_id", controllers.GetFiasID)

	port := os.Getenv("REGIONS_PORT")
	if port == "" {
		port = "1337"
	}
	log.Println("listning on port: " + port)
	http.ListenAndServe(":"+port, nil)
}

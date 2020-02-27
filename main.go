package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/DrShnitzel/rus-regions/controllers"
	"github.com/DrShnitzel/rus-regions/draw"
	"github.com/golang/geo/s2"
)

var parsedData [][]float64

func main() {
	fmt.Println("startu")
	parseData("data/regions4.json")
	polygon := prepareRegionsData()

	x := 57.605012
	y := 38.746949

	http.HandleFunc("/api/v1/fias_id", controllers.GetFiasID)
	http.ListenAndServe(":1337", nil)

	result := polygon.ContainsPoint(s2.PointFromLatLng(s2.LatLngFromDegrees(x, y)))
	fmt.Println(result)

	if os.Getenv("DEBUG") == "true" {
		draw.RegionWithPoint(&parsedData, x, y)
	}

}

func parseData(filePath string) {
	rawData, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(rawData, &parsedData)
	if err != nil {
		panic(err)
	}
}

func prepareRegionsData() *s2.Loop {

	var points []s2.Point
	for _, v := range parsedData {
		points = append(points, s2.PointFromLatLng(s2.LatLngFromDegrees(v[1], v[0])))
	}

	return s2.LoopFromPoints(points)
}

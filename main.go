package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/DrShnitzel/rus-regions/draw"
	"github.com/golang/geo/s2"
)

var parsedData [][]float64

func main() {
	fmt.Println("startu")
	parseData("data/regions3.json")
	polygon := prepareRegionsData()

	result := polygon.ContainsPoint(s2.PointFromLatLng(s2.LatLngFromDegrees(58.042856, 38.921909)))
	fmt.Println(result)

	draw.DrawRegion(&parsedData)
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

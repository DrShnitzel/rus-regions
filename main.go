package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/golang/geo/s2"
)

func main() {
	fmt.Println("startu")
	polygon := prepareRegionsData()
	// fmt.Println(polygon)
	result := polygon.ContainsPoint(s2.PointFromLatLng(s2.LatLngFromDegrees(64.219472, 55.216981)))
	fmt.Println(result)
}

func prepareRegionsData() *s2.Loop {
	rawData, err := ioutil.ReadFile("data/regions.json")
	if err != nil {
		panic(err)
	}

	var parsedData [][]float64
	err = json.Unmarshal(rawData, &parsedData)
	if err != nil {
		panic(err)
	}

	var points []s2.Point
	for _, v := range parsedData {
		points = append(points, s2.PointFromLatLng(s2.LatLngFromDegrees(v[0], v[1])))
	}

	return s2.LoopFromPoints(points)
	// fmt.Println(loop)

	// return s2.PolygonFromLoops([]*s2.Loop{loop})
}

package regions

import (
	"encoding/json"
	"io/ioutil"

	"github.com/golang/geo/s2"
)

var parsedData [][]float64
var polygon *s2.Loop

func parseData(filePath string) {
	rawData, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(rawData, &parsedData)
	if err != nil {
		panic(err)
	}

	log.Println("data file successefuly parsed")
}

func prepareData() {
	filePath := "data/regions4.json"
	parseData(filePath string)

	var points []s2.Point
	for _, v := range parsedData {
		points = append(points, s2.PointFromLatLng(s2.LatLngFromDegrees(v[1], v[0])))
	}

	polygon = s2.LoopFromPoints(points)

	log.Println("regions data loaded into memory")
}

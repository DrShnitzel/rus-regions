package regions

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

var regionsList []*region

func parseData(filePath string) {
	rawData, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(rawData, &regionsList)
	if err != nil {
		panic(err)
	}

	log.Println("data file successefuly parsed")
}

func PrepareRegions() {
	filePath := "data/regions.json"
	parseData(filePath)

	for _, v := range regionsList {
		v.addPolygons()
	}

	log.Println("regions data loaded into memory")
}

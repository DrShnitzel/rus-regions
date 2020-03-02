package regions

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

var regionsList []region

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
		fmt.Println(v.Name)
		fmt.Println(v.FiasID)
		fmt.Println(len(v.Coordinates))
		fmt.Println("polygons")
		fmt.Println(len(v.polygons))
		v.addPolygons()
		fmt.Println(len(v.polygons))
	}

	log.Println("regions data loaded into memory")
}

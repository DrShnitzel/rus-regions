package regions

import "github.com/golang/geo/s2"

type region struct {
	Name    string `json:"name"`
	FiasID  string `json:"fias_id"`
	polygon *s2.Polygon
	loops   [][][]float64 `json:"loops"`
}

func (region) addPolygon() {
	var points []s2.Point
	for _, v := range parsedData {
		points = append(points, s2.PointFromLatLng(s2.LatLngFromDegrees(v[1], v[0])))
	}

	polygon = s2.LoopFromPoints(points)
}

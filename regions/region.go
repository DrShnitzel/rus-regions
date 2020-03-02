package regions

import "github.com/golang/geo/s2"

type region struct {
	Name        string `json:"name"`
	FiasID      string `json:"fias_id"`
	polygons    []*s2.Polygon
	Coordinates [][][][]float64 `json:"loops"`
}

func (r region) addPolygons() {
	for _, loops := range r.Coordinates {
		r.polygons = append(r.polygons, addPolygon(loops))
	}
}

func addPolygon(loops [][][]float64) *s2.Polygon {
	var s2LoopList []*s2.Loop
	for _, loop := range loops {
		var points []s2.Point
		for _, v := range loop {
			points = append(points, s2.PointFromLatLng(s2.LatLngFromDegrees(v[1], v[0])))
		}

		s2LoopList = append(s2LoopList, s2.LoopFromPoints(points))
	}
	return s2.PolygonFromLoops(s2LoopList)
}

func (r region) containsPoint(lat, long float64) bool {
	for _, p := range r.polygons {
		return p.ContainsPoint(s2.PointFromLatLng(s2.LatLngFromDegrees(lat, long)))
	}
	return false
}
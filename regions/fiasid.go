package regions

type notFoundError struct{}

func (e notFoundError) Error() string {
	return "Cannot find region by given coordinates"
}

// FiasIDByLatLong scans regionsList for region that contains
// point with lat long coordinates and returns fias_id of this region.
// If region cannot be found it returns an error.
func FiasIDByLatLong(lat, long float64) (string, string, error) {
	for _, region := range regionsList {
		if region.containsPoint(lat, long) {
			return region.FiasID, region.Name, nil
		}
	}
	return "", "", notFoundError{}
}

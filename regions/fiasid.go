package regions

func FiasIDByLatLong(lat, long float64) string {
	for _, region := range regionsList {
		if region.containsPoint(lat, long) {
			return region.FiasID
		}
	}
	// TODO: tmp do not leave it like this!
	return "too bad"
}

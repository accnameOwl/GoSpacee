package planet

type Earth struct {
	Date       string
	CloudScore bool
	apiKey     string
	Latitude   float32
	Longtitude float32
	// Dimention of image to recieve
	//!Width and height in degrees, default 0.025
	Dimention float32
}

package model

type IempotencyResponse struct {
	// The status code of the API response
	StatusCode string
	// The content length of the API response
	ContentLength string
	// The content type of the API response
	ContentType string
	// The long format date when the API request was received
	Date int64
	// An API resource
	Resource interface{}
	// The URL of the API request
	RequestURL string
}

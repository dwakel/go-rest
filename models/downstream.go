package models

type SampleDownstreamResponse struct {
	StatusCode    string `json:"statusCode"`
	StatusMessage string `json:"statusMessage"`
	Data		  string `json:"data"`
}


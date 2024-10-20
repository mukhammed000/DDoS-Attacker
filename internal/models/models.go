package models

type LoadTestRequest struct {
	URL         string `json:"url" binding:"required"`
	NumRequests int    `json:"num_requests" binding:"required"`
	Delay       int    `json:"delay"`
}

type LoadTestResponse struct {
	
}

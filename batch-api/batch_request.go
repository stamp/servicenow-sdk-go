package batchapi

type RestRequest interface {
}

type BatchRequest struct {
	RestRequests []RestRequest `json:"rest_requests"`
}

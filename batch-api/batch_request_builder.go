package batchapi

import "github.com/michaeldcanady/servicenow-sdk-go/core"

type BatchRequestBuilder struct {
	core.RequestBuilder
	request BatchRequest
}

type BatchableRequest interface {
	ToBatch() RestRequest
}

func NewBatchRequestBuilder(client core.Client, pathParameters map[string]string) *BatchRequestBuilder {
	requestBuilder := core.NewRequestBuilder(
		client,
		"{+baseurl}/batch",
		pathParameters,
	)
	return &BatchRequestBuilder{
		RequestBuilder: *requestBuilder,
		request:        BatchRequest{},
	}
}

func (rB *BatchRequestBuilder) AddRequest(request BatchableRequest) {
	rB.request.RestRequests = append(rB.request.RestRequests, request.ToBatch())
}

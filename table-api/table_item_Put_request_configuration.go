package tableapi

import "github.com/michaeldcanady/servicenow-sdk-go/core"

type TableItemPutRequestConfiguration struct {
	Header          interface{}
	QueryParameters *TableItemRequestBuilderPutQueryParameters
	Data            interface{}
	ErrorMapping    core.ErrorMapping
	response        *TableItemResponse
}

func (rC *TableItemPutRequestConfiguration) toConfiguration() *core.RequestConfiguration {
	return &core.RequestConfiguration{
		Header:          rC.Header,
		QueryParameters: rC.QueryParameters,
		Data:            rC.Data,
		ErrorMapping:    rC.ErrorMapping,
		Response:        rC.response,
	}
}

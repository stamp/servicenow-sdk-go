package servicenowsdkgo

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	tableapi "github.com/michaeldcanady/servicenow-sdk-go/table-api"
)

type NowRequestBuilder struct {
	core.RequestBuilder
}

// NewNowRequestBuilder creates a new instance of the NowRequestBuilder associated with the given URL and Client.
// It accepts the URL and Client as parameters and returns a pointer to the created NowRequestBuilder.
//	package servicenowsdkgo

// import (
//
//	servicenowsdkgo "github.com/michaeldcanady/servicenow-sdk-go"
//
// )
//
//	func main() {
//		url := "instance.service-now.com/api"
//		client := servicenowsdkgo.NewClient()
//		nowRequestBuilder := servicenowsdkgo.NewNowrequestBuilder(url, client)
//	}
func NewNowRequestBuilder(url string, client *ServiceNowClient) *NowRequestBuilder {
	pathParameters := map[string]string{"baseurl": url}
	requestBuilder := core.NewRequestBuilder(client, "{+baseurl}/Now", pathParameters)
	return &NowRequestBuilder{
		*requestBuilder,
	}
}

// Table returns a TableRequestBuilder associated with the NowRequestBuilder.
// It accepts a table name as a parameter and constructs the URL for table-related requests.
// The returned TableRequestBuilder can be used to build and execute table-related requests.
func (N *NowRequestBuilder) Table(tableName string) *tableapi.TableRequestBuilder {
	N.RequestBuilder.PathParameters["table"] = tableName
	return tableapi.NewTableRequestBuilder(N.RequestBuilder.Client.(*ServiceNowClient), N.RequestBuilder.PathParameters)
}

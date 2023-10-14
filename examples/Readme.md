# Examples

This folder exists to provide generic, cookie cutter examples on how to implement [servicenow-sdk-go](https://github.com/michaeldcanady/servicenow-sdk-go).

## Table Requests

### Basic Authentication

#### Get Request

##### With Parameters

```golang
import (
    servicenowsdkgo "github.com/michaeldcanady/servicenow-sdk-go"
    "github.com/michaeldcanady/servicenow-sdk-go/credentials"
    tableapi "github.com/michaeldcanady/servicenow-sdk-go/table-api"
)

func main() {
    //Defined at top for simplicity
    username := ""
    password := ""
    instance := ""
    tableName := "incident"

    // Create a new UsernamePasswordCredential using account credentials
    cred := credentials.NewUsernamePasswordCredential(username, password)

    // Create a new client using previously created credential and instance
    client := servicenowsdkgo.NewClient(cred, instance)

    // Optional parameters
    params := &tableapi.TableRequestBuilderGetQueryParameters{
        Limit: int32(1),
    }

    //extend client with table name and parameters you'd like to use
    //
    // recieve and error and TableCollectionResponse
    response, err := client.Now().Table(tableName).Get(params)
    if err != nil {
        panic(err)
    }
}
```

##### Without Parameters

```golang
import (
    servicenowsdkgo "github.com/michaeldcanady/servicenow-sdk-go"
    "github.com/michaeldcanady/servicenow-sdk-go/credentials"
    tableapi "github.com/michaeldcanady/servicenow-sdk-go/table-api"
)

func main() {
    //Defined at top for simplicity
    username := ""
    password := ""
    instance := ""
    tableName := "incident"

    // Create a new UsernamePasswordCredential using account credentials
    cred := credentials.NewUsernamePasswordCredential(username, password)

    // Create a new client using previously created credential and instance
    client := servicenowsdkgo.NewClient(cred, instance)

    //extend client with table name and parameters you'd like to use
    //
    // recieve and error and TableCollectionResponse
    response, err := client.Now().Table(tableName).Get(nil)
    if err != nil {
        panic(err)
    }
}

```

#### Post Request

##### With Parameters

```golang
import (
    servicenowsdkgo "github.com/michaeldcanady/servicenow-sdk-go"
    "github.com/michaeldcanady/servicenow-sdk-go/credentials"
    tableapi "github.com/michaeldcanady/servicenow-sdk-go/table-api"
)

func main() {
    //Defined at top for simplicity
    username := ""
    password := ""
    instance := ""
    tableName := "incident"

    // Create a new UsernamePasswordCredential using account credentials
    cred := credentials.NewUsernamePasswordCredential(username, password)

    // Create a new client using previously created credential and instance
    client := servicenowsdkgo.NewClient(cred, instance)

    // Optional parameters
    params := &tableapi.TableRequestBuilderGetQueryParameters{
        Limit: int32(1),
    }

    //extend client with table name and parameters you'd like to use
    //
    // recieve and error and TableCollectionResponse
    response, err := client.Now().Table(tableName).POST(params)
    if err != nil {
        panic(err)
    }
}
```

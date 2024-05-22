
## Installation

formalco/go-sdk supports 2 last Go versions and requires support for
[Go modules](https://github.com/golang/go/wiki/Modules). So make sure to initialize a Go module:

```shell
go mod init github.com/myorg/formal
```

And then install formalco/go-sdk/sdk@v2:

```shell
go get github.com/formalco/go-sdk/sdk/v2
```

## Quickstart

### Code

```go
package main

import (
	"context"
	"fmt"
	"os"

	corev1 "buf.build/gen/go/formal/core/protocolbuffers/go/core/v1"
	"connectrpc.com/connect"
	"github.com/formalco/go-sdk/sdk/v2"
)

func main() {
	client := sdk.New("my-api-key")

	request := connect.NewRequest(&corev1.ListUsersRequest{
		Limit: 100,
	})

	response, err := client.UserServiceClient.ListUsers(context.Background(), request)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to list users: %s\n", err.Error())
		os.Exit(1)
	}

	for _, user := range response.Msg.Users {
		fmt.Println(user.GetId())
	}
}
```

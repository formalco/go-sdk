# Formal Go SDK

Official Go client for the Formal core control-plane API (ConnectRPC).

## Install

```bash
go get github.com/formalco/go-sdk/v3@latest
```

## Usage

```go
package main

import (
	"context"
	"os"

	formal "github.com/formalco/go-sdk/v3"
	corev1 "github.com/formalco/go-sdk/v3/core/v1"
)

func main() {
	client, err := formal.New(formal.WithAPIKey(os.Getenv("FORMAL_API_KEY")))
	if err != nil {
		panic(err)
	}
	_, err = client.UserServiceClient.CreateUser(context.Background(), &corev1.CreateUserRequest{
		// ...
	})
	_ = err
}
```

`New` defaults to `https://api.formal.ai`. Pass `WithBaseURL` for other hosts.

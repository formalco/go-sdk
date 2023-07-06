
## Installation

formal/go-sdk supports 2 last Go versions and requires support for
[Go modules](https://github.com/golang/go/wiki/Modules). So make sure to initialize a Go module:

```shell
go mod init github.com/my/repo
```

And then install formal/go-sdk/sdk@v1:

```shell
go get github.com/formal/go-sdk/sdk@v1
```

## Quickstart

```go
import "github.com/formal/go-sdk/sdk"

func main() {
    client := sdk.New("my-api-key")
	
	client.UserServiceClient.CreateUser(ctx, connect.NewRequest(&adminv1.CreateUserRequest{
            User: &types.User{
            FirstName: "Foo",
            LastName:  "Bar",
            Type:     "user",
            Name:      "foo",
            Email:     "foo@bar.com",
            Admin:    false,
        },
    }))
}
```

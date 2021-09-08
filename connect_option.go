package environ

import (
	"context"
	"github.com/ydb-platform/ydb-go-sdk/v3/connect"
)

func WithEnvironCredentials(ctx context.Context) connect.Option {
	return connect.WithCreateCredentialsFunc(FromEnviron)
}

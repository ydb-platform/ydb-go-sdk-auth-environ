package environ

import (
	"context"

	"github.com/ydb-platform/ydb-go-sdk/v3"
)

func WithEnvironCredentials(ctx context.Context) ydb.Option {
	return ydb.WithCreateCredentialsFunc(FromEnviron)
}

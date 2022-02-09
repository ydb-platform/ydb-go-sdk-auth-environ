package environ

import (
	"context"
	"os"

	"github.com/ydb-platform/ydb-go-sdk/v3"
	"github.com/ydb-platform/ydb-go-yc"
)

func WithEnvironCredentials(ctx context.Context) ydb.Option {
	if serviceAccountKeyFile, ok := os.LookupEnv("YDB_SERVICE_ACCOUNT_KEY_FILE_CREDENTIALS"); ok {
		return ydb.MergeOptions(
			yc.WithInternalCA(),
			yc.WithServiceAccountKeyFileCredentials(serviceAccountKeyFile),
		)
	}
	if os.Getenv("YDB_ANONYMOUS_CREDENTIALS") == "1" {
		return ydb.WithAnonymousCredentials()
	}
	if os.Getenv("YDB_METADATA_CREDENTIALS") == "1" {
		return ydb.MergeOptions(
			yc.WithInternalCA(),
			yc.WithMetadataCredentials(ctx),
		)
	}
	if accessToken, ok := os.LookupEnv("YDB_ACCESS_TOKEN_CREDENTIALS"); ok {
		return ydb.WithAccessTokenCredentials(accessToken)
	}
	return ydb.MergeOptions(
		yc.WithInternalCA(),
		yc.WithMetadataCredentials(ctx),
	)
}

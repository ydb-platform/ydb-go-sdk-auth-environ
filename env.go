package environ

import (
	"context"
	"os"
	"runtime"
	"strconv"
	"strings"

	"github.com/ydb-platform/ydb-go-sdk/v3"
	"github.com/ydb-platform/ydb-go-sdk/v3/credentials"
	"github.com/ydb-platform/ydb-go-yc"
	metadata "github.com/ydb-platform/ydb-go-yc-metadata"
)

// WithEnvironCredentials check environment variables and creates credentials by it
func WithEnvironCredentials() ydb.Option {
	c, err := environCredentials(osLookupEnv{}, true)
	if err != nil {
		return func(ctx context.Context, c *ydb.Driver) error {
			return err
		}
	}
	return ydb.WithCredentials(c)
}

type lookupEnv interface {
	LookupEnv(key string) (string, bool)
}

var _ lookupEnv = osLookupEnv{}

type osLookupEnv struct{}

func (o osLookupEnv) LookupEnv(key string) (string, bool) {
	return os.LookupEnv(key)
}

func stackRecord() string {
	function, file, line, _ := runtime.Caller(1)
	name := runtime.FuncForPC(function).Name()
	return name + "(" + fileName(file) + ":" + strconv.Itoa(line) + ")"
}

func fileName(original string) string {
	i := strings.LastIndex(original, "/")
	if i == -1 {
		return original
	}
	return original[i+1:]
}

func environCredentials(env lookupEnv, appendSourceInfo bool) (credentials.Credentials, error) {
	if serviceAccountKey, ok := env.LookupEnv("YDB_SERVICE_ACCOUNT_KEY_CREDENTIALS"); ok {
		if appendSourceInfo {
			return yc.NewClient(
				yc.WithServiceKey(serviceAccountKey),
				yc.WithSourceInfo(
					stackRecord()+"# YDB_SERVICE_ACCOUNT_KEY_CREDENTIALS",
				),
			)
		}
		return yc.NewClient(
			yc.WithServiceKey(serviceAccountKey),
		)
	}
	if serviceAccountKeyFile, ok := env.LookupEnv("YDB_SERVICE_ACCOUNT_KEY_FILE_CREDENTIALS"); ok {
		if appendSourceInfo {
			return yc.NewClient(
				yc.WithServiceFile(serviceAccountKeyFile),
				yc.WithSourceInfo(
					stackRecord()+"# YDB_SERVICE_ACCOUNT_KEY_FILE_CREDENTIALS",
				),
			)
		}
		return yc.NewClient(
			yc.WithServiceFile(serviceAccountKeyFile),
		)
	}
	if v, has := env.LookupEnv("YDB_METADATA_CREDENTIALS"); has && v == "1" {
		if appendSourceInfo {
			return yc.NewInstanceServiceAccount(
				metadata.WithInstanceServiceAccountCredentialsSourceInfo(
					stackRecord() + "# YDB_METADATA_CREDENTIALS",
				),
			), nil
		}
		return yc.NewInstanceServiceAccount(), nil
	}
	if accessToken, ok := env.LookupEnv("YDB_ACCESS_TOKEN_CREDENTIALS"); ok {
		if appendSourceInfo {
			return credentials.NewAccessTokenCredentials(
				accessToken,
				credentials.WithSourceInfo(
					stackRecord()+"# YDB_ACCESS_TOKEN_CREDENTIALS",
				),
			), nil
		}
		return credentials.NewAccessTokenCredentials(
			accessToken,
		), nil
	}
	if appendSourceInfo {
		return credentials.NewAnonymousCredentials(
			credentials.WithSourceInfo(
				stackRecord(),
			),
		), nil
	}
	if user, ok := env.LookupEnv("YDB_STATIC_CREDENTIALS_USER"); ok {
		if password, ok := env.LookupEnv("YDB_STATIC_CREDENTIALS_PASSWORD"); ok {
			if endpoint, ok := env.LookupEnv("YDB_STATIC_CREDENTIALS_ENDPOINT"); ok {
				return credentials.NewStaticCredentials(user, password, endpoint), nil
			}
		}
	}
	return credentials.NewAnonymousCredentials(), nil
}

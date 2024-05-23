package environ

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/ydb-platform/ydb-go-sdk/v3/credentials"
	yc "github.com/ydb-platform/ydb-go-yc"
)

var _ lookupEnv = lookupEnvMap(nil)

type lookupEnvMap map[string]string

func (m lookupEnvMap) LookupEnv(key string) (string, bool) {
	v, has := m[key]
	return v, has
}

func Test_environCredentials(t *testing.T) {
	f, err := os.CreateTemp("", "tmpfile-")
	require.NoError(t, err)
	t.Cleanup(func() {
		f.Close()
		os.Remove(f.Name())
	})
	key := `{
	   "id": "ajet18dckikcb17r",
	   "service_account_id": "an285h0re8nmm5u6",
	   "private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQCphJG/n0TPek7e\n+IBF1v8ts1svA5nA1HSPHMGzs2ic+jMzkCWK/cZrziQ4XACJp2u/HQo+8BB5TTll\n9w+u1JPPWB+9angJ2ZeyZJBshji11P33L4eBYjL41RGxVAxBVoZ4m8S6fnaaoi44\n0qkLuVYp0i89xI9b4kaQjWznOnZjygf7FMolcG6xjMfd0TONhh70tMGE8jIpz6l0\nKiDwfkQRHtWhBV2vRfZkLp8bxZb3m1AncxVwiBNIZBE1WOpObuXuBdtXYRdl6tqv\nYU143vtp5B1K7fi5frxHTrHJ0eBSce+ApIYcVY/BGZvcQYRYcfp5ob8mRFK0x9Ne\nXKdP1QpNAgMBAAECggEBAKEegbER3kMdSu2eX70BthRnzpkG4RJoXEjR2kHEmf/p\nxtSRgz1yKy+4BLEUsOYlHMHPBjf/0Iian52Mj3lqTvraFYDhOyqRiB9keejMQgaB\nZMmWgUK2ZSAg0+opqmRHqxPjab6gCHa9AxW7FDUyePTS81fQBTod/VfTM9kqte/t\nPjaR4Y8P74Ce2X/bRI4wap4j3FsYf2u8dxgt7hvMWN3O/gUpKTwn7hOSV70NxCwt\nFwjYUCMrVl/sjCP9YcMq0pzUfEuUs9hT01qIBaNIddCllzq8md7TdxdPwCS7XJEi\nn5jkxGKeHllCFY2sEG8nDBNbg5bGPQZSdh3Vdw4jTrkCgYEA5yzCovBxhLS38aqE\ntVMa9HsvKtvcQYT0rTtw8H3Uuew1txD46HYm25FqtwR464aImH7QnUkQlt6EfdLH\n0Dd+rTUZBaGZ0ibztHPg7taXM6AySzbL9sdeaKTZu2Ss3qvSQ1HuOmjYv/kfGE2j\n5BE1paiC/K/QO8NB/T6wlu819fsCgYEAu7jLGJjM6WjzZ60vMljPjph3e0sfrZqS\nuSFiOXliUFwQmVJyLI5pZIn0Q/S4ulxMBQWkbgecEslKhVoYJriM0phUcD8wm3UR\nfH7wELsxkyDh7jilJBFa9z4uq6kMgbPHR6iVNvcup9Hn9XlJkmPsnMco38OzwFsO\nfFsmp+p6tlcCgYEAzp7jGF9wFvyvrACMvMSawwmXDueT5bvANVV7jHfrOoI1QHqa\n/qsb8AP5LbuBmIGWdTZjnzE+8pnQMeXDUgdH4egjhTT7FypZiGBKGy8R1cLJMRC5\nHMj1SPKO6T8Cg8NvG1yPYQV1NaCkekRqx93Z5UbITLGXnNLYmFD/5OfJgyECgYBv\nYAVHk6jHtxfq5CqDYYPLo5QIF5s04ee5ZSAk32rAKM9EWFEbNGc9WkgNZY2QLCCC\nPkW/bk5gKwNGuRxpJMeQTwaSDjulkECOr7V5B8cy9qh1MTBxhMaGuGLyP/sGnQZX\n8qKNGPyaXwSTdKF89EI3Bkau9Cqarquahm4Z5BloQwKBgQCkQHZodhs4lJMiqASb\nQ5bBL5axSG1lPEOzuCo7DHXaDU/xtm16BrPcvFogIhOBe1lEn7M2cwDaIPOqRmrK\nhhVePVXOImErTDcwuQpsjmbaXhcKgkQEcTKxWssZHs3a6TgumJkcx+Fa16SHi2yZ\nYfvoGnE0K9GDwHVBQi57oOlhrA==\n-----END PRIVATE KEY-----\n"
	}`
	_, err = f.WriteString(key)
	require.NoError(t, err)

	for _, tt := range []struct {
		name        string
		env         lookupEnvMap
		credentials credentials.Credentials
	}{
		{
			name:        "nothing environment variables",
			env:         map[string]string{},
			credentials: credentials.NewAnonymousCredentials(),
		},
		{
			name: "anonymous credentials only",
			env: map[string]string{
				"YDB_ANONYMOUS_CREDENTIALS": "1",
			},
			credentials: credentials.NewAnonymousCredentials(),
		},
		{
			name: "access token credentials only",
			env: map[string]string{
				"YDB_ACCESS_TOKEN_CREDENTIALS": "test",
			},
			credentials: credentials.NewAccessTokenCredentials("test"),
		},
		{
			name: "static credentials only",
			env: map[string]string{
				"YDB_STATIC_CREDENTIALS_USER":     "user",
				"YDB_STATIC_CREDENTIALS_PASSWORD": "password",
				"YDB_STATIC_CREDENTIALS_ENDPOINT": "endpoint",
			},
			credentials: credentials.NewStaticCredentials("user", "password", "endpoint"),
		},
		{
			name: "incomplete static credentials",
			env: map[string]string{
				"YDB_STATIC_CREDENTIALS_USER":     "user",
				"YDB_STATIC_CREDENTIALS_PASSWORD": "password",
			},
			credentials: credentials.NewAnonymousCredentials(),
		},
		{
			name: "service account key file credentials only",
			env: map[string]string{
				"YDB_SERVICE_ACCOUNT_KEY_FILE_CREDENTIALS": f.Name(),
			},
			credentials: func() credentials.Credentials {
				c, err := yc.NewClient(yc.WithServiceFile(f.Name()))
				require.NoError(t, err)
				return c
			}(),
		},
		{
			name: "service account key body credentials only",
			env: map[string]string{
				"YDB_SERVICE_ACCOUNT_KEY_CREDENTIALS": key,
			},
			credentials: func() credentials.Credentials {
				c, err := yc.NewClient(yc.WithServiceKey(key))
				require.NoError(t, err)
				return c
			}(),
		},
		{
			name: "priority check: all known credentials",
			env: map[string]string{
				"YDB_ANONYMOUS_CREDENTIALS":                "1",
				"YDB_ACCESS_TOKEN_CREDENTIALS":             "test",
				"YDB_STATIC_CREDENTIALS_USER":              "user",
				"YDB_STATIC_CREDENTIALS_PASSWORD":          "password",
				"YDB_STATIC_CREDENTIALS_ENDPOINT":          "endpoint",
				"YDB_SERVICE_ACCOUNT_KEY_FILE_CREDENTIALS": f.Name(),
				"YDB_SERVICE_ACCOUNT_KEY_CREDENTIALS":      key,
			},
			credentials: func() credentials.Credentials {
				c, err := yc.NewClient(yc.WithServiceFile(f.Name()))
				require.NoError(t, err)
				return c
			}(),
		},
		{
			name: "priority check: anonymous vs token vs static vs service account key credentials",
			env: map[string]string{
				"YDB_ANONYMOUS_CREDENTIALS":           "1",
				"YDB_ACCESS_TOKEN_CREDENTIALS":        "test",
				"YDB_STATIC_CREDENTIALS_USER":         "user",
				"YDB_STATIC_CREDENTIALS_PASSWORD":     "password",
				"YDB_STATIC_CREDENTIALS_ENDPOINT":     "endpoint",
				"YDB_SERVICE_ACCOUNT_KEY_CREDENTIALS": key,
			},
			credentials: func() credentials.Credentials {
				c, err := yc.NewClient(yc.WithServiceKey(key))
				require.NoError(t, err)
				return c
			}(),
		},
		{
			name: "priority check: anonymous vs token vs static credentials",
			env: map[string]string{
				"YDB_ANONYMOUS_CREDENTIALS":       "1",
				"YDB_ACCESS_TOKEN_CREDENTIALS":    "test",
				"YDB_STATIC_CREDENTIALS_USER":     "user",
				"YDB_STATIC_CREDENTIALS_PASSWORD": "password",
				"YDB_STATIC_CREDENTIALS_ENDPOINT": "endpoint",
			},
			credentials: credentials.NewAccessTokenCredentials("test"),
		},
		{
			name: "priority check: anonymous vs static credentials",
			env: map[string]string{
				"YDB_ANONYMOUS_CREDENTIALS":       "1",
				"YDB_STATIC_CREDENTIALS_USER":     "user",
				"YDB_STATIC_CREDENTIALS_PASSWORD": "password",
				"YDB_STATIC_CREDENTIALS_ENDPOINT": "endpoint",
			},
			credentials: credentials.NewStaticCredentials("user", "password", "endpoint"),
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			credentials, err := environCredentials(tt.env, false)
			require.NoError(t, err)
			require.Equal(t, fmt.Sprintf("%v", tt.credentials), fmt.Sprintf("%v", credentials))
		})
	}
}

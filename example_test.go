package environ_test

import (
	"context"
	"fmt"
	"os"
	"time"

	environ "github.com/ydb-platform/ydb-go-sdk-auth-environ"
	"github.com/ydb-platform/ydb-go-sdk/v3"
	"github.com/ydb-platform/ydb-go-sdk/v3/query"
)

func Example_withEnvironCredentials() {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second)
	defer cancel()
	// force set up environment variable for example
	os.Setenv("YDB_ANONYMOUS_CREDENTIALS", "1")
	db, err := ydb.Open(ctx, "grpc://localhost:2136/local",
		environ.WithEnvironCredentials(),
	)
	if err != nil {
		panic(err)
	}
	defer db.Close(ctx) // cleanup resources

	var helloWorld string
	row, err := db.Query().QueryRow(ctx, "SELECT 'HELLO WORLD'u", query.WithIdempotent())
	if err != nil {
		panic(err)
	}
	if err := row.Scan(&helloWorld); err != nil {
		panic(err)
	}

	fmt.Println(helloWorld)
	// Output: HELLO WORLD
}

func Example_dsnParameterForUseEnvironCredentials() {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second)
	defer cancel()
	// force set up environment variable for example
	os.Setenv("YDB_ACCESS_TOKEN_CREDENTIALS", "")
	// For use go_environ_credentials parameter in connection string - need write blank import such as
	// import _ "github.com/ydb-platform/ydb-go-sdk-auth-environ"
	db, err := ydb.Open(ctx, "grpc://localhost:2136/local?use_env_credentials")
	if err != nil {
		panic(err)
	}
	defer db.Close(ctx) // cleanup resources

	var helloWorld string
	row, err := db.Query().QueryRow(ctx, "SELECT 'HELLO WORLD'u", query.WithIdempotent())
	if err != nil {
		panic(err)
	}
	if err := row.Scan(&helloWorld); err != nil {
		panic(err)
	}

	fmt.Println(helloWorld)
	// Output: HELLO WORLD
}

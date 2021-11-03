# ydb-go-sdk-auth-environ

> helpers to connect to YDB using environ 

## Installation <a name="Installation"></a>

```bash
go get -u github.com/ydb-platform/ydb-go-sdk-auth-environ
```

## Usage <a name="Usage"></a>

```go
import (
	env "github.com/ydb-platform/ydb-go-sdk-auth-environ"
)
...
    db, err := ydb.New(
        ctx,
        connectParams,
        env.WithEnvironCredentials(ctx), 
    )
    
```

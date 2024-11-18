# FS DB

[![Test](https://github.com/glebziz/fs_db/actions/workflows/test.yml/badge.svg)](https://github.com/glebziz/fs_db/actions/workflows/test.yml)
[![Lint](https://github.com/glebziz/fs_db/actions/workflows/lint.yml/badge.svg)](https://github.com/glebziz/fs_db/actions/workflows/lint.yml)
[![Coverage](https://codecov.io/gh/glebziz/fs_db/branch/master/graph/badge.svg?token=CIBKI0F59J)](https://codecov.io/gh/glebziz/fs_db/)
[![Go Reference](https://pkg.go.dev/badge/github.com/glebziz/fs_db.svg)](https://pkg.go.dev/github.com/glebziz/fs_db)
[![Go Report Card](https://goreportcard.com/badge/github.com/glebziz/fs_db)](https://goreportcard.com/report/github.com/glebziz/fs_db)

FS DB is a simple key-value database for storing files. FS DB has two clients that give you the option to
inline database logic into your application or run an external server and send data using grpc. 

## Install

```shell
go get -u github.com/glebziz/fs_db
```

## Usage

### Store interface

FS DB provides the following store key-value interface:
```go
type Store interface {
	// Set sets the contents of b using the key. 
	Set(ctx context.Context, key string, b []byte) error

	// SetReader sets the reader content using the key. 
	SetReader(ctx context.Context, key string, reader io.Reader) error

	// Get returns content by key. 
	Get(ctx context.Context, key string) ([]byte, error)

	// GetReader returns content as io.ReadCloser by key. 
	GetReader(ctx context.Context, key string) (io.ReadCloser, error)

	// GetKeys returns all keys from the database. 
	GetKeys(ctx context.Context) ([]string, error)

	// Delete delete content by key. 
	Delete(ctx context.Context, key string) error
}
```

### Transaction interface

FS DB provides the following transaction interface:
```go
type TxOps interface {
	// Commit commits the transaction. 
	Commit(ctx context.Context) error

	// Rollback rolls back the transaction. 
	Rollback(ctx context.Context) error
}
```

### DB interface

FS DB provides the following client interface:
```go
type DB interface {
	Store

	// Begin starts a transaction with isoLevel. 
	Begin(ctx context.Context, isoLevel ...model.TxIsoLevel) (Tx, error)
}
```

The database client supports starting a transaction with four standard isolation levels:
* `ReadUncommitted`
* `ReadCommitted`
* `RepeatableRead`
* `Serializable`

Since the set operation contains insert and update operations, the serializable level is equal to the repeatable read.

Use the following constants to select the isolation level:
```go
const (
	IsoLevelReadUncommitted
	IsoLevelReadCommitted
	IsoLevelRepeatableRead
	IsoLevelSerializable
)
```

### Tx interface

FS DB provides the following tx interface:
```go
type Tx interface {
	TxOps
	Store
}
```

### Example

See more examples [here](https://github.com/glebziz/fs_db/tree/master/example/).

```go
package main

import (
	"context"
	"fmt"
	"log"
	
	"github.com/glebziz/fs_db/pkg/inline"
)

func main() {
	db, err := inline.Open(context.Background(), config.Config{
		Storage: config.Storage{
			DbPath:      "test_db",
			MaxDirCount: 1,
			RootDirs:    []string{"./testStorage"},
			GCPeriod:    1 * time.Minute,
		},
		WPool: config.WPool{
			NumWorkers:   runtime.GOMAXPROCS(0),
			SendDuration: 1 * time.Millisecond,
		},
	})
	if err != nil {
		log.Fatalln("Open db inline:", err)
	}

	err = db.Set(context.Background(), "someKey", []byte("some content"))
	if err != nil {
		log.Fatalln("Set:", err)
	}

	b, err := db.Get(context.Background(), "someKey")
	if err != nil {
		log.Fatalln("Get:", err)
	}

	fmt.Println(string(b))
}
```

## License

[MIT](https://choosealicense.com/licenses/mit/)
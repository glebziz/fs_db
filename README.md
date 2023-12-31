# FS DB

[![Test](https://github.com/glebziz/fs_db/actions/workflows/test.yml/badge.svg)](https://github.com/glebziz/fs_db/actions/workflows/test.yml)
[![Coverage](https://codecov.io/gh/glebziz/fs_db/branch/master/graph/badge.svg?token=CIBKI0F59J)](https://codecov.io/gh/glebziz/fs_db/)

FS DB is a simple key-value database for storing files. FS DB has two clients that give you the option to
inline database logic into your application or run an external server and send data using grpc. 

## Install

```shell
go get -u github.com/glebziz/fs_db
```

## Usage

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
	db, err := inline.Open(context.Background(), &config.Storage{
		DbPath:      "test.db",
		MaxDirCount: 1,
		RootDirs:    []string{"./testStorage"},
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
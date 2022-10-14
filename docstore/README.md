# docstore
Document store with built in cache 

Supported storage backend
- Memory
- MongoDB
- SQL (Mysql, Postgres)

Plase see the cache package for supported cache backend

## Quick Start

Installation
    $ go get gitlab.sicepat.tech/platform/golib/docstore

## Usage

### MongoDB

```go

import (
    "fmt"

	"gitlab.sicepat.tech/platform/golib/docstore"
    _ "gitlab.sicepat.tech/platform/golib/docstore/mongo"
)


mconf := &docstore.Config{
    Database:        "userdb",
    Collection:      "user",
    CacheURL:        "mem://",
    CacheExpiration: 3600 * 24,
    IDField:         "id",
    TimestampField:  "created_at",
    Driver:          "mongo",
    Connection:      database.Client{
        URI:     "mongodb://localhost:27017",
        DB:      "userdb",
        AppName: "test",
    },
}


```

### SQL

```go

import (
    "fmt"

	"gitlab.sicepat.tech/platform/golib/docstore"
    _ "gitlab.sicepat.tech/platform/golib/docstore/sql"
)


sconf := &docstore.Config{
    Database:        "userdb",
    Collection:      "user",
    CacheURL:        "mem://",
    CacheExpiration: 3600 * 24,
    IDField:         "id",
    TimestampField:  "created_at",
    Driver:          "mysql",
    Connection:      database.DBConfig{
		MasterDSN:     "root:password@(localhost:3306)/outbox?parseTime=true",
		SlaveDSN:      "root:password@(localhost:3306)/outbox?parseTime=true",
		RetryInterval: 5,
		MaxIdleConn:   10,
		MaxConn:       5,
	},
}


```

### Initialize docstore


```go

import (
    "fmt"

	"gitlab.sicepat.tech/platform/golib/docstore"
)

func main() {
    conf := &docstore.Config{
		Database:        "userdb",
		Collection:      "user",
		CacheURL:        "mem://",
		CacheExpiration: 3600 * 24,
		IDField:         "id",
		TimestampField:  "created_at",
		Driver:          "memory",
		Connection:      nil,
	}

    store, err := docstore.New(conf)
    if err != nil {
        panic(err)
    }


    ms := docstore.NewMemoryStore("user", "id")
    cache := mem.NewMemoryCache()
    conf := &Config{
		IDField:        "id",
		TimestampField: "created_at",
	}

    store = NewDocstore(ms, cache, conf)

}

```

### Example

```go
package main

import (
    "fmt"

	"gitlab.sicepat.tech/platform/golib/cache/mem"
	"gitlab.sicepat.tech/platform/golib/docstore"
)

type User struct {
    ID        string    `json:"id"`
    Name      string    `json:"name"`
    Username  string    `json:"username"`
    Age       int       `json:"age"`
    CreatedAt time.Time `json:"created_at"`
}

func main() {
    ms := docstore.NewMemoryStore("user", "id")
    cache := mem.NewMemoryCache()
    conf := &Config{
		IDField:        "id",
		TimestampField: "created_at",
	}

    cs := NewDocstore(ms, cache, conf)

    usr := &User{
		Name:     "sahal",
		Username: "sahalzain",
		Age:      35,
	}

    if err := cs.Create(ctx, usr); err != nil {
        fmt.Println(err)
    }

    var doc User
	if err := cs.Get(ctx, usr.ID, &doc); err != nil {
        fmt.Println(err)
    }
}

```
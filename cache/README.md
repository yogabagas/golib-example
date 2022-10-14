# cache
cache abstraction for Golang

Supported cache
  - In Memory
  - Redis
  - LRU
  - Embed (Badger)

## Quick Start

Installation
    $ go get gitlab.sicepat.tech/platform/golib/cache


## Usage

Interface:
```
type Cache interface {
	Set(ctx context.Context, key string, value interface{}, expiration int) error
	Increment(ctx context.Context, key string, expiration int) (int64, error)
	Get(ctx context.Context, key string) ([]byte, error)
	GetObject(ctx context.Context, key string, doc interface{}) error
	GetString(ctx context.Context, key string) (string, error)
	GetInt(ctx context.Context, key string) (int64, error)
	GetFloat(ctx context.Context, key string) (float64, error)
	Exist(ctx context.Context, key string) bool
	Delete(ctx context.Context, key string) error
	RemainingTime(ctx context.Context, key string) int
	Close() error
}
```

Example:

```go
package main

import (
	cache "gitlab.sicepat.tech/platform/golib/cache"
	_ "gitlab.sicepat.tech/platform/golib/cache/mem"
	_ "gitlab.sicepat.tech/platform/golib/cache/redis"
	_ "gitlab.sicepat.tech/platform/golib/cache/lru"
	_ "gitlab.sicepat.tech/platform/golib/cache/embed"
)

func main() {
	// Use in-memory store
	memcache, _ := cache.New("mem://")
	rediscache, _ := cache.New("redis://<user>:<pass>@localhost:6379/prefix")
	lru, _ := cache.New("lru://local/1024")
	embed, _ := cache.New("embed://tmp/mydb")
}
```

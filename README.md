# GoLib

## _official go shared library

These helpers come from internal use cases and have started as an experiment to provide generic library.
it may not have 100% test coverage. so if you plan to put exclusively on runtime, you must have a good test suite.

## Installation

```sh
go get github.com/yogabagas/golib-example
```

#### Usage

```sh
import "github.com/yogabagas/golib-example/(feature_name)"
```

## Compatibility

This library aim to support two latest stable version of go. current minimum supported version is 1.17

## Features

#### api

rest client library with retry support mechanism

#### cache

cache driver. list supported cache:

- In Memory
- Redis
- LRU

#### config

config reader based on reflect. supported file

- yaml
- json

#### constant

consists of several lists of commonly used constants

#### custerr

helper to generate error logging field format

#### database

database driver connection helper. list supported database:

- MySQL
- PostgreSQL
- MongoDB

#### env

env value getter

#### event

Event Emitter library ([detail](https://gitlab.sicepat.tech/platform/golib/-/blob/master/event/README.md)). Supported driver

- Logger (sender & writer)
- Kafka (sender)
- MongoDB outbox (sender & writer)
- SQL outbox (sender & writer)

#### grpc

grpc server helper

#### httputil

http server utility based on [fasthttp](https://github.com/valyala/fasthttp) library

#### lock

Distributed Lock package. implementation:

- local
- redis
- etcd
- zookeeper

#### log

log writer and formater utility

#### monitor

monitoring and alerting toolkit

#### response

generic response builder utility

#### router

http router utility

#### slack

helper to post data on slack

#### taskworker {taskworker}

multiprocess taskworker

#### tracing

instrumental utility to conenct with open tracing

#### util

simple helper function. list:

- Value lookup
- Value matcher
- Value decoder
- UID Generator
- ImgProxy URL Generator
- Struct utility

## License

MIT
**Free Software, Hell Yeah!**

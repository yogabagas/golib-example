# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [1.7.0] - 2022-08-26
### Added
- add function name field option when call `WithField` in log package

## [1.6.0] - 2022-08-08
### Added
- add [httpprofiling](profiling/httpprofiling) package

## [1.5.0] - 2022-08-05
### Added
- add `Balancer` and  `BalancerConfig` options to event `Kafka` sender

## [1.4.2] - 2022-07-08
### Fixed
- fix `util.Lookup` may panic on nil pointer field 

## [1.4.1] - 2022-07-06
### Fixed
- update router context, which cause unable to get http params and writer from context  

## [1.4.0] - 2022-05-13
### Added
- add `Increment` on cache

## [1.3.3] - 2022-05-10
### Fixed
- monitor http counter register twice instead of grpc counter

## [1.3.2] - 2022-05-10
### Added
- add `SkipRegisterReflectionServer` on grpc library
### Fixed
- unable to register multiple servers using grpc library, due to conflict in register reflection server

## [1.3.1] - 2022-04-28
### Added
- add response total counter by path to http monitoring
- implement `http.Handler` on MyRouter
- add otel `trace_id` fields to `log.WithContext`
### Fixed
- remove duplicate monitor init
### Deprecated
- `MyRouter.HttpHandler` deprecated in favor of `MyRouter.ServeHttp`

## [1.3.0] - 2022-04-25
### Added
- add [strutil](util/strutil) package
- add `monitor.Init`, to create unique metrics per app
- add monitor metrics for grpc
- add monitor metrics for event consumer
- add `HttpRouter` method to router as replacement of `WrapperHandler`
- add `Handler` method to router to register `http.Handler`
### Changed
- migrate router trace integration from opentracing to opentelemetry
### Fixed
- fix monitor http histogram value doesn't used the right seconds value
### Removed
- removed router `WrapperHandler`, please use `router.HttpRouter` instead

## [1.2.1] - 2022-03-01
### Fixed
- fix event consumer not using specific topic and or group config for specific topic worker pool 
- fix taskworker may not work properly when the running task is panic

## [1.2.0] - 2022-01-26
### Added
- adding option to override log level of kafka event library for sender and listener

## [1.1.6] - 2022-01-18
### Fixed
- upgrade kafka-go library version to 0.4.26 because the previous version causing unexpected EOF [regression](https://github.com/segmentio/kafka-go/issues/814)

## [1.1.5] - 2022-01-18
### Added
- adding stop context to event consumer to gracefully stop until context timeout
- set kafka sender and listener logger to use golib log

### Changed
- set event consumer error log level to warn when error is context canceled on stop

# Backend implementation in Go

## Features

- Chi as HTTP/2 Go Web Framework.
- Endpoint to start a ride -> `POST /rides`.
- Endpoint to finish a ride -> `POST /rides/{id}/finish`.
- Endpoint with exposed application metrics in the Prometheus text format -> `GET /metrics`.
- Prometheus service to browse metrics -> `http://localhost:9090/graph`.
- Input/state validation:
  - Do not start a ride if user_id or vehicle_id are not provided.
  - Do not start a ride if the user or vehicle have another ride ongoing.
  - Do not finish a ride if the ride is already finished.
- Ride price calculation (initial unlocking price when the ride is started, plus the price per minute when it is finished).
  - Prices are treated as integers to avoid problems with floating point numbers.
- API documentation with Swagger.
- HTTP and service tests.
- ORM and DB migrations with go-rel.
- Data persistence in a Postgres database managed with Docker.
- Code live reload using Air.
- Code linting, formatting and testing on pre-commit.
- IDE configuration for Go in VS Code.
- Dependabot for dependency management.

## Installation

_Note: This document assumes that the application is running on a macOS machine._

Pre-requisites:

- Go 1.18+
- docker

Tools:

- Swag: `go install github.com/swaggo/swag/cmd/swag@latest`
- go-rel CLI:

```
brew tap go-rel/tap
brew install rel
```

Dependencies needed for pre-commit(optional):

- Linting: `brew install golangci-lint`
- Imports: `go install golang.org/x/tools/cmd/goimports@latest`
- Pre-commit itself: `brew install pre-commit`

Other optional tools:

- Air

## Running the application

1. Copy the example environment variable to a .env file.
2. Start the Postgres DB with docker: `docker compose up -d`.
3. Run the migrations: `rel migrate`.
4. Start the application:
   1. If you installed air: `air`
   2. Otherwise: `make all`

The application will be available at http://localhost:8080/

## Building the application

Just run `make build`. The output files will be located at `./bin/server`

## Architecture

```
├── api
│   ├── handlers
│   │   ├── errors.go
│   │   └── rides.go
│   └── http.go
├── bin
│   └── server
├── cmd
│   └── server
│       └── main.go
├── db
│   └── migrations
│       └── [migration file]
├── docs
│   └── docs.go
├── rides
│   ├── finish.go
│   ├── ride.go
│   ├── service.go
│   └── start.go
└── utils
│   └── utils.go
└── [other domain]
    ├── [entity a].go
    ├── [business logic].go
    └── service.go
```

The project follows some clear architecture principles that allow for a scalable and maintainable application. The architecture is modular, with loosely coupled dependencies separated by domain. In the case of our application the only domain folder is `rides`. When the application grows we could have other domains such as `users` or `vehicles`.

This `rides` domain folder contains the service, the use cases and the entity struct, so that this part of the application is loosely coupled(there are no shared components with other domain areas). This prevents cyclic dependencies and makes it easier to test the application.

We also have the principle of dependency inversion, which means that the application code should not depend on the infrastructure of choice. As we are making use of `go-rel` as ORM, we do not depend on the specific implementation of a DB to access the data(we could switch postgresql by sqlite and there would be no need to change the code). 

In addition, by defining the **service** interface, we ensure that we are also not tied to where the data is coming from. For example, if we would like store data in a redis DB, we would only have to change the implementation of the service, but the other parts of the application which use the service would not need to change, as the interface would not change.

## Testing

The code is tested using the [Go testing framework](https://golang.org/pkg/testing/) and the testify package for assertions.

Two types of tests have been written: handler tests and service tests.

The **handler tests** check that the response returned by our server corresponds to what is expected in each case. They also validate the response for wrong inputs or wrong states(e.g. trying to finish a ride which is already finished).

The **service tests** check that our business logic is working as intended. They also check the state of the database after each operation. For example they check:
- if the total ride price is calculated correctly.
- if the database is correctly updated when a ride is started or finished. 
- expected error messages if invalid operations are performed(e.g. trying to finish a ride that does not exist).

## Tools & settings

### Chi

[Chi](https://github.com/go-chi/chi) is a HTTP/2 Go Web Framework. Heavily maintained and adopted by the developer community.

### go-rel

[go-rel](https://go-rel.github.io/) is a ORM and DB migration framework for Go. It was chosen because it was built with testability in mind, and we do love tests!

The [go-rel CLI](https://go-rel.github.io/migration/#running-migration) is needed to run the migrations placed in `/db/migrations/`.

### Prometheus

A Prometheus image has been included in the docker compose configuration to serve the Prometheus UI and browse the available metrics. The available metrics include:

- The duration of the requests(with: code, handler, method).
- The count of the requests(with: code, handler, method).
- The size of the responses(with: code, handler, method).
- The number requests being handled concurrently at a given time a.k.a inflight requests (with: handler).

Metrics such as the average duration of the requests can be derived from the metrics above.

### swagger-ui

The application has a swagger-ui interface. It is available at http://localhost:8080/swagger/index.html.

Every endpoint is configured using [Declarative Comments Format](https://github.com/swaggo/swag/blob/master/README.md#declarative-comments-format).

We are using [Swag](https://github.com/swaggo/swag) to generate and serve the API documentation.

The documentation is automatically generated from the comments on each endpoint when the `swag init -d "cmd/server/,rides/,api/,api/handlers/"` command is ran.

### pre-commit

[pre-commit](https://pre-commit.com) is a tool that allows to perform operations before the commit is made by making use of git hooks. Here it is used for linting, formatting and testing the code. If any of the checks fails, the commit is aborted.

We are using a set of hooks for Go defined [here](https://github.com/dnephin/pre-commit-golang).

### air (live reload)

[Air](https://github.com/cosmtrek/air) is a really simple tool which provides live reloading of the project whenever the code changes. The configuration can be found in `.air.toml`.

### makefile

A set of make commands to facilitate the development flow. See the [makefile](makefile) to see the list of available commands.

### vs-code configuration

Inside the `.vscode` folder, there is a `settings.json` file. This file contains the configuration for the Go language for this IDE.

## Assumptions and improvements

- we assume that the start/finish times are proper data
- we assume that the user and vehicle ids are valid
- a CI/CD pipeline should be implemented in the future
- an error tracking tool(such as [Sentry](https://sentry.io/welcome/)) should be set up in the repo

### branch convention

As a CI/CD pipeline has not been implemented, we thought that the environment was not proper for `trunk-based development`. Therefore, we decided to use the older `gitflow` convention, which offers more control about code changes.

There is the `main` branch, which reflects the state of the deployed application.

We have the `develop` branch, with all the latest changes that are being developed.

And finally there are `feature branches`, which are used to develop new features. These branches are first merged to develop and when the cycle is complete, develop is merged to main.

### commit convention

The commits have been made using the [conventional commits style](https://www.conventionalcommits.org/en/v1.0.0/)

## List of resources

- [go-rel](https://go-rel.github.io/)
- [Chi](https://github.com/go-chi/chi)
- [Go testing framework](https://golang.org/pkg/testing/)
- [Testing in Go with Chi](https://www.newline.co/@kchan/testing-a-go-and-chi-restful-api-route-handlers-and-middleware-part-2--5efc9135)
- [Rest API with Chi and Go](https://www.newline.co/@kchan/building-a-simple-restful-api-with-go-and-chi--5912c411)
- [Dependency injection in Go](https://stackoverflow.com/questions/67944863/dependency-injection-in-go)
- [Chi middlewares](https://github.com/go-chi/chi#middlewares)
- [Mock](https://github.com/golang/mock)
- [Spies](https://github.com/nyarly/spies)
- [Using Mock for Golang testing](https://www.sobyte.net/post/2022-03/use-mock-to-test/)
- [Go by example](https://gobyexample.com/json)
- [Env variables in Go](https://dev.to/craicoverflow/a-no-nonsense-guide-to-environment-variables-in-go-a2f)
- [Go Playground](https://go.dev/play/)
- [Air](https://github.com/cosmtrek/air)
- [Swag](https://github.com/swaggo/swag)
- [pre-commit](https://pre-commit.com)
- [pre-commit hooks for Go](https://github.com/dnephin/pre-commit-golang)
- [conventional commits style](https://www.conventionalcommits.org/en/v1.0.0/)
- [Trunk-based development](https://www.atlassian.com/continuous-delivery/continuous-integration/trunk-based-development)
- [go-metrics](https://github.com/slok/go-http-metrics)
- [prometheus](https://prometheus.io/)
- [The RED method](https://www.weave.works/blog/the-red-method-key-metrics-for-microservices-architecture/)

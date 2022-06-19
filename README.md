# Backend implementation in Go

## Features

- Chi as HTTP/2 Go Web Framework.
- Endpoint to start a ride -> `POST /rides`.
- Endpoint to finish a ride -> `POST /rides/{id}/finish`.
- Input validation:
  - Do not start a ride if user_id or vehicle_id are not provided.
  - Do not start a ride if the user or vehicle have another ride ongoing.
  - Do not finish a ride if the ride is already finished.
- Ride price calculation (initial unlocking price when the ride is started, plus the price per minute when it is finished).
  - Prices are treated as integers to avoid problems with floating point numbers.
- API documentation with Swagger.
- Fully tested.
- ORM and DB migrations with go-rel.
- Data persistence in a Postgres database managed with Docker.
- Code live reload using Air.
- Code linting, formatting and testing on pre-commit.
- IDE configuration for Go in VS Code.

TODO:

- Metrics

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

## Testing

The code is fully tested using the [Go testing framework](https://golang.org/pkg/testing/) and the testify package for assertions.

Two types of tests have been written: handler tests and repository tests.

The **handler tests** check that the response returned by our server corresponds to what is expected in each case. They also validate the response for wrong inputs or wrong states(e.g. trying to finish a ride which is already finished).

The **repository tests** check the state of the database after each operation, for example they check if the database is correctly updated when a ride is started or finished. They also check for expected error messages if invalid operations are performed(e.g. trying to finish a ride that does not exist).

## Tools & settings

### Chi

[Chi](https://github.com/go-chi/chi) is a HTTP/2 Go Web Framework. Heavily maintained and adopted by the developer community.

### go-rel

[go-rel](https://go-rel.github.io/) is a ORM and DB migration framework for Go. It was chosen because it was built with testability in mind, and we do love tests!

The [go-rel CLI](https://go-rel.github.io/migration/#running-migration) is needed to run the migrations placed in `/db/migrations/`.

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
- a dependency management tool(such as [dependabot](https://docs.github.com/en/code-security/dependabot/dependabot-version-updates/about-dependabot-version-updates)) should be set up in the repo

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

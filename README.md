# auth0-POC

## Implementation Resources
* [SPA + API](https://auth0.com/docs/architecture-scenarios/spa-api)
* [Auth0 Golang](https://auth0.com/blog/authentication-in-golang/)

## Deps
* Go.

## Running the service
* Run `go get ./...` to install all dependencies.
* Run the command `go run main.go`.

## Features
* Universal Login.
* MFA(Provided by Auth0 From Dashboard).
* SPA + API Integration.
* JWT Middleware and basic verification of token.

## TODOS
* Authorization, RBAC
* Verify and validate id token
* Validate Scopes
* Data Migration

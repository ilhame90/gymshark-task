# Orders


## Description
Order sevice

## Tech stack
Service has written using Go programming language.
* [Echo](https://echo.labstack.com/) as a HTTP framework
* [Viper](https://github.com/spf13/viper) for config management
* [OpenAPI Client and Server Code Generator](https://github.com/deepmap/oapi-codegen) to generate HTTP handlers using [Open API Specification](https://swagger.io/specification/#:~:text=The%20OpenAPI%20Specification%20(OAS)%20defines,or%20through%20network%20traffic%20inspection.)
* [Gomock](https://github.com/golang/mock) for mocking dependencies
* [OpenTelemetry](https://opentelemetry.io/) to export traces
* Docker

## Installation
Clone the project

`git clone https://github.com/ilhame90/gymshark-task`

Install [OpenAPI Client and Server Code Generator](https://github.com/deepmap/oapi-codegen)

`make openapi-download`

Generate handlers

`make openapi-http`




## Usage
Set configurations. Project contains `.env` file to ease setup environment variables

`source .env`

Run service

`go run cmd/orders/main.go`

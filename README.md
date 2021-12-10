# ElianCodes-imageGenerator

This API is still a work in progress.

The API is written in Go, using the gin-gonic RESTful API framework.

The goal of this repo is to become a cloud function that will return a static image with the details you requested from the API

## Development

### Prerequisites

- install Go 1.17
- install Go dependencies with `go install`

### Local

- Install gowatch: `go get github.com/silenceper/gowatch`
- run gowatch: `gowatch`

running this command in your terminal should be enough to make it run locally

## Building

```bash
go run main.go
```

this will output a .exe (depending on your system)

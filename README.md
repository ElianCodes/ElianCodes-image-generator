# ElianCodes-imageGenerator

This API is still a work in progress.

The API is written in Go, using the gin-gonic RESTful API framework.

The goal of this repo is to become a cloud function that will return a static image with the details you requested from the API

## Usage

The API has 3 endpoints:

1. `/`
2. `/health`
3. `/generate`

### `/`

`GET` endpoint that returns a Hello World

### `/health`

`GET` endpoint for health checks. The endpoint just returns a simple `200 OK` response

### `/generate`

`POST` endpoint that returns a PNG image with the body parameters:

#### body

```json
{
    "title": "Elian Codes",
    "pageTitle": "How I automate SEO to fit my needs",
    "link": "www.elian.codes/blog"
}
```

- `title`: will be used as main title of the image
- `pageTitle`: serves as the title of the post or article
- `link`: yet to build

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

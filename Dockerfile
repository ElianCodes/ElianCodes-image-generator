FROM golang:1.17

WORKDIR /go/src/app
COPY . .

RUN ["go", "mod", "tidy"]
RUN ["go", "mod", "vendor"]

RUN ["go", "build", "main.go"]

EXPOSE 3000

CMD ["./main"]
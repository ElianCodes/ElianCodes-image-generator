FROM golang:1.17

WORKDIR /go/src/app
COPY . .

RUN ["go", "mod", "tidy"]
RUN ["go", "mod", "vendor"]

EXPOSE 3000

CMD ["go", "run", "main.go"]
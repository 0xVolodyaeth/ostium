FROM golang:1.20-alpine

WORKDIR /go/src/

COPY . /go/src/

RUN go build -o app cmd/main.go

ENTRYPOINT ["/go/src/app"]
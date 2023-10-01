FROM golang:1.21.1-alpine3.18 AS build

RUN apk add --update git
WORKDIR /go/src/github.com/erik-sostenes/receipt-processor-api
COPY  . .
RUN go get -d -v ./...
RUN  CGO_ENABLED=0 go build -o /go/bin/receipt-processor-api cmd/bootstrap/main.go


# Building image with the binary
FROM scratch
COPY --from=build /go/bin/receipt-processor-api /go/bin/receipt-processor-api
ENTRYPOINT ["/go/bin/receipt-processor-api"]
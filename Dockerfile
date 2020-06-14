FROM golang:alpine as builder
FROM alpine:3.7

WORKDIR /
ENV GOPATH /go
ENV GO111MODULE="on"
COPY go.mod go.sum ./
RUN go mod download

COPY . /go/src/github.com/tensaitensai/Time-US-api
RUN CGO_ENABLED=0 GOOS=linux go build -o TimeUS-api /go/src/github.com/tensaitensai/Time-US-api/main.go

RUN apk add mysql-client
RUN apk add --no-cache ca-certificates

COPY --from=builder /TimeUS-api /go/src/github.com/tensaitensai/Time-US-api/TimeUS-api
CMD ["./TimeUS-api"]

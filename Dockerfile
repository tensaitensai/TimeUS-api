FROM golang:alpine as builder
ENV APPDIR $GOPATH/src/github.com/tensaitensai/TimeUS-api
ENV GO111MODULE on
RUN \
    apk update --no-cache && \
    mkdir -p $APPDIR
ADD . $APPDIR/
WORKDIR $APPDIR
RUN go build -ldflags "-s -w" -o timeus-api cmd/main.go
RUN mv timeus-api /

FROM alpine
RUN apk add --no-cache ca-certificates
RUN apk add mysql-client
COPY --from=builder /timeus-api ./
ENTRYPOINT ["./timeus-api"]

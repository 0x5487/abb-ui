FROM golang:1.9 AS builder
COPY . /go/src/github.com/jasonsoft/abb-ui/
WORKDIR /go/src/github.com/jasonsoft/abb-ui/cmd/web/
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build
RUN rm -rf /go/src/github.com/jasonsoft/abb-ui/cmd/web/run.bat && \
    rm -rf /go/src/github.com/jasonsoft/abb-ui/cmd/web/app.yml && \
    rm -rf /go/src/github.com/jasonsoft/abb-ui/cmd/web/main.go

FROM alpine:3.6
RUN apk update && \
    apk upgrade && \
    apk add --no-cache curl && \
    rm -rf /var/cache/apk/* && \
    mkdir -p /jasonsoft/abb_web

COPY --from=builder /go/src/github.com/jasonsoft/abb-ui/cmd/web/ /jasonsoft/abb_web/
WORKDIR /jasonsoft/abb_web/
CMD ./web


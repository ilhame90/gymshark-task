FROM golang:1.21.1-alpine3.18 as build

WORKDIR /go/src/orders
COPY . .
ARG netrc
ARG release
ENV CGO_ENABLED=0 RELEASE=$release NETRC=$netrc


RUN echo $NETRC | base64 -d > ~/.netrc &&  \
    apk update && apk upgrade && \
    apk add --no-cache bash git openssh && \
    go clean --modcache && \
    go mod download && \
    go build -o app cmd/orders/main.go

FROM alpine

COPY --from=build /go/src/orders/app /usr/local/bin/app

EXPOSE 8080

ENTRYPOINT ["/usr/local/bin/app"]

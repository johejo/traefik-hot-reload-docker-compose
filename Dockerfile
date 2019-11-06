FROM golang:1.13-alpine AS base-builder
WORKDIR /build
ENV GO111MODULE=on

FROM base-builder AS hello-builder
RUN go get -u github.com/cosmtrek/air
COPY hello/ .
RUN go build -v -o hello github.com/johejo/traefik-test/hello
EXPOSE 9001

FROM alpine:3.10 AS base-app

FROM base-app AS hello
COPY --from=hello-builder /build/hello /hello
ENTRYPOINT ["/hello"]
EXPOSE 9001

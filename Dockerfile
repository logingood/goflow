FROM golang:alpine as builder
ARG LDFLAGS=""

RUN apk --update --no-cache add git build-base gcc

WORKDIR /build
COPY go.mod /build
COPY go.sum /build
RUN go mod download

COPY . /build

RUN go build -ldflags "${LDFLAGS}" -o goflow cmd/goflow/goflow.go

FROM alpine:latest
ARG src_dir

RUN apk update --no-cache && \
    adduser -S -D -H -h / flow
USER flow
COPY --from=builder /build/goflow /

ENTRYPOINT ["./goflow"]

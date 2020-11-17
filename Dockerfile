## Builder 
FROM golang:1.13.6 AS builder
ADD . /go/src/api
WORKDIR /go/src/api
RUN go get api && \
    go install && \
    go build -o engine main.go
EXPOSE 443

## Distribution 
FROM alpine:latest AS runtime
LABEL app="api" version="2.0" maintainer="Asromi rOmi"
ENV serverport="443" dbapp="MSSQL" dbcache="redis"
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
WORKDIR /go/src/api
EXPOSE 443
COPY --from=builder /go/src/api/engine /go/src/api
CMD ./engine 
FROM  golang:latest  as builder
WORKDIR  /workdir
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct
COPY  .  .
RUN go mod download
RUN  CGO_ENABLED=0 GOARCH=amd64 go build -o api

FROM alpine:latest
WORKDIR /workdir
COPY --from=builder /workdir .
EXPOSE 8080
ENTRYPOINT [ "./api" ]
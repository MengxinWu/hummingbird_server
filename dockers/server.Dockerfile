FROM golang:1.19-alpine AS builder
ARG main_path
ADD ../server /workdir
WORKDIR /workdir
ENV GO111MODULE=on GOPROXY=https://goproxy.cn,direct
RUN cd /workdir && pwd && ls -l && go mod download && go mod tidy
RUN cd /workdir && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /workdir/main

FROM alpine:3.16
COPY --from=builder /workdir/main /main
RUN chmod +x /main
ENV TZ=Asia/Shanghai
ENTRYPOINT ["/main"]
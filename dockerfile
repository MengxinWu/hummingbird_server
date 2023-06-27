FROM golang:1.19-alpine AS builder
ARG main_path
ADD . /server
WORKDIR /server
ENV GO111MODULE=on GOPROXY=https://goproxy.cn,direct
RUN cd /server && go mod download && go mod tidy
RUN cd /server && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /task/main

FROM alpine:3.16
COPY --from=builder /task/main /main
RUN chmod +x /main
ENV TZ=Asia/Shanghai
ENTRYPOINT ["/main"]
FROM golang:alpine AS builder
ENV GOPROXY=https://goproxy.cn

WORKDIR /app

ADD go.mod .
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o wechatbot main.go

FROM alpine

WORKDIR /app
COPY --from=builder /app/wechatbot .
COPY --from=builder /app/config.json .

CMD ["./wechatbot"]
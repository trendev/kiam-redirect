FROM golang:alpine as builder
WORKDIR /go/src/app
COPY main.go .
RUN go build -o redirect . 

FROM alpine
WORKDIR /app
COPY --from=builder /go/src/app/redirect /app/
EXPOSE 9000
CMD ["./redirect"]
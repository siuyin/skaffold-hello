FROM golang:1.11.5-alpine3.7 as builder
COPY vendor /go/src
COPY main.go .
#RUN go build -o /app *.go
RUN CGO_ENABLED=0 go build -o /app *.go

#FROM alpine:3.7
FROM scratch
CMD ["./app"]
COPY --from=builder /app .

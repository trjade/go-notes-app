FROM golang:1.22-alpine AS builder
WORKDIR /src
COPY . .
RUN go build -o /app ./

FROM alpine:3.18
COPY --from=builder /app /app
EXPOSE 8083
ENTRYPOINT ["/app"]
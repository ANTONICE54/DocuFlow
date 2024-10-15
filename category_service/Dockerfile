#Build stage
FROM golang:1.22.5-alpine3.19 AS builder
WORKDIR /app
COPY category_service .
RUN go build -o main ./cmd/main/main.go

#Run stage
FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/.env .
RUN mkdir migrations
COPY --from=builder /app/internal/database/migrations /app/migrations






EXPOSE 8081
CMD [ "/app/main" ]
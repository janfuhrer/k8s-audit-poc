FROM golang:1.21 AS builder

ENV CGO_ENABLED=0
WORKDIR /app
COPY ./audit-webhook ./audit-webhook
RUN go build -C audit-webhook

FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/audit-webhook/audit-webhook .
EXPOSE 8080
CMD ["./audit-webhook"]

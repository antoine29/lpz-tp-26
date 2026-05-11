FROM golang:1.26-alpine AS builder

WORKDIR /app
COPY go.mod go.sum* ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o adapter .

FROM scratch
COPY --from=builder /app/adapter /adapter
EXPOSE 8084
ENTRYPOINT ["/adapter"]
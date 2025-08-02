FROM golang:latest AS builder

WORKDIR /app
COPY api/icons/index.go /app/api/icons/index.go
COPY build.go /app/build.go
COPY server.go /app/server.go
COPY assets/ /app/assets/
COPY go.mod /app/go.mod
COPY go.sum /app/go.sum

RUN go mod tidy
RUN go run build.go

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -ldflags="-s -w" -o skillicons ./server.go

FROM scratch

COPY --from=builder /app/skillicons /skillicons

EXPOSE 8080

ENTRYPOINT ["/skillicons"]

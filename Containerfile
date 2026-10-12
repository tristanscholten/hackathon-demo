# syntax=docker/dockerfile:1
FROM docker.io/library/golang:1.27.1-alpine AS build
WORKDIR /src
COPY go.mod ./
RUN go mod download
COPY cmd ./cmd
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags='-s -w' -o /out/server ./cmd/server

FROM docker.io/library/alpine:3.23
LABEL org.opencontainers.image.source="https://github.com/tristanscholten/hackathon-demo"
RUN addgroup -S app && adduser -S -G app -u 10001 app
USER app
WORKDIR /app
COPY --from=build /out/server /app/server
EXPOSE 8080
ENTRYPOINT ["/app/server"]

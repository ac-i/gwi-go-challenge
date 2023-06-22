# [1] build
FROM docker.io/library/golang:1.20 AS build

WORKDIR /app-src
COPY . .

ENV CGO_ENABLED=0
RUN env && \
    go mod tidy && \
    go build -o /app-run

# [2] run
FROM scratch
LABEL app.version="0.1.0"

COPY --from=build /app-run /

ENTRYPOINT ["/app-run"]

EXPOSE 9090 9080
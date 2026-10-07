FROM golang:1.25.3-trixie AS build
WORKDIR /src
COPY . .
RUN go mod tidy && CGO_ENABLED=0 go build -trimpath -o /service ./cmd/candidate
FROM debian:trixie-slim
RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates && rm -rf /var/lib/apt/lists/*
COPY --from=build /service /usr/local/bin/candidate
USER 65532:65532
EXPOSE 8080
CMD ["candidate"]

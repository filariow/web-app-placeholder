FROM gcr.io/distroless/base as final
LABEL org.opencontainers.image.source https://github.com/filariow/web-app-placeholder
ENV GIN_MODE=release
ENV PORT=80
WORKDIR /app
CMD ["/app/webapp"]

FROM golang:1.17 as build-env
WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build \
        -trimpath \
        --ldflags='-s -w' \
        -o ./bin/webapp \
        main.go

FROM final
COPY --from=build-env /app/bin/webapp .


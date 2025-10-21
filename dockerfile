FROM golang:1.25.1 AS build

WORKDIR /src

COPY . .

RUN go build -o /app/main .

FROM debian:bookworm-slim

WORKDIR /app

COPY --from=build /app/main /app/main

EXPOSE 3030

ENTRYPOINT ["/app/main"]

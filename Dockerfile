FROM golang:1.22-alpine AS build

WORKDIR /app
COPY . .

RUN go build -o ytscrape -ldflags="-w -s" cmd/api/main.go

FROM alpine:latest AS run

RUN apk add --no-cache make

WORKDIR /app
COPY --from=build /app/ytscrape ./ytscrape

EXPOSE 8080

CMD ["./ytscrape"]

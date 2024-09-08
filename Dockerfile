FROM golang:1.22-alpine AS build

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /app/main .

FROM alpine:3.18

WORKDIR /app

COPY --from=build /app/main .
COPY static static

EXPOSE 3050

CMD ["./main"]
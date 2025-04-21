FROM golang:1.23.4-alpine3.20 AS build

WORKDIR /app

# Modules layer
COPY go.mod go.sum ./
RUN go mod download

# Build layer
COPY . .
RUN CGO_ENABLED=0 go build -o /myapp ./cmd/app

FROM alpine:3.20 AS run

COPY --from=build /myapp /myapp

EXPOSE 3000

CMD ["/myapp"]
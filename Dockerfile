# --------------------------------------------- Build the image
FROM golang:1.21.0-alpine3.18 as builder
LABEL authors="Ramin Farmani <ramin.farmani@gmail.com>"
LABEL maintainer="Ramin Farmani <ramin.farmani@gmail.com>"

RUN apk add --no-cache git
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o ./build/sharebuy

# --------------------------------------------- Run the image
FROM alpine:latest AS production
LABEL authors="Ramin Farmani <ramin.farmani@gmail.com>"
LABEL maintainer="Ramin Farmani <ramin.farmani@gmail.com>"

WORKDIR /app
COPY --from=builder /build/sharebuy .
ENTRYPOINT ["sharebuy", "api-v1"]
# Stage 1: compile bixi-around
FROM golang:1.19 as build-stage
WORKDIR /app
COPY . .
RUN go mod download
COPY bin .
RUN make build

# Stage 2: build the image
FROM alpine:latest
RUN apk --no-cache add ca-certificates libc6-compat
WORKDIR /app/
COPY --from=build-stage /app/bin .
ENV GIN_MODE=release
CMD ["./bixi-around"]
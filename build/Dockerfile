FROM golang:1.23 AS build-stage

WORKDIR /app


COPY go.mod go.sum ./
RUN go mod download

COPY . .


RUN CGO_ENABLED=0 GOOS=linux go build -o /wellknown ./cmd/well-known

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian12 AS build-release-stage

WORKDIR /

COPY --from=build-stage /wellknown /wellknown

EXPOSE 3000

USER nonroot:nonroot

ENTRYPOINT ["/wellknown"]
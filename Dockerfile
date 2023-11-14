FROM golang:1.21.1-bullseye AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -v -o /go-app .

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11

WORKDIR /

COPY --from=build /go-app /go-app

# exposes the specified port and makes it available only for inter-container communication
EXPOSE 3000

USER nonroot:nonroot

ENTRYPOINT ["/go-app"]

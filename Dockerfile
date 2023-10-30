FROM golang:1.21.1-bullseye AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download && go mod verify

COPY . .

RUN go build -v -o /my-app .

# FROM gcr.io/distroless/base-debian11

# COPY --from=build /my-app /my-app

RUN env

ENTRYPOINT ["/my-app"]


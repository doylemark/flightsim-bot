FROM golang:1.15-alpine as DEV

WORKDIR /flightsim-bot

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go install github.com/doylemark/flightsim-bot

CMD ["go", "run", "*.go"]

FROM alpine

WORKDIR /bin

COPY --from=DEV /go/bin/flightsim-bot ./flightsim-bot

CMD ["sh", "-c", "flightsim-bot -p"]
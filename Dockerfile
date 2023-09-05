FROM golang:1.21 AS builder

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

ENV GOPROXY=https://goproxy.io

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /tagfile

FROM gcr.io/distroless/base-debian10

WORKDIR /app

COPY --from=builder /tagfile /tagfile

USER root

VOLUME [ "/app" ]

ENTRYPOINT [ "/tagfile" ]

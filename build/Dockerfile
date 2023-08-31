# BUILD
FROM golang:1.20.2-alpine AS builder 

WORKDIR /app 

COPY go.mod /app/

RUN go mod download

COPY . /app/

RUN go build -o /app/build/amhooker


## FINAL
FROM alpine:3.14 

WORKDIR / 

COPY --from=builder /app/build/amhooker /amhooker

ENTRYPOINT ["/amhooker"]
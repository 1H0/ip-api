FROM golang:1.19-alpine as build

WORKDIR /app

COPY go.mod ./
COPY *.go ./

RUN go build -o /out

FROM alpine:3.16

WORKDIR /

COPY --from=build /out /app

ENTRYPOINT ["/app"]


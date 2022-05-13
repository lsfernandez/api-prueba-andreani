FROM golang:alpine as builder

EXPOSE 8080

RUN apk update && apk add --no-cache git

RUN mkdir /build 
COPY go.mod /build/
COPY go.sum /build/
WORKDIR /build
RUN go mod download
ADD . /build/
RUN go build -o api-prueba .

FROM alpine:latest
EXPOSE 8080
RUN apk --no-cache add ca-certificates tzdata && \
    cp /usr/share/zoneinfo/America/Argentina/Buenos_Aires /etc/localtime && \
    apk del tzdata && rm -rf /var/cache/apk/* && date

WORKDIR /app
COPY --from=builder /build/api-prueba /app/

CMD ["./api-prueba"]

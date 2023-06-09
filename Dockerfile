FROM golang:1.18.2-alpine3.16 as base
RUN apk update 
WORKDIR /src/go-pay
COPY .env .
ADD . . 
RUN go mod download
RUN go build -o go-pay ./cmd

FROM alpine:3.16 as binary
WORKDIR /src/app
COPY --from=base /src/go-pay/* .
EXPOSE 8080
CMD ["/src/app/go-pay"]
FROM alpine:3.12 AS builder
MAINTAINER Roshan Bhatia (rshnbhatia@gmail.com) 

RUN apk update && apk add openssl
RUN openssl req -x509 -nodes -days 365 -subj "/C=CA/ST=QC/O=Company Inc/CN=example.com" -newkey rsa:2048 -keyout /etc/ssl/private/localhost.key -out /etc/ssl/certs/localhost.crt

FROM golang:1.15.6-alpine
ARG SERVICE_PORT


COPY ./bin/echo-service ./echo-service
RUN chmod +x echo-service

COPY --from=builder /etc/ssl/certs/localhost.crt /etc/ssl/certs/localhost.crt
COPY --from=builder /etc/ssl/private/localhost.key /etc/ssl/private/localhost.key

ENV SSL_CERT_PATH /etc/ssl/certs/localhost.crt
ENV SSL_KEY_PATH /etc/ssl/private/localhost.key
ENV SERVICE_PORT $SERVICE_PORT

EXPOSE $SERVICE_PORT
EXPOSE 443
ENTRYPOINT [ "./echo-service" ]

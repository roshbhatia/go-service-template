FROM alpine:3.12 AS builder
MAINTAINER Roshan Bhatia (rshnbhatia@gmail.com) 

RUN apk update && apk add openssl
RUN openssl req -x509 -nodes -days 200 -subj "/C=CA/ST=QC/O=Company Inc/CN=127.0.0.1" -newkey rsa:2048 -keyout /etc/ssl/private/localhost.key -out /etc/ssl/certs/localhost.crt

FROM golang:1.15.6-alpine
ARG SERVICE_PORT


COPY ./bin/echo-service ./echo-service
RUN chmod +x echo-service

COPY --from=builder /etc/ssl/certs/localhost.crt localhost.crt
COPY --from=builder /etc/ssl/private/localhost.key localhost.key

ENV SSL_CERT_PATH localhost.crt
ENV SSL_KEY_PATH localhost.key
ENV SERVICE_PORT $SERVICE_PORT

EXPOSE $SERVICE_PORT
ENTRYPOINT [ "./echo-service" ]

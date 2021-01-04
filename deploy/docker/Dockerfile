FROM golang:1.15.6-alpine
COPY ./bin/echo-service ./echo-service
RUN chmod +x echo-service

ENTRYPOINT [ "./echo-service" ]
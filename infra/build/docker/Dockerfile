FROM golang:1.18.0-alpine

WORKDIR /apt/app

EXPOSE 8080

COPY . .

USER root

RUN apk -U add libc-dev libc6-compat

USER appuser

CMD ["./app"]

# build : docker build -t app .
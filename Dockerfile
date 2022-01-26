FROM golang:alpine

WORKDIR /app

RUN adduser --disabled-password golang \
  && chown golang:golang -R /app

USER golang

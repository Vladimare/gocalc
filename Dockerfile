FROM golang:1.19 AS build

ADD . /usr/src/gocalc

WORKDIR /usr/src/gocalc
RUN go build .

FROM ubuntu:20.04

COPY --from=build /usr/src/gocalc/gocalc /usr/bin/gocalc

WORKDIR /usr/bin

ENTRYPOINT ["gocalc"]

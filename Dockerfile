FROM golang:alpine AS base

WORKDIR /root

COPY *.go .
COPY go.* ./

RUN go build .

FROM alpine
RUN mkdir /input
WORKDIR /input
VOLUME  /input

COPY --from=base /root/mars_rover /usr/local/bin/mars_rover

ENTRYPOINT ["/usr/local/bin/mars_rover"]


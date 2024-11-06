FROM golang:alpine AS base

WORKDIR /root

COPY *.go .
COPY go.* ./

RUN go build .

FROM alpine

COPY --from=base /root/mars_rover /root/mars_rover
COPY input.txt .

ENTRYPOINT ["/root/mars_rover"]

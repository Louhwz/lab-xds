#FROM golang:latest as builder
#
#WORKDIR /builder
#ADD main.go go.mod ./
#RUN go build -o main

FROM alpine:latest
ADD main /bin
EXPOSE 80
CMD ["/bin/main"]
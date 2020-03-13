FROM golang:1.14.0-alpine as builder

RUN apk --no-cache add ca-certificates gcc libc-dev nodejs npm git make

VOLUME /root/go
COPY ./ /server
RUN cd /server &&\
 npm install -g jshint rollup &&\
 make test jshint all

FROM alpine:3.11.3
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /server/output/mapserver-linux-x86_64 /bin/mapserver

EXPOSE 8080
CMD ["/bin/mapserver"]

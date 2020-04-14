FROM golang:1.14.2-alpine as builder

RUN apk --no-cache add ca-certificates gcc libc-dev nodejs npm git make

VOLUME /root/go
COPY ./ /server
RUN cd /server &&\
 npm install -g jshint rollup &&\
 make test jshint all

FROM alpine:3.11.5
RUN apk --no-cache add ca-certificates curl
WORKDIR /app
COPY --from=builder /server/output/mapserver-linux-x86_64 /bin/mapserver

HEALTHCHECK --interval=15s --timeout=3s \
  CMD curl -f http://localhost:8080/ || exit 1

EXPOSE 8080
CMD ["/bin/mapserver"]

FROM node:22.16.0 as bundle-builder
COPY public /public
WORKDIR /public
RUN npm ci && npm run bundle

FROM golang:1.24.3 as go-builder
COPY . /data
COPY --from=bundle-builder /public/js/bundle* /data/public/js/
WORKDIR /data
RUN CGO_ENABLED=0 go build .

FROM alpine:3.21.3
COPY --from=go-builder /data/mapserver /bin/mapserver
ENV MT_CONFIG_PATH "mapserver.json"
ENV MT_LOGLEVEL "INFO"
ENV MT_READONLY "false"
EXPOSE 8080
ENTRYPOINT ["/bin/mapserver"]
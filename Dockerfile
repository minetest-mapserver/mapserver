FROM ubuntu as builder

RUN apt-get update &&\
	apt-get install -y software-properties-common git &&\
	add-apt-repository ppa:longsleep/golang-backports &&\
	apt-get update &&\
	apt-get install -y golang-go npm nodejs git

VOLUME /root/go
COPY ./ /server
RUN cd /server &&\
 npm install -g jshint rollup &&\
 make test jshint all

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /server/mapserver /bin/mapserver

EXPOSE 8080
CMD ["/bin/mapserver"]

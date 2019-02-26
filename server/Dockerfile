FROM ubuntu as builder

RUN apt-get update &&\
	apt-get install -y software-properties-common git &&\
	add-apt-repository ppa:longsleep/golang-backports &&\
	apt-get update &&\
	apt-get install -y golang-go

COPY ./ /server
RUN cd /server && make build

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /server/output/mapserver-linux-x86_64 /bin/mapserver
CMD ["/bin/mapserver"]


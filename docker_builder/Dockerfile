FROM ubuntu:bionic

RUN apt-get update &&\
	apt-get install -y gcc-mingw-w64 gcc-5-arm-linux-gnueabihf gcc-i686-linux-gnu &&\
	apt-get install -y software-properties-common git

RUN add-apt-repository ppa:longsleep/golang-backports &&\
	apt-get update &&\
	apt-get install -y golang-go

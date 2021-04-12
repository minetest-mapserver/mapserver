FROM ubuntu:hirsute

# for tzdata
ENV DEBIAN_FRONTEND=noninteractive

RUN apt-get update

# cross compile and go stuff
RUN apt-get install -y gcc-mingw-w64 gcc-8-arm-linux-gnueabihf gcc-i686-linux-gnu &&\
	apt-get install -y software-properties-common git golang-go

# TODO: move this to a separate package.json and docker container
# jshint
RUN apt-get install -y nodejs npm
RUN npm install -g jshint@2.12.0

# rollup
RUN npm install -g rollup@2.35.1

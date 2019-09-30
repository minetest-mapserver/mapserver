FROM ubuntu:bionic

# cross compile stuff
RUN apt-get update &&\
	apt-get install -y gcc-mingw-w64 gcc-5-arm-linux-gnueabihf gcc-i686-linux-gnu &&\
	apt-get install -y software-properties-common git

# go stuff
RUN add-apt-repository ppa:longsleep/golang-backports &&\
	apt-get update &&\
	apt-get install -y golang-go

# luacheck
RUN apt-get install -y luarocks liblua5.1-dev
RUN luarocks install luacheck

# jshint
RUN apt-get install -y nodejs npm
RUN npm install -g jshint

# rollup
RUN npm install -g rollup

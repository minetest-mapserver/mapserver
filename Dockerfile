#== Container for running rolllup. That's all.
FROM node:22-alpine as rollup

RUN npm install --global rollup

COPY . /src
WORKDIR /src

WORKDIR /src/public
RUN rollup -c rollup.config.js

#== The runtime golang container. This is the one to be exported.
FROM golang:1.22-alpine as runtime

# Get the rolled up files
COPY --from=rollup /src /src
WORKDIR /src

# Build the binary
RUN go build

# Keep backward compatibility
RUN ln -s ../src/mapserver /bin/mapserver

# Set up default env variables
ENV MT_CONFIG_PATH "mapserver.json"
ENV MT_LOGLEVEL "INFO"
ENV MT_READONLY "false"

# Final definitions
EXPOSE 8080
ENTRYPOINT ["/src/mapserver"]

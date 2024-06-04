#== Define versions here.
ARG ALPHINE_VER=3.20
ARG NODE_VER=22.2
ARG GO_VER=1.21

#== Container for running rollup. That's all.
FROM node:${NODE_VER}-alpine${ALPHINE_VER} as rollup

RUN npm install --global rollup

COPY ./public /public
WORKDIR /public

RUN rollup -c rollup.config.js

#== The container building Go codes.
FROM golang:${GO_VER}-alpine${ALPHINE_VER} AS build

# Get the rolled up files
COPY . /src
COPY --from=rollup /public/js/bundle.js /src/public/js/bundle.js
COPY --from=rollup /public/js/bundle.js.map /src/public/js/bundle.js.map
WORKDIR /src

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build

#== Run the tests
# Use this command to ensure it runs:
## docker build . --progress plain --no-cache --target run-test
FROM build AS run-test
RUN go test -v ./...

#== Export the image
FROM scratch AS release

# Copy the binary and license
COPY --from=build /src/mapserver /bin/mapserver
COPY license.txt license_mapserver.txt

# Set up default env variables
ENV MT_CONFIG_PATH "mapserver.json"
ENV MT_LOGLEVEL "INFO"
ENV MT_READONLY "false"

# Final definitions
EXPOSE 8080
ENTRYPOINT ["/bin/mapserver"]

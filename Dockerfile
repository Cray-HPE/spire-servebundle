########### Build ##########
FROM arti.dev.cray.com/baseos-docker-master-local/golang:alpine3.12 AS build

RUN apk add --no-cache git build-base

RUN mkdir -p /build
COPY . /build
RUN cd /build && go test -v ./...
RUN cd /build && go build -o /usr/local/bin/spire-bundle
RUN ls /build
RUN ls /usr/local/bin

########## Runtime ##########
FROM arti.dev.cray.com/baseos-docker-master-local/alpine:3.13.2 AS runtime

RUN apk add --no-cache bash curl

COPY --from=build /usr/local/bin/spire-bundle /usr/local/bin/spire-bundle
COPY ./.version /bundle-version
EXPOSE 8080/tcp
ENTRYPOINT ["/usr/local/bin/spire-bundle"]

## backend build-phase
## --------------------------------------------------------------------------
FROM golang:alpine AS BACKEND-BUILD

ARG buildtime_variable_version=1.0.0
ENV VERSION=${buildtime_variable_version}

WORKDIR /simpleapi
COPY ./go.mod ./
COPY ./go.sum ./
COPY ./main.go ./
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-s -w -X main.Version=${VERSION}" -o simpleapi ./*.go
## --------------------------------------------------------------------------

## runtime
## --------------------------------------------------------------------------
FROM alpine:latest
LABEL author="Henrik Binggl"
LABEL description="simple example of a pub-sub golang component for a dapr demo"

WORKDIR /opt/simpleapi

RUN mkdir -p /opt/simpleapi/
COPY --from=BACKEND-BUILD /simpleapi/simpleapi /opt/simpleapi

EXPOSE 3000

# Do not run as root user
## alpine specific user/group creation
RUN addgroup -S simpleapi && \
    adduser -S -D -H -u 1000 -G simpleapi -h /opt/simpleapi simpleapi

RUN chown -R simpleapi:simpleapi /opt/simpleapi
USER simpleapi

CMD ["/opt/simpleapi/simpleapi"]

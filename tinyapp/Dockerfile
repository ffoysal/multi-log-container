#multi stage docker file

# build stage
FROM golang:alpine AS build-env

ADD main.go /app/

RUN cd /app && go build -o tinyapp

# Final stage
FROM alpine

WORKDIR /go-app

COPY --from=build-env /app/tinyapp /go-app/
EXPOSE 8080
# application logs will be in /go-app/logs folder

ENTRYPOINT ./tinyapp


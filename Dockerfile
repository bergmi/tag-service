FROM golang:alpine

# Install build tools
# Git is required for fetching the dependencies.
RUN apk update && apk --no-cache add make git gcc libtool musl-dev ca-certificates dumb-init 

# RUN apk add build-base

WORKDIR /go/src/tag-service
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["go", "run", "src/server.go"]
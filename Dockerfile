# maintainer

LABEL Maintainer="Luca Paterlini <paterlini.luca@gmail.com>"

# building docker
FROM golang:alpine as builder

RUN mkdir /build 
ADD . /build/
WORKDIR /build
RUN apk add --no-cache git
ENV GOBIN /go/bin

# Force the go compiler to use modules
ENV GO111MODULE=on

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

## install all the go modules required
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main .

# execution docker
FROM scratch
COPY --from=builder /build/main /app/
WORKDIR /app
CMD ["./main"]

# Opening the ports for the service
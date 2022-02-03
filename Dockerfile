FROM golang:alpine

RUN apk update && apk add --no-cache git

WORKDIR /home/icemessanger/chatnews-restapi/

COPY . .

RUN go mod tidy

RUN go build -o binary

ENTRYPOINT ["/home/icemessanger/chatnews-restapi/binary"]
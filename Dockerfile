FROM golang:latest

ADD . /go/src/filechallenge
WORKDIR $GOPATH/src/filechallenge

RUN go get gorm.io/gorm
RUN go get gorm.io/driver/postgres

RUN go build -o /go/bin/filechallenge

ENTRYPOINT /go/bin/filechallenge

EXPOSE 8080
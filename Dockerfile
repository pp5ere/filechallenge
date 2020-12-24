FROM golang:latest

ADD . /go/src/desafioNeoWay
WORKDIR $GOPATH/src/desafioNeoWay

RUN go get gorm.io/gorm
RUN go get gorm.io/driver/postgres

RUN go build -o /go/bin/challengefile

ENTRYPOINT /go/bin/challengefile

EXPOSE 8080
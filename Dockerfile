FROM golang:1.14rc1-alpine3.11

WORKDIR /go/src/app
COPY . /go/src/app

RUN go get -d -v ./...
RUN go install -v ./...

RUN go build

EXPOSE 80

CMD ["./listable-backend"]

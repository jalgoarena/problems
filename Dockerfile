FROM golang:1.10.1

ENV GOBIN /go/bin

RUN mkdir /app
RUN mkdir /go/src/app
ADD . /go/src/app
WORKDIR /go/src/app

RUN go get -u github.com/golang/dep/...
RUN dep ensure

EXPOSE 8080

RUN go build -o /app/main .
CMD ["/app/main"]

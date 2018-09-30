FROM golang:1.11 as builder

ENV GOBIN /go/bin

RUN mkdir /app
RUN mkdir /go/src/app
ADD . /go/src/app
WORKDIR /go/src/app

RUN go get -u github.com/gobuffalo/packr/packr
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
RUN dep ensure

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 packr build -a -installsuffix cgo -ldflags="-w -s" -o /app/main .

FROM scratch

COPY --from=builder /app/main /app/main

EXPOSE 8080

CMD ["/app/main"]

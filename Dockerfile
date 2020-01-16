FROM golang:latest

ENV GOBIN $GOPATH/bin

RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o main .
CMD ["app/main"]

FROM golang:1.14.2-alpine3.11

RUN mkdir /covid19_api_consumer
ADD . /covid19_api_consumer
WORKDIR /covid19_api_consumer
RUN go build cmd/covid19_api_consumer/main.go
CMD ["/covid19_api_consumer/main"]

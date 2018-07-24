FROM golang

ADD . /go/src/github.com/tsmcalister/cc-cedict-microservice
RUN go get github.com/boltdb/bolt
RUN go get github.com/gorilla/mux
RUN cd src/github.com/tsmcalister/cc-cedict-microservice && ./dictionary_bootstrapper.sh && go build main.go
WORKDIR "/go/src/github.com/tsmcalister/cc-cedict-microservice/"
ENTRYPOINT ["./main"]
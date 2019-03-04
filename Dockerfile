FROM golang

RUN go get github.com/stripe/stripe-go

RUN go get github.com/gorilla/mux

WORKDIR /go/src/github.com/heaptracetechnology/microservice-stripe

ADD . /go/src/github.com/heaptracetechnology/microservice-stripe

RUN go install github.com/heaptracetechnology/microservice-stripe

ENTRYPOINT microservice-stripe

EXPOSE 3000
FROM golang

ADD . /go/src/github.com/ykim135/DateApi
RUN go install github.com/ykim135/DateApi
ENTRYPOINT /go/bin/DateApi

FROM golang

RUN mkdir -p /go/src/om-kit

ADD . /go/src/om-kit

RUN go get  -t -v ./...
RUN go get  github.com/canthefason/go-watcher
RUN go install github.com/canthefason/go-watcher/cmd/watcher

ENTRYPOINT  watcher -run om-kit/education/cmd  -watch om-kit/education

FROM golang:1.15.7-alpine3.13
 RUN apk update && apk upgrade && \
     apk add --no-cache bash git openssh
 ENV GIT_TERMINAL_PROMPT=1
 RUN go get -u github.com/sorokinmax/apicli
 RUN GIT_COMMIT=$(git rev-list -1 HEAD)
 ENV GO111MODULE=on
 ENV GOFLAGS=-mod=vendor
 ENV APP_HOME /go/src/github.com/sorokinmax/apicli/src
 ENV PATH "$PATH:/go/src/github.com/sorokinmax/apicli/src"
 WORKDIR $APP_HOME
 EXPOSE 89
 RUN go build -ldflags="-X 'main.version=v1.0.0'"
 CMD apicli
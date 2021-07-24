FROM golang:1.16.5:latest

PULL 
Run mkdir / build
WORKDIR /build

RUN export GO111MODULE=on
RUN go get 
EXPOSE 8080
ENTRYPOINT ["/build/DEBBIE/sca.go"]

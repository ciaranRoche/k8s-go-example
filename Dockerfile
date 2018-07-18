FROM golang:1.7.4-alpine
MAINTAINER ciaranRoche

ENV SOURCES /go/src/github.com/ciaranRoche/cloud-native-go/

COPY . ${SOURCES}
RUN cd ${SOURCES} && CGO_ENABLED=0 go install

ENTRYPOINT cloud-native-go
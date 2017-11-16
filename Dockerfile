FROM golang:onbuild

WORKDIR /go/src/github.com/RedPatchTechnologies/postmurum-server/
COPY . .

RUN go get -v -u github.com/serenize/snaker
#RUN go get -v -u github.com/golang/lint/golint
RUN go get -v -u github.com/markbates/pop/...
RUN go install -v
RUN go get github.com/codegangsta/gin


EXPOSE 8080
#CMD gin run
FROM golang:latest 
WORKDIR /go/src/github.com/RedPatchTechnologies/postmurum-server/
RUN pwd
COPY . .

RUN go get -v -u github.com/serenize/snaker
#RUN go get -v -u github.com/golang/lint/golint
RUN go get -v -u github.com/markbates/pop/...
RUN go install -v
RUN go get github.com/codegangsta/gin

EXPOSE 3000
CMD gin run
FROM golang:onbuild

ADD . /go/src/github.com/RedPatchTechnologies/postmurum-server/
WORKDIR /go/src/github.com/RedPatchTechnologies/postmurum-server/
COPY . .

RUN go get -v -u github.com/serenize/snaker
#RUN go get -v -u github.com/golang/lint/golint
RUN go get -v -u github.com/markbates/pop/...

RUN go install -v
#

ENTRYPOINT /go/bin/postmurum-server

EXPOSE 3000



#if [ "$RUN_GIN_CMD" = "true" ] ; 
#then CMD ["gin"];
#fi
#RUN if [ "$RUN_GIN_CMD" = "true" ] ; then EXPOSE 3000; else EXPOSE 8080; fi
#
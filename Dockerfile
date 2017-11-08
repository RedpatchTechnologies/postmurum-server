FROM golang:latest 
RUN mkdir /app 
ADD . /app/ 
ADD handlers /src/github.com/RedPatchTechnologies/server/handlers
ADD middleware /src/github.com/RedPatchTechnologies/server/middleware
WORKDIR /app 
RUN go get ./
RUN go build -o main . 
CMD ["/app/main"]
FROM golang:1.8

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["app"]


#https://hub.docker.com/_/golang/
#You can then build and run the Docker image:
#
#$ docker build -t my-golang-app .
#$ docker run -it --rm --name my-running-app my-golang-app


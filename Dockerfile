FROM golang:1.19
EXPOSE 9090

WORKDIR /prj

COPY . /prj/
RUN apt-get install gcc
RUN go build main.go
CMD ./main
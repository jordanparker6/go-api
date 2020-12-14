FROM golang:latest

RUN apt-get update && apt-get -y upgrade
RUN apt-get -y install make && apt-get clean

COPY . .

RUN make install



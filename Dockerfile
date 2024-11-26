FROM golang:1.14-buster

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN apt-get update
RUN apt-get -y install postgresql-client

RUN chmod +x wait-for-postgres.sh

RUN go mod download
RUN build -o todo-app ./cma/main.go

CMD ["./todo-app"]
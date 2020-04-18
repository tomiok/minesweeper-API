FROM golang:1.14.2

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN make build
EXPOSE 8080
CMD ["./ms-api"]
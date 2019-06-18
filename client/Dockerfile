FROM golang:alpine

RUN mkdir /app
ADD . /app/
WORKDIR /app 

RUN CGO_ENABLED=0 go build -a -o client .

CMD ["/app/client"]

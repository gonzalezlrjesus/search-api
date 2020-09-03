FROM golang:1.14.7-alpine3.11 
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 
EXPOSE 8000
RUN go build -o /app/bin/main /app/main.go
CMD ["/app/bin/main"]

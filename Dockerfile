FROM golang:1.14.4

WORKDIR /app

COPY . /app/
RUN go get -u github.com/gin-gonic/gin
EXPOSE 8080

RUN go build -o main .
CMD ["/app/main"]


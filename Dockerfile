FROM golang

RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go get -u github.com/gin-gonic/gin
EXPOSE 8080

RUN go build -o main .
CMD ["/app/main"]


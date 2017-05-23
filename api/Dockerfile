FROM golang:1.8
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go get github.com/gorilla/mux
# RUN go get github.com/lib/pq
RUN go get github.com/go-sql-driver/mysql
RUN go build -o main .
CMD ["/app/main"]
EXPOSE 80

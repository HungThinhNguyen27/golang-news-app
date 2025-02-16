FROM golang:1.23

WORKDIR /app 

COPY . .

# add necessary dependencies
RUN go mod tidy 

RUN go build -o main .

CMD ["/app/main"]

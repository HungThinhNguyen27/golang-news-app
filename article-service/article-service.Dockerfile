FROM golang:1.24.0

WORKDIR /app 

COPY . .

# add necessary dependencies
RUN go mod tidy 

# run file main.go in folder cmd 
RUN go build -o main ./cmd 

CMD ["/app/main", "-yaml", "local.yaml"]

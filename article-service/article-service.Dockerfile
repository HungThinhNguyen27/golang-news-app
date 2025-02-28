FROM golang:1.24.0

WORKDIR /app 

COPY . .

# Cài đặt dependencies
RUN go mod tidy 

RUN go build -o /app/main ./cmd 

CMD ["/app/main", "-yaml", "local.yaml"]

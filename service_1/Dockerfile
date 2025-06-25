FROM golang:1.20

WORKDIR /app

COPY . .

# Add this line to fix the missing go.mod issue
RUN go mod init service1

RUN go build -o main .

CMD ["./main"]

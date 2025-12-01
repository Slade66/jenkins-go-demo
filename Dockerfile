FROM golang:1.20-alpine
WORKDIR /app
COPY . .
RUN go build -o myapp main.go
EXPOSE 8090
CMD ["./myapp"]
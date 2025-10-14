FROM golang:1.24.2
ENV GO111MODULE=on
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod tidy
COPY . .
RUN go build -o /app/server github.com/AliMumtazDev/Go_Chat_App
RUN chmod +x /app/server
EXPOSE 8005
CMD ["/app/server"]
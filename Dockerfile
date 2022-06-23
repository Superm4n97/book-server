FROM golang:latest
WORKDIR /app
ENV UNAME="admin"
ENV UPASS="1234"
ENV SSKEY="super-secret"
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main .
EXPOSE 8080
CMD ["./main"]
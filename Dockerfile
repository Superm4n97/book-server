#===========single staged=============
#FROM golang:latest
#WORKDIR /app
#ENV UNAME="admin"
#ENV UPASS="1234"
#ENV SSKEY="super-secret"
#COPY go.mod go.sum ./
#RUN go mod download
#COPY . .
#RUN go build -o main .
#EXPOSE 8080
#CMD ["./main"]


#===============multistaged=================
FROM golang:latest as builder

#ENV UNAME="admin"
#ENV UPASS="1234"
#ENV SSKEY="superm4n"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

###########start a new stage #################
FROM alpine:latest
RUN apk --no-cache add ca-certificates

ENV UNAME="admin"
ENV UPASS="1234"
ENV SSKEY="superm4n"

WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]
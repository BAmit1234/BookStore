FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download
ADD  . ./
COPY *.go ./
RUN pwd
RUN go build -o main

EXPOSE 8000

CMD [ "./main" ]

FROM golang:1.17-alpine

WORKDIR /app

COPY src .
COPY .env ../
RUN go mod download

RUN go build -o main .

EXPOSE 8080

CMD [ "./main" ]

FROM golang:1.17-alpine
WORKDIR /app
# COPY .env ../
COPY src .
RUN go mod download && go build -o main .
EXPOSE 5000
CMD [ "./main" ]

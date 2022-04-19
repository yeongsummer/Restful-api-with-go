FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

EXPOSE 8080

COPY ./ ./

RUN go build

CMD [ "/restful-api-with-go" ]
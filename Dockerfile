FROM golang:1.21-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /learnacademy-app

EXPOSE 8080

CMD ["/learnacademy-app"]

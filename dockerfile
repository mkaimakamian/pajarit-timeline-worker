# pajarit-timeline-worker/Dockerfile

FROM golang:1.24-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o app

EXPOSE 8081

CMD ["./app"]
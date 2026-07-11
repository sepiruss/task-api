FROM golang:1.25

WORKDIR /app

COPY go.mod go.sum ./

ENV GOPROXY=https://package-mirror.liara.ir/repository/go/
ENV GOSUMDB=off

RUN go mod download

COPY . .

RUN go build -o task-api ./cmd/api

EXPOSE 8080

CMD ["./task-api"]
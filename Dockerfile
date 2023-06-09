FROM golang:1.19

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

EXPOSE 3000

ENTRYPOINT [ "go", "run", "cmd/api/main.go" ]
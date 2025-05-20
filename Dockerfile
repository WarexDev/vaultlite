FROM golang:1.23-alpine

# Installer les dépendances (make, etc.)
RUN apk update && apk add --no-cache make

WORKDIR /app

# Copier les fichiers Go
COPY go.mod ./
RUN go mod download

COPY . .

RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init -g ./cmd/main.go

RUN go mod tidy && go build -o server ./cmd/main.go

EXPOSE 8080

CMD ["./server"]
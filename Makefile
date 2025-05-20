# Nom du binaire à produire
BINARY_NAME=vaultlite
ENTRY_POINT=./cmd/main.go

# Chemins
SRC_DIR=./cmd
LINT=golangci-lint

.PHONY: all build run clean lint docker docker-run test

# Build du binaire
build:
	go build -o $(BINARY_NAME) $(ENTRY_POINT)

# Exécuter localement
run:
	go run $(ENTRY_POINT)

# Linter avec golangci-lint
lint:
	$(LINT) run ./...

# Supprimer le binaire
clean:
	rm -f $(BINARY_NAME)

# Construction de l'image Docker
docker:
	docker build -t $(BINARY_NAME):latest .

# Lancer le conteneur Docker localement
docker-run:
	docker run -p 8080:8080 $(BINARY_NAME):latest

# Tests unitaires
test:
	go test ./...

# Tout faire
all: clean build lint test

# Nom du binaire à produire
BINARY_NAME=vaultlite
ENTRY_POINT=./cmd/main.go

# Outils
LINT=golangci-lint

# Cibles phony
.PHONY: all build run clean lint docker docker-run test doc setup

# Setup des dépendances
setup:
	go install github.com/swaggo/swag/cmd/swag@latest
	go install github.com/swaggo/swag/cmd/swag@latest
	go mod tidy -go="1.23.0"

# Génération de la documentation Swagger
doc:
	swag init -g $(ENTRY_POINT)

# Build du binaire
build: doc
	go build -o $(BINARY_NAME) $(ENTRY_POINT)

# Exécution locale
run:
	go run $(ENTRY_POINT)

# Analyse statique avec golangci-lint
lint:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	$(LINT) run ./...

# Suppression du binaire compilé
clean:
	rm -f $(BINARY_NAME)

# Construction de l'image Docker
docker:
	docker build -t $(BINARY_NAME):latest .

# Lancement du conteneur Docker
docker-run:
	docker run -p 8080:8080 $(BINARY_NAME):latest

# Lancer les tests
test:
	go test ./...

# Exécuter tout le pipeline (clean, build, lint, test)
all: clean setup build lint test
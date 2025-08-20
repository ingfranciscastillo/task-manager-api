# Variables
APP_NAME=task-manager
DOCKER_IMAGE=$(APP_NAME):latest
DB_NAME=taskmanager
DB_USER=postgres
DB_PASSWORD=password

# Comandos de desarrollo
.PHONY: run
run:
	@echo "🚀 Iniciando servidor..."
	go run cmd/server/main.go

.PHONY: build
build:
	@echo "🔨 Compilando aplicación..."
	go build -o bin/$(APP_NAME) cmd/server/main.go

.PHONY: test
test:
	@echo "🧪 Ejecutando tests..."
	go test -v ./...

.PHONY: clean
clean:
	@echo "🧹 Limpiando archivos compilados..."
	rm -rf bin/

# Comandos de base de datos
.PHONY: db-create
db-create:
	@echo "📊 Creando base de datos..."
	createdb -U $(DB_USER) $(DB_NAME)

.PHONY: db-drop
db-drop:
	@echo "🗑️ Eliminando base de datos..."
	dropdb -U $(DB_USER) $(DB_NAME)

.PHONY: db-reset
db-reset: db-drop db-create
	@echo "🔄 Base de datos reiniciada"

# Comandos de Docker
.PHONY: docker-build
docker-build:
	@echo "🐳 Construyendo imagen Docker..."
	docker build -t $(DOCKER_IMAGE) .

.PHONY: docker-run
docker-run:
	@echo "🐳 Ejecutando contenedor..."
	docker run -p 8080:8080 --env-file .env $(DOCKER_IMAGE)

.PHONY: docker-compose-up
docker-compose-up:
	@echo "🐳 Iniciando servicios con Docker Compose..."
	docker-compose up -d

.PHONY: docker-compose-down
docker-compose-down:
	@echo "🐳 Deteniendo servicios..."
	docker-compose down

# Comandos de dependencias
.PHONY: deps
deps:
	@echo "📦 Instalando dependencias..."
	go mod download
	go mod tidy

.PHONY: deps-update
deps-update:
	@echo "📦 Actualizando dependencias..."
	go get -u ./...
	go mod tidy

# Comando de ayuda
.PHONY: help
help:
	@echo "Comandos disponibles:"
	@echo "  run              - Ejecutar servidor de desarrollo"
	@echo "  build            - Compilar aplicación"
	@echo "  test             - Ejecutar tests"
	@echo "  clean            - Limpiar archivos compilados"
	@echo "  db-create        - Crear base de datos"
	@echo "  db-drop          - Eliminar base de datos"
	@echo "  db-reset         - Reiniciar base de datos"
	@echo "  docker-build     - Construir imagen Docker"
	@echo "  docker-run       - Ejecutar contenedor Docker"
	@echo "  deps             - Instalar dependencias"
	@echo "  deps-update      - Actualizar dependencias"
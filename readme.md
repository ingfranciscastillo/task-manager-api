# 🚀 Task Manager API

Una API REST robusta para gestión de tareas construida con Go, Gin, GORM y PostgreSQL.

## ✨ Características

- **Autenticación JWT** - Sistema seguro de autenticación con tokens
- **CRUD de Tareas** - Crear, leer, actualizar y eliminar tareas
- **Concurrencia** - Endpoint de resumen que utiliza goroutines para consultas paralelas
- **Base de datos** - PostgreSQL con migraciones automáticas
- **Arquitectura limpia** - Código organizado en capas (controllers, services, models)
- **Docker ready** - Contenedorización completa con Docker Compose
- **Hashing seguro** - Contraseñas encriptadas con bcrypt

## 🛠️ Tecnologías Utilizadas

- **Go** - Lenguaje de programación
- **Gin** - Framework web HTTP
- **GORM** - ORM para Go
- **PostgreSQL** - Base de datos relacional
- **JWT** - Autenticación con tokens
- **Docker** - Contenedorización
- **bcrypt** - Hashing de contraseñas

## 📋 Endpoints

### Autenticación

- `POST /auth/register` - Registrar usuario
- `POST /auth/login` - Iniciar sesión

### Tareas (requieren JWT)

- `GET /tasks` - Obtener todas las tareas del usuario
- `POST /tasks` - Crear nueva tarea
- `PUT /tasks/:id` - Actualizar tarea
- `DELETE /tasks/:id` - Eliminar tarea
- `GET /tasks/summary` - Obtener resumen de tareas (con concurrencia)

### Utilidades

- `GET /health` - Health check de la API

## 🚀 Instalación y Ejecución

### Prerrequisitos

- Go
- PostgreSQL
- Docker (opcional)
- Make (opcional)

### Opción 1: Ejecución Local

#### Clonar el repositorio

```bash
git clone https://github.com/ingfranciscastillo/task-manager-api
cd task-manager
```

#### Instalar dependencias

```bash
make deps
# o manualmente: go mod download
```

#### Configurar variables de entorno

```bash
cp .env.example .env
# Editar .env con tus configuraciones
```

#### Crear base de datos

```bash
createdb -U postgres taskmanager
# o usar: make db-create
```

#### Ejecutar la aplicación

```bash
make run
# o manualmente: go run cmd/server/main.go
```

### Opción 2: Docker Compose (Recomendado)

#### Clonar y ejecutar

```bash
git clone https://github.com/ingfranciscastillo/task-manager-api
cd task-manager
docker-compose up -d
```

La aplicación estará disponible en `http://localhost:9090`

## 📖 Uso de la API

### Registrar usuario

```bash
curl -X POST http://localhost:9090/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "usuario@ejemplo.com",
    "password": "password123"
  }'
```

### Iniciar sesión

```bash
curl -X POST http://localhost:9090/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "usuario@ejemplo.com",
    "password": "password123"
  }'
```

### Crear tarea (requiere token)

```bash
curl -X POST http://localhost:9090/tasks \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <tu-jwt-token>" \
  -d '{
    "title": "Mi primera tarea",
    "description": "Descripción de la tarea",
    "status": "pending"
  }'
```

### Obtener tareas

```bash
curl -X GET http://localhost:9090/tasks \
  -H "Authorization: Bearer <tu-jwt-token>"
```

### Obtener resumen de tareas

```bash
curl -X GET http://localhost:9090/tasks/summary \
  -H "Authorization: Bearer <tu-jwt-token>"
```

### Actualizar tarea

```bash
curl -X PUT http://localhost:9090/tasks/1 \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <tu-jwt-token>" \
  -d '{
    "title": "Tarea actualizada",
    "status": "completed"
  }'
```

### Eliminar tarea

```bash
curl -X DELETE http://localhost:9090/tasks/1 \
  -H "Authorization: Bearer <tu-jwt-token>"
```

## 🏗️ Arquitectura del Proyecto

```
task-manager/
├── cmd/server/           # Punto de entrada de la aplicación
├── internal/
│   ├── controllers/      # Controladores HTTP (manejo de requests)
│   ├── services/         # Lógica de negocio
│   ├── models/          # Modelos de datos (GORM)
│   ├── middleware/      # Middleware (autenticación, CORS, etc.)
│   ├── routes/          # Definición de rutas
│   └── config/          # Configuración (BD, etc.)
├── pkg/utils/           # Utilidades compartidas
└── migrations/          # Migraciones de BD (futuro)
```

## ⚡ Características de Concurrencia

El endpoint `/tasks/summary` implementa concurrencia usando goroutines y channels:

- Ejecuta 3 consultas en paralelo:
  - Total de tareas
  - Tareas completadas
  - Tareas pendientes
- Utiliza `sync.WaitGroup` para sincronización
- Manejo de errores con channels no bloqueantes

## 🔒 Seguridad

- Contraseñas hasheadas con bcrypt (cost 12)
- JWT con expiración (24 horas)
- Middleware de autenticación para rutas protegidas
- Validación de entrada con Gin binding
- CORS configurado para desarrollo

## 🐳 Docker

El proyecto incluye:

- Dockerfile multi-stage para optimizar tamaño
- docker-compose.yml con PostgreSQL
- Health checks para servicios
- Volumes persistentes para datos

## 📝 Variables de Entorno

| Variable      | Descripción         | Valor por defecto |
| ------------- | ------------------- | ----------------- |
| `DB_HOST`     | Host de PostgreSQL  | `localhost`       |
| `DB_USER`     | Usuario de BD       | `postgres`        |
| `DB_PASSWORD` | Contraseña de BD    | `password`        |
| `DB_NAME`     | Nombre de BD        | `taskmanager`     |
| `DB_PORT`     | Puerto de BD        | `5432`            |
| `JWT_SECRET`  | Secreto para JWT    | `requerido`       |
| `PORT`        | Puerto del servidor | `8080`            |
| `GIN_MODE`    | Modo de Gin         | `debug`           |

## 🧪 Testing

```bash
# Ejecutar tests
make test

# Con coverage
go test -cover ./...

# Test específico
go test ./internal/services -v
```

## 📚 Comandos Make Útiles

```bash
make run           # Ejecutar servidor
make build         # Compilar aplicación
make test          # Ejecutar tests
make docker-build  # Construir imagen Docker
make db-create     # Crear BD
make db-reset      # Reiniciar BD
make deps          # Instalar dependencias
make help          # Ver todos los comandos
```

## 🚨 Troubleshooting

### Error de conexión a BD

- Verificar que PostgreSQL esté ejecutándose
- Revisar variables de entorno en `.env`
- Verificar que la BD `taskmanager` exista

### Error de JWT

- Verificar que `JWT_SECRET` esté configurado
- Verificar formato del token: `Bearer <token>`

### Error de puerto

- Verificar que el puerto 8080 esté libre
- Cambiar `PORT` en `.env` si es necesario

## 🤝 Contribución

1. Fork el proyecto
2. Crear branch de feature (`git checkout -b feature/nueva-caracteristica`)
3. Commit cambios (`git commit -am 'Agregar nueva característica'`)
4. Push al branch (`git push origin feature/nueva-caracteristica`)
5. Crear Pull Request

## 📄 Licencia

Este proyecto está bajo la Licencia MIT - ver el archivo LICENSE para detalles.

## 👨‍💻 Autor

Desarrollado con ❤️ usando Go

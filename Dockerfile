# Build stage
FROM golang:1-alpine3.22 AS builder

WORKDIR /app

# Instalar dependencias del sistema
RUN apk add --no-cache git

# Copiar archivos de dependencias
COPY go.mod go.sum ./
RUN go mod download

# Copiar código fuente
COPY . .

# Compilar aplicación
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main cmd/server/main.go

# Production stage
FROM alpine:latest

# Instalar certificados CA para HTTPS
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copiar binario desde build stage
COPY --from=builder /app/main .
COPY --from=builder /app/.env.example .env

# Exponer puerto
EXPOSE 8080

# Comando para ejecutar
CMD ["./main"]
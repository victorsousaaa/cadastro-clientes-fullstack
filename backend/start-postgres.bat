@echo off
echo === INICIANDO BACKEND COM POSTGRESQL ===
echo.

echo Verificando dependencias...
go mod tidy

echo.
echo Configurando para PostgreSQL...
set DB_TYPE=postgres
set DB_HOST=localhost
set DB_PORT=5432
set DB_USER=postgres
set DB_PASSWORD=postgres
set DB_NAME=cadastro_clientes
set DB_SSLMODE=disable

echo.
echo Iniciando servidor...
echo URL: http://localhost:8080
echo Health Check: http://localhost:8080/health
echo.
echo Para parar: Ctrl+C
echo.

go run main.go

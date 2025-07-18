@echo off
echo === TESTANDO BACKEND ===
echo.

echo 1. Instalando dependencias...
go mod tidy

echo.
echo 2. Usando SQLite (nao precisa PostgreSQL)
set DB_TYPE=sqlite

echo.
echo 3. Iniciando servidor...
echo Servidor rodara em: http://localhost:8080
echo Para parar: Ctrl+C
echo.

go run main.go

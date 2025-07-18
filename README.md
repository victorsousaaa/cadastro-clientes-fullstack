# Cadastro de Clientes - Backend

Backend em Go com PostgreSQL para o sistema de cadastro de clientes.

## 🚀 Como rodar

### Pré-requisitos
- Go 1.21+
- PostgreSQL
- Git

### 1. Instalar dependências
```bash
cd backend
go mod tidy
```

### 2. Configurar banco de dados
Crie um banco PostgreSQL chamado `cadastro_clientes`:
```sql
CREATE DATABASE cadastro_clientes;
```
### 3. Configurar variáveis de ambiente (opcional)
Crie um arquivo `.env` ou configure as variáveis:

```bash
export DB_HOST=localhost
export DB_PORT=5432
export DB_USER=postgres
export DB_PASSWORD=postgres
export DB_NAME=cadastro_clientes
export DB_SSLMODE=disable
```

### 4. Rodar o servidor
```bash
go run main.go
```

O servidor estará rodando em `http://localhost:8080`

## 📋 Endpoints da API

### Clientes
- `GET /api/clientes` - Listar todos os clientes
- `GET /api/clientes/:id` - Buscar cliente por ID
- `POST /api/clientes` - Criar novo cliente
- `PUT /api/clientes/:id` - Atualizar cliente
- `DELETE /api/clientes/:id` - Deletar cliente

### Parâmetros de busca
- `?nome=João` - Filtrar por nome
- `?email=joao@email.com` - Filtrar por email
- `?order_by=nome&order_dir=asc` - Ordenar

### Exemplo de payload (POST/PUT):
```json
{
  "nome": "João Silva",
  "email": "joao@email.com",
  "cpf": "12345678901",
  "nascimento": "1990-01-01",
  "telefone": "11999999999"
}
```

## 🧪 Testar a API

### Health Check
```bash
curl http://localhost:8080/health
```

### Criar cliente
```bash
curl -X POST http://localhost:8080/api/clientes \
  -H "Content-Type: application/json" \
  -d '{
    "nome": "João Silva",
    "email": "joao@email.com",
    "cpf": "12345678901",
    "nascimento": "1990-01-01",
    "telefone": "11999999999"
  }'
```

### Listar clientes
```bash
curl http://localhost:8080/api/clientes
```

## 🏗️ Estrutura do projeto

```
backend/
├── main.go                    # Arquivo principal
├── go.mod                     # Dependências
├── models/
│   └── cliente.go            # Modelo do cliente
├── controllers/
│   └── cliente_controller.go # Controladores
├── database/
│   └── connection.go         # Conexão com banco
└── routes/
    └── routes.go             # Definição das rotas
```

## ✨ Funcionalidades implementadas

- ✅ CRUD completo de clientes
- ✅ Validação de CPF brasileiro
- ✅ Validação de email único
- ✅ Busca por nome e email
- ✅ Ordenação por campos
- ✅ CORS configurado
- ✅ Migrations automáticas
- ✅ Validações de entrada

# 🎯 Sistema de Cadastro de Clientes - Full Stack

## 📋 Pré-requisitos

Antes de executar o projeto, certifique-se de ter instalado:

- **Go 1.21+** - [Download aqui](https://golang.org/dl/)
- **PostgreSQL 12+** - [Download aqui](https://www.postgresql.org/download/)
- **Navegador web moderno** (Chrome, Firefox, Edge)
- **Git** - [Download aqui](https://git-scm.com/)

---

## 🚀 Instalação e Execução

### 1️⃣ Clonar o Repositório

```bash
git clone https://github.com/victorsousaaa/cadastro-clientes-fullstack.git
cd cadastro-clientes-fullstack
```

### 2️⃣ Configurar Banco de Dados PostgreSQL

#### Opção A: Via Interface Gráfica (pgAdmin)
1. Abra o **pgAdmin** ou **DBeaver**
2. Conecte no PostgreSQL
3. Execute este SQL:
```sql
-- Criar banco de dados
CREATE DATABASE cadastro_clientes;

-- Criar usuário (se necessário)
CREATE USER postgres WITH PASSWORD '800425';
GRANT ALL PRIVILEGES ON DATABASE cadastro_clientes TO postgres;
```

#### Opção B: Via Linha de Comando
```bash
# Conectar no PostgreSQL
psql -U postgres

# Executar comandos SQL
CREATE DATABASE cadastro_clientes;
\q
```

### 3️⃣ Configurar e Executar o Backend

```bash
# Entrar na pasta do backend
cd backend

# Baixar dependências do Go
go mod download

# Executar o servidor
go run main.go
```

**✅ Saída esperada:**
```
🔑 Tentando senha: '800425'...
✅ Sucesso com senha: '800425'
✅ PostgreSQL conectado e tabelas prontas!
Servidor iniciando na porta 8080...
[GIN-debug] Listening and serving HTTP on :8080
```

> **📝 Nota:** O sistema criará as tabelas automaticamente na primeira execução!

### 4️⃣ Executar o Frontend

```bash
# Voltar para a pasta raiz
cd ..

# Abrir o formulário de cadastro
# Método 1: Duplo clique no arquivo
Index.html

# Método 2: Pelo navegador
# Abrir navegador e acessar:
file:///C:/caminho/para/seu/projeto/Index.html
```

**📱 URLs de Acesso:**
- **Cadastro:** `Index.html`
- **Lista de Clientes:** `cliente.html`  
- **API:** `http://localhost:8080/api/clientes`

---

## 🗄️ Estrutura das Tabelas

O sistema criará automaticamente a tabela no banco:

```sql
-- Tabela criada automaticamente pelo GORM
CREATE TABLE clientes (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    cpf VARCHAR(11) UNIQUE NOT NULL,
    nascimento VARCHAR NOT NULL,
    telefone VARCHAR(15)
);

-- Índices únicos para email e CPF
CREATE UNIQUE INDEX idx_clientes_email ON clientes(email);
CREATE UNIQUE INDEX idx_clientes_cpf ON clientes(cpf);
```

> **🔧 GORM Auto-Migration:** As tabelas são criadas automaticamente quando o backend inicia!

---

## 🔧 Configuração de Ambiente

### Senhas Testadas Automaticamente
O sistema testa estas senhas do PostgreSQL automaticamente:
1. `800425`
2. `Roch@8004` 
3. `postgres`

### Configuração Manual (Opcional)
Se usar senha diferente, edite `backend/database/connection.go`:

```go
// Linha ~20: Adicione sua senha na lista
passwords := []string{"800425", "Roch@8004", "postgres", "SUA_SENHA_AQUI"}
```

---

## 📖 Como Usar o Sistema

### 1️⃣ Cadastrar Cliente
1. Abra `Index.html` no navegador
2. Preencha os campos (nome, email, CPF, data nascimento)
3. Veja as **máscaras automáticas** funcionando:
   - CPF: `123.456.789-01`
   - Telefone: `(11) 99999-9999`
4. Clique **"Cadastrar"**

### 2️⃣ Visualizar Clientes
1. Após cadastrar → redirecionamento automático
2. Ou abra `cliente.html` diretamente
3. **Funcionalidades disponíveis:**
   - 🔍 **Buscar** por nome ou email
   - 📊 **Ordenar** clicando nos headers
   - ✏️ **Editar** clicando no cliente
   - 🗑️ **Remover** com confirmação

### 3️⃣ Testar API
```bash
# Verificar status
curl http://localhost:8080/health

# Listar clientes
curl http://localhost:8080/api/clientes

# Criar cliente via API
curl -X POST http://localhost:8080/api/clientes \
  -H "Content-Type: application/json" \
  -d '{
    "nome": "João Silva",
    "email": "joao@email.com", 
    "cpf": "12345678901",
    "nascimento": "1990-01-15",
    "telefone": "11999999999"
  }'
```

---

## 🐛 Solução de Problemas

### ❌ Erro: "connection refused"
**Problema:** PostgreSQL não está rodando
```bash
# Windows
net start postgresql-x64-12

# Linux/Mac  
sudo service postgresql start
```

### ❌ Erro: "senha incorreta"
**Problema:** Senha do PostgreSQL diferente
1. Edite `backend/database/connection.go`
2. Adicione sua senha na lista `passwords`
3. Reinicie o backend

### ❌ Erro: "go: command not found"
**Problema:** Go não instalado
1. Baixe: https://golang.org/dl/
2. Instale e reinicie o terminal
3. Teste: `go version`

### ❌ Frontend não carrega
**Problema:** Arquivo não está sendo servido corretamente
1. **Método 1:** Duplo clique no `Index.html`
2. **Método 2:** Use Live Server do VS Code
3. **Método 3:** Python: `python -m http.server 3000`

### ❌ Máscaras não funcionam
**Problema:** JavaScript com erro
1. Abra F12 (Console do navegador)
2. Veja se há erros em vermelho
3. Certifique-se que `script.js` está carregando

---

## 🔗 Endpoints da API

| Método | URL | Descrição | Body |
|--------|-----|-----------|------|
| `GET` | `/health` | Status da API | - |
| `GET` | `/api/clientes` | Listar todos | - |
| `POST` | `/api/clientes` | Criar novo | JSON cliente |
| `PUT` | `/api/clientes/:id` | Atualizar | JSON cliente |
| `DELETE` | `/api/clientes/:id` | Remover | - |

### Exemplo de Cliente JSON:
```json
{
  "nome": "Victor Hugo",
  "email": "victor@exemplo.com",
  "cpf": "12345678901",
  "nascimento": "1995-10-15", 
  "telefone": "11999999999"
}
```

---

## ✨ Funcionalidades Implementadas

### 🔢 Máscaras Automáticas
- **CPF:** Digite `12345678901` → vira `123.456.789-01`
- **Telefone:** Digite `11999999999` → vira `(11) 99999-9999`

### 🛡️ Validações
- **CPF:** Algoritmo brasileiro oficial
- **Email:** Formato válido + único no banco  
- **Campos obrigatórios:** Nome, email, CPF, nascimento

### 🔍 Interface Interativa
- **Busca em tempo real** por nome/email
- **Ordenação** clicando nos headers da tabela
- **Edição inline** com máscaras mantidas
- **Remoção** com confirmação

### 💾 Resilência
- **Fallback automático** para localStorage se API indisponível
- **Validação dupla** frontend + backend
- **Tratamento de erros** amigável

---

## 🧪 Dados para Teste

### CPFs Válidos
```
123.456.789-09
987.654.321-00
111.444.777-35
```

### CPFs Inválidos (para testar validação)
```
111.111.111-11  // Todos iguais
123.456.789-00  // Dígito errado
123.456.789     // Incompleto
```

### Clientes de Exemplo
```json
{
  "nome": "Ana Silva",
  "email": "ana@teste.com",
  "cpf": "12345678901", 
  "nascimento": "1990-05-15",
  "telefone": "11987654321"
}
```

---

## 📁 Estrutura do Projeto

```
cadastro-clientes-fullstack/
├── 📁 backend/                 # Servidor Go
│   ├── main.go                 # Entrada principal
│   ├── models/cliente.go       # Modelo dados
│   ├── controllers/            # Lógica negócio
│   ├── routes/                 # Rotas API
│   ├── database/               # Conexão BD
│   └── go.mod                  # Dependências
├── 📁 frontend/                # Assets estáticos
│   ├── style.css              # Estilos
│   ├── img/tecmise.jpg        # Imagem fundo
│   └── fonts/AnAwesome.ttf    # Fonte
├── 📄 Index.html              # Formulário cadastro
├── 📄 cliente.html            # Lista clientes  
├── 📄 script.js               # Lógica cadastro
├── 📄 cliente.js              # Lógica listagem
└── 📄 README.md               # Esta documentação
```

---

## 🔧 Tecnologias e Dependências

### Backend (Go)
```go
// go.mod
module cadastro-clientes

go 1.21

require (
    github.com/gin-gonic/gin v1.9.1
    gorm.io/gorm v1.25.1
    gorm.io/driver/postgres v1.5.2
)
```

### Frontend
- **Vue.js 3** (via CDN)
- **HTML5/CSS3** nativo
- **JavaScript ES6+** 
- **Fetch API** para HTTP

---

## 👨‍💻 Desenvolvedor

**Victor Hugo**  
📧 Email: victorsousaa07@gmail.com  
🔗 GitHub: [@victorsousaaa](https://github.com/victorsousaaa)

---


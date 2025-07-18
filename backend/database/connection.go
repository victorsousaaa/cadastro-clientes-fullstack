package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	_ "github.com/lib/pq"
	"cadastro-clientes-api/models"
)

var DB *gorm.DB

func Connect() {
	
	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "5432")
	user := getEnv("DB_USER", "postgres")
	password := getEnv("DB_PASSWORD", "postgres")
	dbname := getEnv("DB_NAME", "cadastro_clientes")
	sslmode := getEnv("DB_SSLMODE", "disable")

	passwords := []string{
		"800425",     // SUA SENHA QUE FUNCIONA!
		password,     // senha do ambiente
	}

	for _, pwd := range passwords {
		log.Printf("🔑 Tentando senha: '%s'...", pwd)
		
		if connectPostgreSQL(host, port, user, pwd, dbname, sslmode) {
			log.Printf("✅ Sucesso com senha: '%s'", pwd)
			return
		}
	}

	log.Println("❌ PostgreSQL falhou. Configure a senha correta!")
	log.Println("💡 Exemplo: $env:DB_PASSWORD=\"SUA_SENHA\"; go run .")
	log.Fatal("Falha na conexão com PostgreSQL")
}
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func connectPostgreSQL(host, port, user, password, dbname, sslmode string) bool {
	dsnBase := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=postgres sslmode=%s",
		host, port, user, password, sslmode)
	
	dbBase, err := sql.Open("postgres", dsnBase)
	if err != nil {
		return false
	}
	defer dbBase.Close()

	if err := dbBase.Ping(); err != nil {
		return false
	}

	_, err = dbBase.Exec(fmt.Sprintf("CREATE DATABASE %s", dbname))
	if err != nil {
		log.Printf("⚠️  Banco '%s' já existe ou erro: %v", dbname, err)
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode)
	
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return false
	}

	DB = database

	log.Println("🔄 Criando tabelas...")
	err = DB.AutoMigrate(&models.Cliente{})
	if err != nil {
		log.Printf("❌ Erro GORM migration: %v", err)
		log.Println("🔧 Tentando criar tabela manualmente...")
		createTableSQL := `
		CREATE TABLE IF NOT EXISTS clientes (
			id SERIAL PRIMARY KEY,
			nome VARCHAR(100) NOT NULL,
			email VARCHAR(100) UNIQUE NOT NULL,
			cpf VARCHAR(11) UNIQUE NOT NULL,
			nascimento VARCHAR(10) NOT NULL,
			telefone VARCHAR(20),
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)`
		
		sqlDB, _ := DB.DB()
		_, errCreate := sqlDB.Exec(createTableSQL)
		if errCreate != nil {
			log.Printf("❌ Erro criar tabela manual: %v", errCreate)
			return false
		}
		log.Println("✅ Tabela criada manualmente!")
	} else {
		log.Println("✅ Tabelas criadas via GORM!")
	}

	log.Println("✅ PostgreSQL conectado e tabelas prontas!")
	return true
}


package controllers

import (
	"cadastro-clientes-api/database"
	"cadastro-clientes-api/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func GetClientes(c *gin.Context) {
	var clientes []models.Cliente

	nome := c.Query("nome")
	email := c.Query("email")
	orderBy := c.Query("order_by")   
	orderDir := c.Query("order_dir") 

	query := database.DB

	
	if nome != "" {
		query = query.Where("nome ILIKE ?", "%"+nome+"%")
	}
	if email != "" {
		query = query.Where("email ILIKE ?", "%"+email+"%")
	}
	if orderBy == "nome" || orderBy == "email" {
		order := orderBy
		if orderDir == "desc" {
			order += " DESC"
		} else {
			order += " ASC"
		}
		query = query.Order(order)
	} else {
		query = query.Order("created_at DESC")
	}

	if err := query.Find(&clientes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar clientes"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": clientes})
}

func GetClienteByID(c *gin.Context) {
	var cliente models.Cliente
	id := c.Param("id")

	if err := database.DB.First(&cliente, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cliente não encontrado"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": cliente})
}


func CreateCliente(c *gin.Context) {
	var cliente models.Cliente

	if err := c.ShouldBindJSON(&cliente); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}


	if err := validate.Struct(&cliente); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos", "details": err.Error()})
		return
	}

	cliente.CPF = strings.ReplaceAll(cliente.CPF, ".", "")
	cliente.CPF = strings.ReplaceAll(cliente.CPF, "-", "")

	if !cliente.ValidarCPF() {
		c.JSON(http.StatusBadRequest, gin.H{"error": "CPF inválido"})
		return
	}

	var existingCliente models.Cliente
	if err := database.DB.Where("cpf = ? OR email = ?", cliente.CPF, cliente.Email).First(&existingCliente).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "CPF ou email já cadastrados"})
		return
	}

	if err := database.DB.Create(&cliente).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar cliente"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": cliente, "message": "Cliente criado com sucesso"})
}

func UpdateCliente(c *gin.Context) {
	var cliente models.Cliente
	id := c.Param("id")

	if err := database.DB.First(&cliente, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cliente não encontrado"})
		return
	}

	var updateData models.Cliente
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validate.Struct(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos", "details": err.Error()})
		return
	}

	updateData.CPF = strings.ReplaceAll(updateData.CPF, ".", "")
	updateData.CPF = strings.ReplaceAll(updateData.CPF, "-", "")

	if !updateData.ValidarCPF() {
		c.JSON(http.StatusBadRequest, gin.H{"error": "CPF inválido"})
		return
	}

	var existingCliente models.Cliente
	if err := database.DB.Where("(cpf = ? OR email = ?) AND id != ?", updateData.CPF, updateData.Email, id).First(&existingCliente).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "CPF ou email já cadastrados"})
		return
	}

	cliente.Nome = updateData.Nome
	cliente.Email = updateData.Email
	cliente.CPF = updateData.CPF
	cliente.Nascimento = updateData.Nascimento
	cliente.Telefone = updateData.Telefone

	if err := database.DB.Save(&cliente).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar cliente"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": cliente, "message": "Cliente atualizado com sucesso"})
}

func DeleteCliente(c *gin.Context) {
	var cliente models.Cliente
	id := c.Param("id")

	if err := database.DB.First(&cliente, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cliente não encontrado"})
		return
	}

	if err := database.DB.Delete(&cliente).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar cliente"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Cliente deletado com sucesso"})
}

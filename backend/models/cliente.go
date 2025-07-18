package models

import (
	"time"
)

type Cliente struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	Nome       string    `json:"nome" gorm:"not null" validate:"required,min=2,max=100"`
	Email      string    `json:"email" gorm:"unique;not null" validate:"required,email"`
	CPF        string    `json:"cpf" gorm:"unique;not null" validate:"required,len=11"`
	Nascimento string    `json:"nascimento" gorm:"not null" validate:"required"`
	Telefone   string    `json:"telefone,omitempty"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (c *Cliente) ValidarCPF() bool {
	cpf := c.CPF
	if len(cpf) != 11 {
		return false
	}
	igual := true
	for i := 1; i < 11; i++ {
		if cpf[i] != cpf[0] {
			igual = false
			break
		}
	}
	if igual {
		return false
	}
	soma := 0
	for i := 0; i < 9; i++ {
		soma += int(cpf[i]-'0') * (10 - i)
	}

	resto := soma % 11
	if resto < 2 {
		resto = 0
	} else {
		resto = 11 - resto
	}

	if int(cpf[9]-'0') != resto {
		return false
	}

	soma = 0
	for i := 0; i < 10; i++ {
		soma += int(cpf[i]-'0') * (11 - i)
	}

	resto = soma % 11
	if resto < 2 {
		resto = 0
	} else {
		resto = 11 - resto
	}

	return int(cpf[10]-'0') == resto
}

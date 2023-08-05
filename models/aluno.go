package models

import "gorm.io/gorm"

// tys - for shortcut
type Aluno struct {
	gorm.Model        // with property gorm.Model automatically will be included: id, createdAt, updatedAt, deletedAt.
	Nome       string `json:"nome"`
	CPF        string `json:"cpf"`
	RG         string `json:"rg"`
}

var Alunos []Aluno

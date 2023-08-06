package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

// tys - for shortcut

// serializing data
type Aluno struct {
	gorm.Model        // with property gorm.Model automatically will be included: id, createdAt, updatedAt, deletedAt.
	Nome       string `json:"nome" validate:"nonzero"`
	CPF        string `json:"cpf" validate:"len=9, regexp=^[0-9]*$"`
	RG         string `json:"rg" validate:"len=12, regexp=^[0-9]*$"`
}

func ValidateStudents(aluno *Aluno) error { // this function will be called in /controllers.
	if err := validator.Validate(aluno); err != nil {
		return err

	}
	return nil

}

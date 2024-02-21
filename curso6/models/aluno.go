package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Aluno struct {
	gorm.Model
	Id   int    `json:"id"`
	Nome string `json:"nome" validate:"nonzero"`
	Cpf  string `json:"cpf" validate:"len=11,regexp=^[0-9]*$"`
	Rg   string `json:"rg" validate:"len=9,regexp=^[0-9]*$"`
}

var Alunos []Aluno

func ValidarDadosAluno(aluno *Aluno) error {
	if err := validator.Validate(aluno); err != nil {
		return err
	}

	return nil
}

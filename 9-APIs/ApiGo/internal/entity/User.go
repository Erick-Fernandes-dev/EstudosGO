package entity

import (
	"github.com/APIs/pkg/entity"
	"golang.org/x/crypto/bcrypt"
)

// VO - Value Object
type User struct {
	ID       entity.ID `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"-"`
}

func NewUser(name, email, password string) (*User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	return &User{
		ID:    entity.NewId(),
		Name:  name,
		Email: email,
		//Convetendo o hash para string
		Password: string(hash),
	}, nil
}

// Validando a senha
func (u *User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

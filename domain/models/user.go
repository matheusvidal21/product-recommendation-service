package models

import "golang.org/x/crypto/bcrypt"

type UserDomain interface {
	GetID() string
	GetName() string
	GetEmail() string
	GetPassword() string
	ValidatePassword(string) bool
}

type user struct {
	id       string
	name     string
	email    string
	password string
}

func NewUser(id, name, email, password string) UserDomain {
	encryptedPass, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return &user{
		id:       id,
		name:     name,
		email:    email,
		password: string(encryptedPass),
	}
}

func (u *user) GetID() string {
	return u.id
}

func (u *user) GetName() string {
	return u.name
}

func (u *user) GetEmail() string {
	return u.email
}

func (u *user) GetPassword() string {
	return u.password
}

func (u *user) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.password), []byte(password))
	return err == nil
}

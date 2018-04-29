package controllers

import (
 "Phonebook/models"
)

type (
	UserResource struct {
		Data models.User `json:"data"`
	}

	LoginResource struct {
		Data LoginModel `json:"data"`
	}

	AuthUserResource struct {
		Data AuthUserModel `json:"data"`
	}

	LoginModel struct {
		Email string `json:"email"`
		Password string `json:"password"`
	}
	AuthUserModel struct {
		User models.User `json:"user"`
		Token string     `json:"token"`
	}
	NumberResource struct {
		Data models.Number `json:"data"`
	}
	NumbersResource struct {
		Data []models.Number `json:"data"`
	}

)

package controllers

import (
	"fmt"
	"net/http"

	"github.com/sixsat/lenslocked/models"
)

type User struct {
	Templates struct {
		New Template
	}
	UserService *models.UserService
}

func (u User) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Email: ", r.FormValue("email"))
	fmt.Fprintln(w, "Password: ", r.FormValue("password"))
}

func (u User) New(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}
	data.Email = r.FormValue("email")
	u.Templates.New.Execute(w, data)
}

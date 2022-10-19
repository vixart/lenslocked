package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sixsat/lenslocked/controllers"
	"github.com/sixsat/lenslocked/models"
	"github.com/sixsat/lenslocked/templates"
	"github.com/sixsat/lenslocked/views"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", controllers.StaticHandler(views.Must(views.ParseFS(
		templates.FS,
		"tailwind.html", "home.html",
	))))

	r.Get("/contact", controllers.StaticHandler(views.Must(views.ParseFS(
		templates.FS,
		"tailwind.html", "contact.html",
	))))

	r.Get("/faq", controllers.FAQ(views.Must(views.ParseFS(
		templates.FS,
		"tailwind.html", "faq.html",
	))))

	// Setup a database connection
	cfg := models.DefaultPostgresConfig()
	db, err := models.Open(cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Setup model service
	userService := models.UserService{
		DB: db,
	}

	// Setup user controller
	userCtrl := controllers.User{
		UserService: &userService,
	}
	userCtrl.Templates.New = views.Must(views.ParseFS(
		templates.FS,
		"tailwind.html", "signup.html",
	))
	userCtrl.Templates.SignIn = views.Must(views.ParseFS(
		templates.FS,
		"tailwind.html", "signin.html",
	))
	r.Get("/signup", userCtrl.New)
	r.Post("/user", userCtrl.Create)
	r.Get("/signin", userCtrl.SignIn)
	r.Post("/signin", userCtrl.ProcessSignIn)
	r.Get("/user/me", userCtrl.CurrentUser)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}

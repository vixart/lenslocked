package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/csrf"
	"github.com/sixsat/lenslocked/controllers"
	"github.com/sixsat/lenslocked/models"
	"github.com/sixsat/lenslocked/templates"
	"github.com/sixsat/lenslocked/views"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", controllers.StaticHandler(views.Must(views.ParseFS(
		templates.FS,
		"tailwind.gohtml", "home.gohtml",
	))))

	r.Get("/contact", controllers.StaticHandler(views.Must(views.ParseFS(
		templates.FS,
		"tailwind.gohtml", "contact.gohtml",
	))))

	r.Get("/faq", controllers.FAQ(views.Must(views.ParseFS(
		templates.FS,
		"tailwind.gohtml", "faq.gohtml",
	))))

	// Setup a database connection
	cfg := models.DefaultPostgresConfig()
	db, err := models.Open(cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Setup model services
	userService := models.UserService{
		DB: db,
	}
	sessionService := models.SessionService{
		DB: db,
	}

	// Setup user controller
	userCtrl := controllers.User{
		UserService:    &userService,
		SessionService: &sessionService,
	}
	userCtrl.Templates.New = views.Must(views.ParseFS(
		templates.FS,
		"tailwind.gohtml", "signup.gohtml",
	))
	userCtrl.Templates.SignIn = views.Must(views.ParseFS(
		templates.FS,
		"tailwind.gohtml", "signin.gohtml",
	))
	r.Get("/signup", userCtrl.New)
	r.Post("/user", userCtrl.Create)
	r.Get("/signin", userCtrl.SignIn)
	r.Post("/signin", userCtrl.ProcessSignIn)
	r.Post("/signout", userCtrl.ProcessSignOut)
	r.Get("/user/me", userCtrl.CurrentUser)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found.", http.StatusNotFound)
	})

	// Setup CSRF middleware
	csrfKey := "gFvi45R4fy5xNBlnEeZtQbfAVCYEIAUX"
	csrfMw := csrf.Protect(
		[]byte(csrfKey),
		// TODO: Remove this before deploying
		csrf.Secure(false),
	)

	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", csrfMw(r))
}

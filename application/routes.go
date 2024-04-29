package application

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/Infael/gogoVseProject/controller"
)

func (app *App) loadRoutes() {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	router.Route("/newsletters", app.loadNewslettersRoutes)
	router.Route("/users", app.loadUserRoutes)

	app.router = router
}

func (app *App) loadNewslettersRoutes(router chi.Router) {
	newsletterController := &controller.Newsletter{}

	router.Post("/", newsletterController.Create)
	router.Get("/", newsletterController.List)
	router.Get("/{id}", newsletterController.GetById)
	router.Put("/{id}", newsletterController.UpdateById)
	router.Delete("/{id}", newsletterController.DeleteById)

	router.Post("/{id}/posts", newsletterController.CreatePost)
	router.Get("/{id}/posts", newsletterController.GetPosts)
}

func (app *App) loadUserRoutes(router chi.Router) {
	userController := &controller.User{}

	router.Post("/register", userController.Register)
	router.Post("/login", userController.Login)
	router.Delete("/", userController.DeleteAccount)
	router.Post("/reset-password", userController.ResetPassword)
	router.Get("/newsletters", userController.GetNewsletters)
}

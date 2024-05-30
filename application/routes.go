package application

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/Infael/gogoVseProject/controller"
	"github.com/Infael/gogoVseProject/middlewares"
)

func (app *App) loadRoutes() {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	router.Route("/users", app.loadUserRoutes)
	router.Route("/newsletters", app.loadNewslettersRoutes)
	router.Route("/auth", app.loadAuthRoutes)
	router.Route("/password", app.loadPasswordRoutes)

	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "static"))
	loadFileServer(router, "/static", filesDir)

	app.router = router
}

func (app *App) loadNewslettersRoutes(router chi.Router) {
	newsletterController := controller.NewNewsletterController(&app.services.NewsletterService, &app.services.PostService, &app.services.UserService, &app.services.SubscriberService)

	// Protected endpoints
	router.Group(
		func(r chi.Router) {
			r.Use(middlewares.JwtAuthMiddleware)
			r.Post("/", newsletterController.Create)
			r.Get("/", newsletterController.List)
			r.Get("/{id}", newsletterController.GetById)
			r.Put("/{id}", newsletterController.UpdateById)
			r.Delete("/{id}", newsletterController.DeleteById)

			r.Post("/{id}/posts", newsletterController.CreatePost)
			r.Get("/{id}/posts", newsletterController.GetPosts)

			r.Get("/{id}/subscribers", newsletterController.GetSubscribers)
		},
	)

	// Public endpoints
	router.Post("/{id}/subscribers", newsletterController.Subscribe)
	router.Get("/{id}/subscribers/unsubscribe/{subId}", newsletterController.Unsubscribe)
	router.Delete("/{id}/subscribers/unsubscribe/{subId}", newsletterController.Unsubscribe)
	router.Get("/{id}/subscribers/verify/{token}", newsletterController.VerifySubscriber)
	router.Post("/{id}/subscribers/verify/{token}", newsletterController.VerifySubscriber)
}

func (app *App) loadUserRoutes(router chi.Router) {
	userController := controller.NewUserController(&app.services.UserService)

	// Protected endpoints
	router.Group(
		func(r chi.Router) {
			r.Use(middlewares.JwtAuthMiddleware)

			r.Put("/", userController.UpdateAccount)
			// I wanted to try asynchronous operations in Go, so I added a 60-second delay to the deletion of the user.
			r.Delete("/", userController.DeleteAccount)
			r.Post("/cancel-delete", userController.CancelUserDeletion)
		},
	)
}

func (app *App) loadAuthRoutes(router chi.Router) {
	authController := controller.NewAuthController(&app.services.AuthService)

	router.Post("/register", authController.Register)
	router.Post("/login", authController.Login)
}

func (app *App) loadPasswordRoutes(router chi.Router) {
	passwordController := controller.NewPasswordController(&app.services.PasswordService)

	router.Post("/request-reset", passwordController.SendResetPasswordEmail)
	router.Post("/reset/{token}", passwordController.SetNewPasswordWithResetToken)
}

// FileServer conveniently sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
func loadFileServer(router chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	// Ensure the path ends with a slash
	if path != "/" && path[len(path)-1] != '/' {
		router.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	// Add wildcard for file matching
	path += "*"
	log.Printf("Setting up file server at path: %s", path)

	router.Get("/static", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	router.Get(path, func(w http.ResponseWriter, r *http.Request) {
		log.Printf("File server route hit for URL: %s", r.URL.Path)
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		log.Printf("Path prefix: %s", pathPrefix)
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}

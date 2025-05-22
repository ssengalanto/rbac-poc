package main

import (
	"github.com/casbin/casbin/v2"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"rbac-poc/middleware"
)

func main() {
	r := chi.NewRouter()

	enforcer, err := casbin.NewEnforcer("model.conf", "policy.csv")
	if err != nil {
		log.Fatalf("Failed to create Casbin enforcer: %v", err)
	}

	// Public routes
	r.Group(func(public chi.Router) {
		public.Post("/login", writeText("Login"))
		public.Get("/healthz", writeText("OK"))
	})

	// Protected routes
	r.Group(func(protected chi.Router) {
		protected.Use(middleware.AuthMiddleware)
		protected.Use(middleware.RBACMiddleware(enforcer))

		// Workspaces routes
		protected.Route("/workspaces", func(r chi.Router) {
			r.Get("/", writeText("Read Workspaces"))
			r.Post("/", writeText("Create Workspaces"))
			r.Put("/", writeText("Update Workspaces"))
			r.Delete("/", writeText("Delete Workspaces"))
			r.Post("/duplicate", writeText("Duplicate Workspace"))
			r.Put("/toggle", writeText("Toggle Workspace"))
		})

		// Users routes
		protected.Route("/users", func(r chi.Router) {
			r.Get("/", writeText("Read Users"))
			r.Post("/", writeText("Create Users"))
			r.Put("/", writeText("Update Users"))
			r.Delete("/", writeText("Delete Users"))
			r.Get("/view-assigned", writeText("View Assigned Users"))
		})

		// Guests routes
		protected.Route("/guests", func(r chi.Router) {
			r.Get("/", writeText("Read Guests"))
			r.Post("/", writeText("Create Guests"))
			r.Put("/", writeText("Update Guests"))
			r.Delete("/", writeText("Delete Guests"))
			r.Put("/bulk-upload", writeText("Bulk Upload Guests"))
		})
	})

	http.ListenAndServe(":8080", r)
}

// writeText returns a simple handler that writes the given text
func writeText(text string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(text))
	}
}

package middleware

import (
	"github.com/casbin/casbin/v2"
	"log"
	"net/http"
)

var httpMethodToAction = map[string]string{
	"GET":     "read",
	"POST":    "create",
	"PUT":     "update",
	"PATCH":   "update",
	"DELETE":  "delete",
	"OPTIONS": "read",
	"HEAD":    "read",
}

func RBACMiddleware(e *casbin.Enforcer) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			role := r.Context().Value("role")
			if role == nil {
				log.Println("Missing role in context")
				http.Error(w, "Forbidden: no role found", http.StatusForbidden)
				return
			}

			sub := role.(string)
			obj := r.URL.Path
			act := httpMethodToAction[r.Method]
			act, exists := httpMethodToAction[r.Method]
			if !exists {
				log.Printf("Unsupported HTTP method: %s", r.Method)
				http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
				return
			}

			ok, err := e.Enforce(sub, obj, act)
			if err != nil {
				log.Printf("Casbin enforcement error: %v (sub=%s, obj=%s, act=%s)", err, sub, obj, act)
				http.Error(w, "Internal Server Error during enforcement", http.StatusInternalServerError)
				return
			}

			log.Printf("RBAC check: role=%s path=%s method=%s action=%s allowed=%v", sub, obj, r.Method, act, ok)

			if !ok {
				http.Error(w, "Forbidden: access denied", http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

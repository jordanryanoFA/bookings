package main

import (
	"net/http"

	"github.com/jordanryanoFA/bookings/internal/helpers"
	"github.com/justinas/nosurf"
)

//	func WriteToConsole(next http.Handler) http.Handler {
//		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//			fmt.Println("hit the page")
//			next.ServeHTTP(w, r)
//		})
//	}
//

// Nosurf adds CSRF protection to all post request
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

// SessionLoad and saves the session on every request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !helpers.IsAutenticated(r) {
			session.Put(r.Context(), "error", "log in first!")
			http.Redirect(w, r, "/user/login", http.StatusSeeOther)
		}
		next.ServeHTTP(w, r)
	})
}

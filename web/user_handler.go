package web

import (
	"net/http"
	"time"
)

// 登録ページ
func (s *server) willSignupHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.renderTemplate(w, r, "signup.tmpl", nil)
	})
}

// 登録処理
func (s *server) signupHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name, password := r.FormValue("name"), r.FormValue("password")
		if err := s.app.CreateNewUser(name, password); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		user, err := s.app.FindUserByName(name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		expiresAt := time.Now().Add(24 * time.Hour)
		token, err := s.app.CreateNewToken(user.ID, expiresAt)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.SetCookie(w, &http.Cookie{
			Name:    sessionKey,
			Value:   token,
			Expires: expiresAt,
		})
		http.Redirect(w, r, "/", http.StatusSeeOther)
	})
}

func (s *server) willSigninHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.renderTemplate(w, r, "signin.tmpl", nil)
	})
}

func (s *server) signinHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name, password := r.FormValue("name"), r.FormValue("password")
		if ok, err := s.app.LoginUser(name, password); err != nil || !ok {
			http.Error(w, "user not found or invalid password", http.StatusBadRequest)
			return
		}
		user, err := s.app.FindUserByName(name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		expiresAt := time.Now().Add(24 * time.Hour)
		token, err := s.app.CreateNewToken(user.ID, expiresAt)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.SetCookie(w, &http.Cookie{
			Name:    sessionKey,
			Value:   token,
			Expires: expiresAt,
		})
		http.Redirect(w, r, "/", http.StatusSeeOther)
	})
}

func (s *server) signoutHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.SetCookie(w, &http.Cookie{
			Name:    sessionKey,
			Value:   "",
			Expires: time.Unix(0, 0),
		})
		http.Redirect(w, r, "/", http.StatusSeeOther)
	})
}

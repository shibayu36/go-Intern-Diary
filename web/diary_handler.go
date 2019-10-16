package web

import (
	"net/http"
)

// ダイアリー一覧
func (s *server) diariesHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := s.findUser(r)
		if user == nil {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		diaries, err := s.app.ListDiariesByUserID(user.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		s.renderTemplate(w, r, "diaries.tmpl", map[string]interface{}{
			"User":    user,
			"Diaries": diaries,
		})
	})
}

// ダイアリー作成ページ
func (s *server) willDiaryCreateHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := s.findUser(r)
		if user == nil {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		s.renderTemplate(w, r, "diary_create.tmpl", map[string]interface{}{
			"User": user,
		})
	})
}

// ダイアリー作成
func (s *server) diaryCreateHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := s.findUser(r)
		if user == nil {
			http.Error(w, "please login", http.StatusBadRequest)
			return
		}

		name := r.FormValue("name")
		_, err := s.app.CreateNewDiary(user.ID, name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/diaries", http.StatusSeeOther)
	})
}

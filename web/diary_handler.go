package web

import (
	"net/http"
)

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
		if err := s.app.CreateNewDiary(user.ID, name); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	})
}

package web

import "net/http"

// 登録ページ
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

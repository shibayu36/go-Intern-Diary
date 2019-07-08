package web

//go:generate go-assets-builder --package=web --output=./templates-gen.go --strip-prefix="/templates/" --variable=Templates ../templates

import (
	"html/template"
	"io/ioutil"
	"net/http"

	"github.com/dimfeld/httptreemux"
	"github.com/justinas/nosurf"

	"github.com/hatena/go-Intern-Diary/model"
	"github.com/hatena/go-Intern-Diary/service"
)

type Server interface {
	Handler() http.Handler
}

const sessionKey = "DIARY_SESSION"

var templates map[string]*template.Template

func init() {
	var err error
	templates, err = loadTemplates()
	if err != nil {
		panic(err)
	}
}

func loadTemplates() (map[string]*template.Template, error) {
	templates := make(map[string]*template.Template)
	bs, err := ioutil.ReadAll(Templates.Files["main.tmpl"])
	if err != nil {
		return nil, err
	}
	mainTmpl := template.Must(template.New("main.tmpl").Parse(string(bs)))
	for fileName, file := range Templates.Files {
		bs, err := ioutil.ReadAll(file)
		if err != nil {
			return nil, err
		}
		mainTmpl := template.Must(mainTmpl.Clone())
		templates[fileName] = template.Must(mainTmpl.New(fileName).Parse(string(bs)))
	}
	return templates, nil
}

func NewServer(app service.DiaryApp) Server {
	return &server{app: app}
}

type server struct {
	app service.DiaryApp
}

func (s *server) Handler() http.Handler {
	router := httptreemux.New()

	handle := func(method, path string, handler http.Handler) {
		router.UsingContext().Handler(method, path,
			csrfMiddleware(loggingMiddleware(headerMiddleware(handler))),
		)
	}

	handle("GET", "/", s.indexHandler())

	// ユーザー登録・ログイン(user_handler.go)
	handle("GET", "/signup", s.willSignupHandler())
	handle("POST", "/signup", s.signupHandler())
	handle("GET", "/signin", s.willSigninHandler())
	handle("POST", "/signin", s.signinHandler())
	handle("POST", "/signout", s.signoutHandler())

	// ダイアリー系(diary_handler.go)
	handle("GET", "/diaries", s.diariesHandler())
	handle("GET", "/diaries/create", s.willDiaryCreateHandler())
	handle("POST", "/diaries/create", s.diaryCreateHandler())

	return router
}

func (s *server) findUser(r *http.Request) (user *model.User) {
	cookie, err := r.Cookie(sessionKey)
	if err == nil && cookie.Value != "" {
		user, _ = s.app.FindUserByToken(cookie.Value)
	}
	return
}

func (s *server) indexHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := s.findUser(r)
		s.renderTemplate(w, r, "index.tmpl", map[string]interface{}{
			"User": user,
		})
	})
}

var csrfMiddleware = func(next http.Handler) http.Handler {
	return nosurf.New(next)
}

var csrfToken = func(r *http.Request) string {
	return nosurf.Token(r)
}

func (s *server) renderTemplate(w http.ResponseWriter, r *http.Request, tmpl string, data map[string]interface{}) {
	if data == nil {
		data = make(map[string]interface{})
	}
	data["CSRFToken"] = csrfToken(r)
	err := templates[tmpl].ExecuteTemplate(w, "main.tmpl", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

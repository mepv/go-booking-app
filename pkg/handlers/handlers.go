package handlers

import (
	"net/http"

	"github.com/mepv/go-booking-app/pkg/config"
	"github.com/mepv/go-booking-app/pkg/model"
	"github.com/mepv/go-booking-app/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (repo *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.html", &model.TemplateData{})
}

func (repo *Repository) About(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."

	render.RenderTemplate(w, "about.page.html", &model.TemplateData{
		StringMap: stringMap,
	})
}

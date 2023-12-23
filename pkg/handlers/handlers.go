package handlers

import (
	"net/http"

	"github.com/mitanshukr/bookings-app/pkg/config"
	"github.com/mitanshukr/bookings-app/pkg/models"
	"github.com/mitanshukr/bookings-app/pkg/render"
)

type Repository struct {
	App *config.AppConfig
}

var Repo *Repository

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again, kumar!!!"

	aboutTempData := models.TemplateData{
		StringMap: stringMap,
	}

	render.RenderTemplate(w, "about.page.tmpl", &aboutTempData)
}

// func divideValues(x, y float32) (float32, error) {
// 	if y == 0 {
// 		var err error = fmt.Errorf("cannot %f divide by zero", x)
// 		return 0, err
// 	}
// 	result := x / y
// 	return result, nil
// }

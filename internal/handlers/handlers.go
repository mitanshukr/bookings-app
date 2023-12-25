package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mitanshukr/bookings-app/internal/config"
	"github.com/mitanshukr/bookings-app/internal/models"
	"github.com/mitanshukr/bookings-app/internal/render"
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
	render.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again, kumar!!!"

	aboutTempData := models.TemplateData{
		StringMap: stringMap,
	}

	render.RenderTemplate(w, r, "about.page.tmpl", &aboutTempData)
}

func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "make-reservation.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "generals.page.tmpl", &models.TemplateData{})
}

// Majors renders the room page
func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "majors.page.tmpl", &models.TemplateData{})
}

// Availability renders the search availability page
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "search-availability.page.tmpl", &models.TemplateData{})
}

// PostAvailability renders the search availability page
func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	startDate := r.Form.Get("start")
	endDate := r.Form.Get("end")

	w.Write([]byte(fmt.Sprintf("start date is %s and end date is %s", startDate, endDate)))
}

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

func (m *Repository) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {
	res := jsonResponse{
		OK:      true,
		Message: "Available!",
	}

	// out, _ := json.MarshalIndent(res, "", "  ")
	out, _ := json.Marshal(res)
	w.Write(out)
}

// Contact renders the contact page
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "contact.page.tmpl", &models.TemplateData{})
}

// func divideValues(x, y float32) (float32, error) {
// 	if y == 0 {
// 		var err error = fmt.Errorf("cannot %f divide by zero", x)
// 		return 0, err
// 	}
// 	result := x / y
// 	return result, nil
// }

package handlers

import (
	"net/http"

	"github.com/angusbean/weather-check/config"
)

//Repo the repository used by the handlers
var Repo *Repository

//Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

//PostLocation accepts JSON input and returns JSON weather information for that location
func (m *Repository) PostJSON(w http.ResponseWriter, r *http.Request) {

}

package routes

import (
	"net/http"

	"github.com/wilian1808/apirest-users/controllers"
)

// UpdateRoute func
func UpdateRoute(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		controllers.UpdateGetController(w, r)
	}
	if r.Method == http.MethodPost {
		controllers.UpdatePostController(w, r)
	}
}

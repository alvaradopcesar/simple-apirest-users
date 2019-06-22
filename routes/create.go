package routes

import (
	"net/http"

	"github.com/wilian1808/simple-apirest-users/controllers"
)

// CreateRoute func
func CreateRoute(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		controllers.CreateController(w, r)
	}
}

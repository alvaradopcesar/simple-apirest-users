package routes

import (
	"net/http"

	"github.com/wilian1808/simple-apirest-users/controllers"
)

// ReadRoute func
func ReadRoute(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		controllers.ReadController(w, r)
	}
}

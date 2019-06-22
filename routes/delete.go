package routes

import (
	"net/http"

	"github.com/wilian1808/simple-apirest-users/controllers"
)

// DeleteRoute func
func DeleteRoute(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		controllers.DeleteController(w, r)
	}
}

package helpers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/wilian1808/apirest-users/models"
)

// ResponseHelper func
func ResponseHelper(w http.ResponseWriter, response models.Response) {
	jsn, err := json.Marshal(response)
	if err != nil {
		log.Fatal("Error al formatear el mensaje: ", err.Error())
		return
	}
	w.Write(jsn)
}

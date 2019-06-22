package controllers

import (
	"log"
	"net/http"
	"strings"

	"github.com/wilian1808/apirest-users/configs"
	"github.com/wilian1808/apirest-users/helpers"
	"github.com/wilian1808/apirest-users/models"
)

// DeleteController func
func DeleteController(w http.ResponseWriter, r *http.Request) {

	id := strings.TrimPrefix(r.URL.Path, "/delete/")

	db := configs.DatabaseConfig()
	defer db.Close()

	query, err := db.Prepare("DELETE FROM users WHERE id=?")
	if err != nil {
		log.Fatal("Error al preparar la consulta: ", err.Error())
		return
	}

	_, err = query.Exec(&id)
	if err != nil {
		log.Fatal("Error al ejecutar la consulta: ", err.Error())
		return
	}

	var response models.Response
	response.Code = http.StatusOK
	response.Message = "Usuario eliminado con exito"

	helpers.ResponseHelper(w, response)
}

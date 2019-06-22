package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/wilian1808/apirest-users/configs"
	"github.com/wilian1808/apirest-users/helpers"
	"github.com/wilian1808/apirest-users/models"
)

// CreateController func
func CreateController(w http.ResponseWriter, r *http.Request) {
	var response models.Response
	var user models.User

	json.NewDecoder(r.Body).Decode(&user)

	db := configs.DatabaseConfig()
	defer db.Close()

	query, err := db.Prepare("INSERT users SET fullname=?, paternal=?, maternal=?, email=?")
	if err != nil {
		log.Fatal("Error al preparar la consulta: ", err.Error())
		return
	}

	_, err = query.Exec(&user.Fullname, &user.Paternal, &user.Maternal, &user.Email)
	if err != nil {
		log.Fatal("Error al executar la consulta: ", err.Error())
		return
	}

	response.Code = http.StatusOK
	response.Message = "Usuario registrado exitosamente"
	helpers.ResponseHelper(w, response)
}

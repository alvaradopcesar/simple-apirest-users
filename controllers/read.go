package controllers

import (
	"log"
	"net/http"

	"github.com/wilian1808/simple-apirest-users/configs"
	"github.com/wilian1808/simple-apirest-users/helpers"
	"github.com/wilian1808/simple-apirest-users/models"
)

// ReadController func
func ReadController(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	var response models.Response

	db := configs.ConnectionConfig()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal("Error al consultar la db: ", err.Error())
		return
	}

	for rows.Next() {
		var user models.User
		err = rows.Scan(&user.ID, &user.Fullname, &user.Paternal, &user.Maternal, &user.Email)
		if err != nil {
			log.Fatal("Error al formatear los datos: ", err.Error())
			return
		}
		users = append(users, user)
	}

	response.Code = http.StatusOK
	response.Message = "datos traidos con exito"
	response.Data = users

	helpers.ResponseHelper(w, response)
}

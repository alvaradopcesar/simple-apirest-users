package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/wilian1808/simple-apirest-users/configs"
	"github.com/wilian1808/simple-apirest-users/models"
)

// ReadController func
func ReadController(w http.ResponseWriter, r *http.Request) {
	var users []models.User

	db := configs.DatabaseConfig()
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

	jsn, err := json.Marshal(users)
	if err != nil {
		log.Fatal("Error al convertir a json: ", err.Error())
		return
	}

	w.Write(jsn)
}

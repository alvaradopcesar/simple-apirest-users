package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/wilian1808/simple-apirest-users/configs"
	"github.com/wilian1808/simple-apirest-users/helpers"
	"github.com/wilian1808/simple-apirest-users/models"
)

// UpdateGetController func
func UpdateGetController(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/update/")
	user := models.User{}

	db := configs.ConnectionConfig()
	defer db.Close()

	row := db.QueryRow("SELECT id, fullname, paternal, maternal, email FROM users WHERE id=?", &id)
	err := row.Scan(&user.ID, &user.Fullname, &user.Paternal, &user.Maternal, &user.Email)
	if err != nil {
		log.Fatal("Error al formatear los datos: ", err.Error())
		return
	}

	var response models.Response
	var users []models.User

	users = append(users, user)

	response.Code = http.StatusOK
	response.Message = "dato traido con exito"
	response.Data = users

	helpers.ResponseHelper(w, response)
}

// UpdatePostController func
func UpdatePostController(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/update/")
	user := models.User{}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatal("Error al formatear los datos: ", err.Error())
		return
	}

	db := configs.ConnectionConfig()
	defer db.Close()

	query, err := db.Prepare("UPDATE users SET fullname=?, paternal=?, maternal=?, email=? WHERE id=?")
	if err != nil {
		log.Fatal("Error al preparar la consulta: ", err.Error())
		return
	}

	_, err = query.Exec(&user.Fullname, &user.Paternal, &user.Maternal, &user.Email, &id)
	if err != nil {
		log.Fatal("Error al ejecutar la consulta: ", err.Error())
		return
	}

	var response models.Response
	response.Code = http.StatusOK
	response.Message = "Usuario actualizado con exito"

	helpers.ResponseHelper(w, response)
}

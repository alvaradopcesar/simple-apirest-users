package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/wilian1808/simple-apirest-users/helpers"

	"github.com/wilian1808/simple-apirest-users/models"

	"github.com/wilian1808/simple-apirest-users/configs"
)

// UpdateGetController func
func UpdateGetController(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/update/")
	user := models.User{}

	db := configs.DatabaseConfig()
	defer db.Close()

	row := db.QueryRow("SELECT id, fullname, paternal, maternal, email FROM users WHERE id=?", &id)
	err := row.Scan(&user.ID, &user.Fullname, &user.Paternal, &user.Maternal, &user.Email)
	if err != nil {
		log.Fatal("Error al formatear los datos: ", err.Error())
		return
	}

	jsn, err := json.Marshal(&user)
	if err != nil {
		log.Fatal("Error al convertir a json: ", err.Error())
		return
	}

	w.Write(jsn)
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

	db := configs.DatabaseConfig()
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

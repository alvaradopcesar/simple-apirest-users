package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	db "github.com/wilian1808/simple-apirest-users/configs"
	"github.com/wilian1808/simple-apirest-users/helpers"
	"github.com/wilian1808/simple-apirest-users/models"
)

// CreateController func
func CreateController(w http.ResponseWriter, r *http.Request) {
	// var response models.Response
	var user models.User

	json.NewDecoder(r.Body).Decode(&user)
	fmt.Println("hola")

	query, err := db.Dbmap.Prepare("INSERT users SET fullname=?, paternal=?, maternal=?, email=?")
	if err != nil {
		log.Fatal("Error al preparar la consulta: ", err.Error())
		return
	}

	_, err = query.Exec(&user.Fullname, &user.Paternal, &user.Maternal, &user.Email)
	if err != nil {
		log.Fatal("Error al executar la consulta: ", err.Error())
		return
	}

	helpers.ResponseHelper(w, models.Response{Code: http.StatusOK, Message: "Usuario registrado exitosamente"})
}

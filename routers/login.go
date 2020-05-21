package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/frank1995alfredo/twittor-go/bd"
	"github.com/frank1995alfredo/twittor-go/jwt"
	"github.com/frank1995alfredo/twittor-go/models"
)

//Login ... Realiza el login
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Usuario y/o Contraseña invalidos "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El email de usuario es requerido", 400)
		return
	}

	documento, existe := bd.IntentoLogin(t.Email, t.Password)
	if existe == false {
		http.Error(w, "Usuario y/o Contraseña invalidos", 400)
		return
	}

	jwtKey, err := jwt.GeneroJWT(documento)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar generar el Token correspondiente "+err.Error(), 400)
		return
	}

	resp := models.RespuestaLogin{
		Token: jwtKey,
	}

	w.Header().Set("Content-Tyoe", "application/json")
	w.WriteHeader(http.StatusCreated) //devuelve un 200
	json.NewEncoder(w).Encode(resp)

	//guardar una coockie
	expirationTime := time.Now().Add(24 * time.Hour) //calcula la fecha de hoy
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})

}

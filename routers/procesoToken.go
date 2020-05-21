package routers

import (
	"errors"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/frank1995alfredo/twittor-go/bd"
	"github.com/frank1995alfredo/twittor-go/models"
)

//Email ... valor de Email usado en todos los EndPoints
var Email string 

//IDUsuario ... es el ID devuelto del modelo, que se usara en todos los EndPoints
var IDUsario string

//ProcesoToken ... koin 
func ProcesoToken(tk string) (*models.Claim, bool, string, error) {
	miClave := []byte("1234")
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2{
		return claims, false, string(""), errors.New("formato de token invalido")
	}
    //quita los espacios del token
	tk = strings.TrimSpace(splitToken[1])


	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token)(interface{}, error){
		return miClave, nil
	})
	if err == nil{
		_, encontrado, _ := bd.ChequeoYaExisteUsuario(claims.Email)
		if encontrado == true{
			Email = claims.Email
			IDUsario = claims.ID.Hex()
		}
		return claims, encontrado, IDUsario, nil
	}
    if !tkn.Valid {
	    return claims, false, string(""), errors.New("token Invalido")
	}

	return claims, false, string(""), err

}

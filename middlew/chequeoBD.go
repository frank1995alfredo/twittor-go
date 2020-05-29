package middlew

import (
	"net/http"

	"github.com/frank1995alfredo/twittor-go/bd"
)

//ChequeoBD ... func
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConnection() == 0 {
			http.Error(w, "Conexion perdida con la Base de Datos", 500)
			return
		}
		next.ServeHTTP(w, r)

	}

}

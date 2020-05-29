package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/frank1995alfredo/twittor-go/middlew"
	"github.com/frank1995alfredo/twittor-go/routers"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

//Manejadores ... rutas
func Manejadores() {
	router := mux.NewRouter()

	//tipo de llamado, en este caso es por POST
	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")
	router.HandleFunc("/login", middlew.ChequeoBD(routers.Login)).Methods("POST")
	router.HandleFunc("/verperfil", middlew.ChequeoBD(middlew.ValidoJWT(routers.VerPerfil))).Methods("GET")
	router.HandleFunc("/modificarPerfil", middlew.ChequeoBD(middlew.ValidoJWT(routers.ModificarPerfil))).Methods("PUT")
	router.HandleFunc("/tweet", middlew.ChequeoBD(middlew.ValidoJWT(routers.GraboTweet))).Methods("POST")
	router.HandleFunc("/leoTweet", middlew.ChequeoBD(middlew.ValidoJWT(routers.LeoTweets))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router) //da permiso a cualquiera para que pueda acceder a la aplicacion
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}

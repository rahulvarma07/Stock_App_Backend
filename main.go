package main

import (
	"context"
	"log"
	"net/http"

	database "rahulvarma07/github.com/DATABASE"
	routers "rahulvarma07/github.com/ROUTERS"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("There's an err loading env files")
	}
}

func main() {
	defer database.GetMongoCLient().Disconnect(context.Background())

	rout := mux.NewRouter()

	rout.HandleFunc("/hey", routers.HomePage).Methods("GET")
	rout.HandleFunc("/login", routers.LoginTheUser).Methods("POST")
	rout.HandleFunc("/signup", routers.SignUpTheUser).Methods("POST")

	c := cors.New(cors.Options{
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST"},
	})

	handler := c.Handler(rout)
	log.Println("Running Successfully")
	http.ListenAndServe(":9000", handler)
}

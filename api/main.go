package main

import (
	"log"
	"net/http"

	"github.com/fedegmz/api-calendar/db"
	"github.com/fedegmz/api-calendar/models"
	"github.com/fedegmz/api-calendar/routes"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {

	log.Println("Starting API server...")
	log.Println("Server running on port 4000")

	db.DBConection()
	db.DB.AutoMigrate(models.Citas{})

	r := mux.NewRouter()
	
	r.HandleFunc("/api/citas", routes.GetCitasHandler).Methods("GET")
	r.HandleFunc("/api/citas", routes.CreateCitasHandler).Methods("POST")
	r.HandleFunc("/api/citas/{id}", routes.GetCitaHandler).Methods("GET")
	r.HandleFunc("/api/citas/year/{year}", routes.GetCitasFormYearHandler).Methods("GET")
	r.HandleFunc("/api/citas/{id}", routes.DeleteCitaHandler).Methods("DELETE")
	r.HandleFunc("/api/citas/{id}", routes.UpdateCitaHandler).Methods("PUT")
    
	// Configurar las opciones CORS
	corsOptions := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	})

    // Aplicar el middleware CORS a las rutas
	handler := corsOptions.Handler(r)

	http.ListenAndServe(":3000", handler)

}

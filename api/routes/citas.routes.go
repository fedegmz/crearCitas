package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/fedegmz/api-calendar/db"
	"github.com/fedegmz/api-calendar/models"
	"github.com/gorilla/mux"
)


func GetCitasHandler(w http.ResponseWriter, r *http.Request) {

	var citas []models.Citas

	db.DB.Find(&citas)

	citasCount := len(citas)

	if citasCount == 0 {
		w.WriteHeader(http.StatusNotFound)
		respose:= models.OnlyMessage{
			Success: false,
			Message: "Citas not found",
		}
		json.NewEncoder(w).Encode(&respose)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&citas)

}

func GetCitasFormYearHandler(w http.ResponseWriter, r *http.Request) {

	var citas []models.Citas
	params := mux.Vars(r)
	log.Printf("Month: %v\n", params["year"])
	db.DB.Where("fecha_inicio LIKE ?", params["year"]+"%").Find(&citas)

	citasCount := len(citas)

	if citasCount == 0 {
		w.WriteHeader(http.StatusNotFound)
		respose:= models.OnlyMessage{
			Success: false,
			Message: "Citas not found",
		}
		json.NewEncoder(w).Encode(&respose)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&citas)

}

func CreateCitasHandler(w http.ResponseWriter, r *http.Request) {
	var citas models.Citas
	

	json.NewDecoder(r.Body).Decode(&citas)

	if !CheckDate(citas){
		w.WriteHeader(http.StatusOK)
		response := models.OnlyMessage{
			Success: false,
			Message: "La fecha de inicio no puede ser menor a la fecha actual",
		}
		json.NewEncoder(w).Encode(&response)
		return
	}

	//select where fecha_inicio between fecha_inicio and fecha_fin
	db.DB.Where("fecha_inicio <= ? AND fecha_fin >= ?", citas.FechaInicio, citas.FechaInicio).First(&citas)

	if citas.ID != 0 {
		w.WriteHeader(http.StatusOK)
		response := models.OnlyMessage{
			Success: false,
			Message: "La fecha de inicio no puede estar entre la fecha de inicio y fin de otra cita",
		}
		json.NewEncoder(w).Encode(&response)
		return
	}

	createdCitas := db.DB.Create(&citas)

	err := createdCitas.Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Error al crear la cita");
		return
	}

	response := models.OnlyMessage{
		Success: true,
		Message: "Cita created successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&response)

	//json.NewEncoder(w).Encode(&citas)
}

func GetCitaHandler(w http.ResponseWriter, r *http.Request) {

	var citas models.Citas
	params := mux.Vars(r)

	db.DB.First(&citas, params["id"])

	if citas.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Type", "application/json")
		response:= models.OnlyMessage{
			Success: false,
			Message: "Cita not found",
		}
		json.NewEncoder(w).Encode(&response)
		return
	}
	response:= models.Response{
		Success: true,
		Message: "Cita found successfully",
		Citas: citas,
	}
	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&response)


}

func DeleteCitaHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("DeleteCitaHandler")
	var citas models.Citas
	params := mux.Vars(r)

	db.DB.First(&citas, params["id"])

	if citas.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - Cita not found"))
		return
	}

	db.DB.Unscoped().Delete(&citas)

	w.WriteHeader(http.StatusOK)
	response:= models.OnlyMessage{
		Success: true,
		Message: "Cita deleted successfully",
	}
	json.NewEncoder(w).Encode(&response)

}

func UpdateCitaHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("UpdateCitaHandler")
	var citas models.Citas
	params := mux.Vars(r)

	db.DB.First(&citas, params["id"])

	if citas.ID == 0 {
		w.WriteHeader(http.StatusOK)
		response:= models.OnlyMessage{
			Success: false,
			Message: "Cita not found",
		}
		json.NewEncoder(w).Encode(&response)
		return
	}

	json.NewDecoder(r.Body).Decode(&citas)

	if !CheckDate(citas){
		w.WriteHeader(http.StatusOK)
		response := models.OnlyMessage{
			Success: false,
			Message: "La fecha de inicio no puede ser menor a la fecha actual",
		}
		json.NewEncoder(w).Encode(&response)
		return
	}

	//select where fecha_inicio between fecha_inicio and fecha_fin
	existingCitas := []models.Citas{}
	db.DB.Where("fecha_inicio <= ? AND fecha_fin >= ? AND id != ?", citas.FechaInicio, citas.FechaInicio, citas.ID).Find(&existingCitas)

	if len(existingCitas) > 0 {
		w.WriteHeader(http.StatusOK)
		response := models.OnlyMessage{
			Success: false,
			Message: "La fecha de inicio no puede estar entre la fecha de inicio y fin de otra cita",
		}
		json.NewEncoder(w).Encode(&response)
		return
	}


	updateCita:= db.DB.Save(&citas)

	err := updateCita.Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Error al actualizar la cita");
		return
	}

	

	w.WriteHeader(http.StatusOK)
	response:= models.OnlyMessage{
		Success: true,
		Message: "Cita updated successfully",
	}
	json.NewEncoder(w).Encode(&response)
}

func CheckDate(citas models.Citas) bool {
    // Cargar la ubicación de la zona horaria correcta (por ejemplo, "America/Mexico_City")
    location, err := time.LoadLocation("America/Mexico_City")
    if err != nil {
		log.Println("Error al cargar la ubicación de la zona horaria")
        log.Println(err)
        return false
    }

    now := time.Now().In(location)
    fechaInicio, err := time.Parse("2006-01-02 15:04:05 -0700 MST", citas.FechaInicio+" -0600 CST")
	fechaFin, err := time.Parse("2006-01-02 15:04:05 -0700 MST", citas.FechaFin+" -0600 CST")
    
	if err != nil {
		log.Println("Error al parsear la fecha")
        log.Println(err)
        return false
    }


	log.Printf("Fecha actual: %v\n", now)
	log.Printf("Fecha inicio: %v\n", fechaInicio.In(location))
	log.Printf("Fecha fin: %v\n", fechaFin.In(location))
    
	// Comparar la fechaInicio y fechaFin
	if fechaInicio.In(location).After(fechaFin.In(location)) {
		log.Println("La fecha de inicio no puede ser mayor a la fecha de fin")
		return false
	}

	// Comparar la fecha y hora completa
    if now.After(fechaInicio.In(location)) {
		log.Println("La fecha de inicio no puede ser menor a la fecha actual")
        return false
    }

    return true
}


package main

import (
	"fmt"
	"net/http"
)

const rutaMutant = "mutant/"
const rutaStats = "stats"

//Invocar controlador para servicio Mutant
func handleMutant(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Mutante"))
	//context.JSON(http.StatusOK, map[string]interface{}{"ruta": "mutatn"})
	//mutantController(w, r)
}

//Invocar controlador para servicio Stats
func handleStats(w http.ResponseWriter, r *http.Request) {
	//context.JSON(http.StatusOK, map[string]interface{}{"ruta": "stats"})
	//statsController(w, r)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Stats"))
}

// SetupRoutes :
func SetupRoutes(rutabase string) {
	mutantHandler := http.HandlerFunc(handleMutant)
	statsHandler := http.HandlerFunc(handleStats)
	http.Handle(fmt.Sprintf("%s/%s", rutabase, rutaMutant), mutantHandler)
	http.Handle(fmt.Sprintf("%s/%s", rutabase, rutaStats), statsHandler)
}

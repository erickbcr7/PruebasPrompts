package main

import (
	"encoding/json"
	"fmt"
	"net/http"

)

func saludoHandler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	nombre := queryParams.Get("nombre")

	w.Header().Set("Content-Type", "application/json")

	if nombre != "" {
		response := map[string]string{"mensaje": fmt.Sprintf("Hola, %s", nombre)}
		json.NewEncoder(w).Encode(response)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		response := map[string]string{"error": "El parámetro \"nombre\" es requerido"}
		json.NewEncoder(w).Encode(response)
	}
}

func main() {
	http.HandleFunc("/saludo", saludoHandler)
	fmt.Println("Servidor escuchando en el puerto 8080")
	http.ListenAndServe(":8080", nil)
}
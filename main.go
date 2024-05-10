package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/tasks", handleTasks)
	http.HandleFunc("/tasks/", handleTask)

	fmt.Println("Servidor en ejecución en el puerto 8080...")
	http.ListenAndServe(":8080", nil)
}

func handleTasks(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		GetTasks(w, r)
	} else if r.Method == http.MethodPost {
		CreateTask(w, r)
	} else {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}
}

func handleTask(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		GetTask(w, r)
	} else if r.Method == http.MethodPut {
		UpdateTask(w, r)
	} else if r.Method == http.MethodDelete {
		DeleteTask(w, r)
	} else {
		http.Error(w, "Métpdo no permitido", http.StatusMethodNotAllowed)
		return
	}
}

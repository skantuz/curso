package controllers

import (
	"encoding/json"
	"net/http"
)

func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, Content-Length, Accept-Encoding")
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

var Index = func(w http.ResponseWriter, r *http.Request) {
	Respond(w, Message(true, "Systema en linea"))
}

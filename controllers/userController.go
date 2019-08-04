package controllers 

import (
	m "github.com/skantuz/curso/models"
	"net/http"
)

var CreateUser = func(w http.ResponseWriter, r *http.Request){
	user := &m.USer{}
	err := json.NewDecoder(r.body).Decode(user)
	if err != nil{
		Respond(w, Message(false,err))
		return
	}
	resp := user.Create()
	Respond(w,resp)
}

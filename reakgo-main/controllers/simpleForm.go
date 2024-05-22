package controllers

import (
	"log"
	"net/http"
	"reakgo/models"
)

func AddForm(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		} else {
			// Logic
			taskname := r.FormValue("taskname")
			deadline := r.FormValue("deadline")

			log.Println(taskname, deadline)
			models.FormAddView{}.Add(taskname, deadline)
		}
	}
	Helper.RenderTemplate(w, r, "addForm", nil)
}

func ViewForm(w http.ResponseWriter, r *http.Request) {
	result, err := models.FormAddView{}.View()
	if err != nil {
		log.Println(err)
	}
	Helper.RenderTemplate(w, r, "viewForm", result)
}

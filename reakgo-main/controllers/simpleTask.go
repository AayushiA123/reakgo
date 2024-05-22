package controllers

import (
	"fmt"
	"log"
	"net/http"
	"reakgo/models"
	"strconv"
	"time"
)

func AddTask(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		} else {
			// Logic
			taskname := r.FormValue("taskname")
			deadline := r.FormValue("deadline")

			format := "2006-01-02T15:04"
			t, err := time.Parse(format, deadline)
			fmt.Println("Error", err)
			fmt.Printf("%v", t.Unix())
			log.Println(taskname, "deadline", t.Unix())

			if taskname == "" || deadline == "" {
				http.Error(w, "Please fill out all fields", http.StatusBadRequest)
				return
			}
			models.TaskAddView{}.Add(taskname, t.Unix())
		}
	}
	result, err := models.TaskAddView{}.View()

	if err != nil {
		log.Println(err)
	}
	Helper.RenderTemplate(w, r, "addTask", result)
}

// func DeleteTask(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == "POST" {
// 		err := r.ParseForm()

// 		// Retrieve the task ID from the form
// 		taskID := r.FormValue("id")

// 		// Convert task ID to integer
// 		id, err := strconv.Atoi(taskID)

// 		// Delete the task from the database
// 		//err = models.DeleteTaskByID(id)
// 		err = models.TaskAddView{}.Delete(id)
// 		if err != nil {
// 			log.Println(err)
// 			http.Error(w, "Failed to delete task", http.StatusInternalServerError)
// 			return
// 		}

// 		// Redirect to the add task page after successful deletion
// 		http.Redirect(w, r, "/addTask", http.StatusSeeOther)
// 	} else {
// 		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
// 	}
// }

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Retrieve the task ID from the form
		taskID := r.FormValue("id")
		if taskID == "" {
			http.Error(w, "Task ID is missing", http.StatusBadRequest)
			return
		}

		// Convert task ID to int
		id, err := strconv.Atoi(taskID)
		if err != nil {
			log.Println(err)
			http.Error(w, "Invalid Task ID", http.StatusBadRequest)
			return
		}

		// Convert int to int32
		id32 := int32(id)

		// Delete the task using TaskAddView's Delete method
		err = models.TaskAddView{}.Delete(id32)
		if err != nil {
			log.Println(err)
			http.Error(w, "Failed to delete task", http.StatusInternalServerError)
			return
		}

		// Redirect to the add task page after successful deletion
		http.Redirect(w, r, "/addTask", 301)
	}
}

func CompleteTask(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Retrieve the task ID from the form
		taskID := r.FormValue("id")
		if taskID == "" {
			http.Error(w, "Task ID is missing", http.StatusBadRequest)
			return
		}

		// Convert task ID to int
		id, err := strconv.Atoi(taskID)
		if err != nil {
			log.Println(err)
			http.Error(w, "Invalid Task ID", http.StatusBadRequest)
			return
		}

		// Convert int to int32
		id32 := int32(id)

		err = models.TaskAddView{}.Complete(id32)
		if err != nil {
			log.Println(err)
			http.Error(w, "Failed to delete task", http.StatusInternalServerError)
			return
		}

		// Redirect to the add task page after successful deletion
		http.Redirect(w, r, "/addTask", 301)
	}
}

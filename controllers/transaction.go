package controllers

import (
	"finance-management-web/models"
	"html/template"
	"net/http"
	"strconv"
)

const REDIRECT = 301

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	trasanctions := models.SearchTransactions()
	temp.ExecuteTemplate(w, "index", trasanctions)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "new", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		value, _ := strconv.ParseFloat(r.FormValue("value"), 64)
		date := r.FormValue("date")

		models.CreateTransaction(name, description, date, value)
	}

	http.Redirect(w, r, "/", REDIRECT)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	models.DeleteTransaction(id)
	http.Redirect(w, r, "/", REDIRECT)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	transaction := models.EditProduct(id)
	temp.ExecuteTemplate(w, "edit", transaction)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id, _ := strconv.Atoi(r.FormValue("id"))
		name := r.FormValue("name")
		description := r.FormValue("description")
		value, _ := strconv.ParseFloat(r.FormValue("value"), 64)
		date := r.FormValue("date")

		models.UpdateTransaction(id, name, description, date, value)
	}

	http.Redirect(w, r, "/", REDIRECT)
}

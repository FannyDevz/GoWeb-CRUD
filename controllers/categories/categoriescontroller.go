package categories

import (
	"goweb/entities"
	"goweb/models/categoriesmodel"
	"html/template"
	"net/http"
	"strconv"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	var categories = categoriesmodel.GetAll()
	data := map[string]any{
		"categories": categories,
	}
	temp, err := template.ParseFiles("views/categories/index.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(w, data)
}
func Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/categories/create.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(w, nil)
	}
	if r.Method == "POST" {
		var category entities.Category
		category.Name = r.FormValue("name")
		category.Type = r.FormValue("type")
		category.CreatedAt = time.Now()
		category.UpdatedAt = time.Now()
		if ok := categoriesmodel.Add(category); !ok {
			temp, _ := template.ParseFiles("views/categories/create.html")
			temp.Execute(w, nil)
		}
		http.Redirect(w, r, "/categories", 301)
	}
}

func Edit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var category = categoriesmodel.GetById(r.URL.Query().Get("id"))
		data := map[string]any{
			"category": category,
		}
		temp, err := template.ParseFiles("views/categories/edit.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(w, data)
	}
	if r.Method == "POST" {
		var category entities.Category
		idStr := r.FormValue("id")

		id, err := strconv.Atoi(idStr)
		if err != nil {
			panic(err)
		}

		category.Name = r.FormValue("name")
		category.Type = r.FormValue("type")
		category.UpdatedAt = time.Now()

		if ok := categoriesmodel.Update(id, category); !ok {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
			return
		}
		http.Redirect(w, r, "/categories", http.StatusSeeOther)
	}
}
func Delete(w http.ResponseWriter, r *http.Request) {

	idStr := r.FormValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		panic(err)
	}

	if err := categoriesmodel.Delete(id); err != nil {
		panic(err)
	}
	http.Redirect(w, r, "/categories", http.StatusSeeOther)

}

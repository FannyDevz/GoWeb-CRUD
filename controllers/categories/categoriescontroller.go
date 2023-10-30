package categories

import "net/http"
import "goweb/models/categoriesmodel"

func Index(w http.ResponseWriter, r *http.Request) {
	categories := categoriesmodel.GetAll()
	data := map[stromg]any{
		"categories": categories,
	}
}
func Add(w http.ResponseWriter, r *http.Request) {

}
func Edit(w http.ResponseWriter, r *http.Request) {

}
func Delete(w http.ResponseWriter, r *http.Request) {

}

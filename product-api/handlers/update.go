package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
	"shop.com/v1/data"
	"strconv"
)

func (p *Products) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle put Products")
	vars :=mux.Vars(r)
	id, err:=strconv.Atoi(vars["id"])
	if err!=nil {
		http.Error(w, "unable to convert id", http.StatusBadRequest)
	}
	//get the data being received from the request body and convert it into json
	prod := r.Context().Value(KeyProduct{}).(data.Product)
	err = prod.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "unable to unmarshal json", http.StatusBadRequest)
	}
	err = data.UpdateProduct(id, &prod)
	if err == data.ErrProductNotFound {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "Product not found", http.StatusInternalServerError)
		return
	}
}

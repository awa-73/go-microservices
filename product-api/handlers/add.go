package handlers

import (
	"net/http"
	"shop.com/v1/data"
)

func (p *Products) AddProduct(w http.ResponseWriter, r *http.Request) {
	prod := r.Context().Value(KeyProduct{}).(data.Product)
	data.AddProduct(&prod)
}
